// Code for genetic algorithms
package genetic

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kevinburke/rct/Godeps/_workspace/src/github.com/pborman/uuid"

	"github.com/kevinburke/rct/tracks"
)

var directory *string

func init() {
	directory = flag.String("directory", "/usr/local/rct", "Path to the folder storing RCT experiment data")
}

// constants which may be altered to affect the ride runtime
const PARENTS = 2
const FIT_PERCENTAGE = 0.2
const MUTATION_RATE = 0.05

// crossover with a probability of 0.6 (taken from the book & De Jong 1975)
const CROSSOVER_PROBABILITY = 0.6
const POOL_SIZE = 500
const ITERATIONS = 1000
const PRINT_RESULTS_EVERY = 5

type ExperimentMetadata struct {
	Hash                 string
	Date                 time.Time
	Notes                string
	Runtime              time.Duration
	PoolSize             int16
	Iterations           int32
	CrossoverProbability float32
	MutationRate         float32
}

func Run(packageRoot string) error {
	if directory == nil {
		return fmt.Errorf("invalid directory - need to specify it")
	}
	expDir := path.Join(*directory, "experiments")
	err := os.MkdirAll(expDir, 0755)
	if err != nil {
		return err
	}
	id := fmt.Sprintf("exp_%s", uuid.New())
	expIdDir := path.Join(expDir, id)
	err = os.MkdirAll(expIdDir, 0755)
	if err != nil {
		return err
	}
	iterationsDir := path.Join(expIdDir, "iterations")
	err = os.MkdirAll(iterationsDir, 0755)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = packageRoot
	hashb, err := cmd.Output()
	if err != nil {
		return err
	}
	mtd := ExperimentMetadata{
		Hash:                 strings.TrimSpace(string(hashb)),
		Date:                 time.Now().UTC(),
		Notes:                "(none)", // XXX
		CrossoverProbability: CROSSOVER_PROBABILITY,
		PoolSize:             POOL_SIZE,
		MutationRate:         MUTATION_RATE,
		Iterations:           ITERATIONS,
	}
	metadataPath := path.Join(expIdDir, "meta.json")
	err = encode(metadataPath, mtd)
	if err != nil {
		return err
	}
	fmt.Printf("Experiment %s\n======================================\n", id)
	pool := SeedPool(POOL_SIZE)
	pool.Id = id
	for i := 0; i < ITERATIONS; i++ {
		pool = pool.Crossover()
		pool.Mutate(MUTATION_RATE)
		pool.Evaluate()
		iterationDir := path.Join(iterationsDir, strconv.Itoa(i))
		err = os.MkdirAll(iterationDir, 0755)
		if err != nil {
			return err
		}
		pool.Statistics(i, *directory)
	}
	return nil
}

type Pool struct {
	Id      string
	Members []*Member
}

type Member struct {
	Id    string
	Score int64
	// Advantage or disadvantage in reproducing
	Fitness   float64
	Runtime   time.Duration
	Track     []tracks.Element
	ScoreData scoreData
}

type scoresArray [500]int64

func (a *scoresArray) Len() int {
	return len(a)
}

func (a *scoresArray) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a *scoresArray) Less(i, j int) bool {
	return a[i] < a[j]
}

func (p *Pool) Statistics(iteration int, outputDirectory string) {
	var scores scoresArray
	var highestScore int64 = -1
	var worstScore int64 = 1000 * 1000 * 1000
	bestMember := new(Member)
	worstMember := new(Member)
	for i := 0; i < len(p.Members); i++ {
		scores[i] = p.Members[i].Score
		if p.Members[i].Score > highestScore {
			highestScore = p.Members[i].Score
			bestMember = p.Members[i]
		}
		if p.Members[i].Score < worstScore {
			worstScore = p.Members[i].Score
			worstMember = p.Members[i]
		}
	}
	if iteration%PRINT_RESULTS_EVERY == 0 {
		middle := len(scores) / 2
		median := (scores[middle] + scores[middle-1]) / 2
		sort.Sort(&scores)
		bestScorer := fmt.Sprintf("\t(length: %d, collisions: %d, to completion: %d, negative speed points: %d)\n\n",
			bestMember.ScoreData.Length, bestMember.ScoreData.Collisions,
			bestMember.ScoreData.Distance, bestMember.ScoreData.NegativeSpeed)
		fmt.Printf("Iteration %d: %d members, best member %s has score %d, "+
			"median %d, worst has score %d\n%s",
			iteration, len(p.Members), bestMember.Id, bestMember.Score, median,
			worstScore, bestScorer)
		if os.Getenv("DEBUG_BEST_TRACK") == "true" {
			if iteration%20 == 0 && iteration > 0 {
				for _, elem := range bestMember.Track {
					fmt.Printf("%s %t\n", elem.Segment.String(), elem.ChainLift)
				}
				fmt.Println("==================")
			}
		}
	}

	pth := path.Join(outputDirectory, "experiments", p.Id,
		"iterations", strconv.Itoa(iteration), fmt.Sprintf("%s.json", bestMember.Id))
	err := encode(pth, bestMember)
	if err != nil {
		log.Print(err)
	}
	pth = path.Join(outputDirectory, "experiments", p.Id,
		"iterations", strconv.Itoa(iteration), fmt.Sprintf("%s.json", worstMember.Id))
	err = encode(pth, worstMember)
	if err != nil {
		log.Print(err)
	}
}

// Create an initial pool
func SeedPool(size int) *Pool {
	// 1. Create a station of length 10
	// 2. For the start station piece, generate a list of possible pieces.
	// 3. Choose one at random. Advance a pointer one forward.
	// 4. Repeat for 50 pieces (Woodchip is length 108. Mischief is 123)
	members := make([]*Member, POOL_SIZE)
	for i := 0; i < POOL_SIZE; i++ {
		track := CreateStation()
		for j := STATION_LENGTH - 1; j < INITIAL_TRACK_LENGTH-STATION_LENGTH; j++ {
			poss := track[j].Possibilities()
			track = append(track, poss[rand.Intn(len(poss))])
		}
		score, d := GetScore(track)
		members[i] = &Member{
			Id:        fmt.Sprintf("iter_%s", uuid.New()),
			Track:     track,
			Score:     score,
			ScoreData: d,
		}
		//if i%100 == 0 {
		//for _, elem := range track {
		//fmt.Println(elem.Segment.String())
		//}
		//fmt.Printf("========\n")
		//}
	}
	return &Pool{Members: members}
}

func (p *Pool) Mutate(rate float64) {
	// for each ride:
	// for each position in the ride:
	// add in a possibility of mutation - addition
	// if addition:
	//	find a piece that is compatible with both ends.
	//  if no piece:
	//	  advance to the next track piece and try again
}

// Assign scores for every member of the pool
func (p *Pool) Evaluate() {
	var wg sync.WaitGroup
	for i := 0; i < POOL_SIZE; i++ {
		wg.Add(1)
		go func(i int, track []tracks.Element) {
			p.Members[i].Score, p.Members[i].ScoreData = GetScore(track)
			wg.Done()
		}(i, p.Members[i].Track)
	}
	wg.Wait()

	// Assign fitness for every member. In the future, consider a smarter
	// algorithm higher score members a better chance of reproducing.
	for i := 0; i < POOL_SIZE; i++ {
		p.Members[i].Fitness = float64(p.Members[i].Score)
	}
}

// Select chooses a member of the population at random
func (p *Pool) Select() (int, *Member) {
	// Stupid dumb version of this taken from here:
	// http://eli.thegreenplace.net/2010/01/22/weighted-random-generation-in-python
	// If it's a bottleneck, rewrite it.
	var weightedTotal float64 = 0
	totals := make([]float64, len(p.Members))
	for i := 0; i < len(p.Members); i++ {
		weightedTotal += max(p.Members[i].Fitness, 0)
		totals[i] = weightedTotal
	}
	rnd := rand.Float64() * weightedTotal
	for index, element := range totals {
		if rnd < element {
			return index, p.Members[index]
		}
	}
	return -1, &Member{}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// crossPoint1 = splice point in parent1.
func crossoverAtPoint(parent1 *Member, parent2 *Member, crossPoint1 int) (*Member, *Member) {
	// crossPoint2 = splice point in parent2.
	crossPoint2 := crossPoint1 + 1
	foundMatch := false
	for {
		// say cross point 1 = 3, track length = 6
		// parent 1 == 0-3, parent 2 = 4-5
		if tracks.Compatible(parent1.Track[crossPoint1], parent2.Track[crossPoint2]) {
			// Increment so we splice *after* the cross point in parent 1
			crossPoint1++
			foundMatch = true
			break
		}
		crossPoint1++
		if crossPoint1 >= len(parent1.Track) {
			break
		}
		if tracks.Compatible(parent1.Track[crossPoint1], parent2.Track[crossPoint2]) {
			foundMatch = true
			// Increment so we splice *after* the cross point in parent 1
			crossPoint1++
			break
		}
		crossPoint2++
		if crossPoint2 >= len(parent2.Track) {
			break
		}
	}
	//	swap the track pieces at the chosen point on track A and track B
	if foundMatch {
		c1, c2 := Swap(parent1, parent2, crossPoint1, crossPoint2)
		if os.Getenv("DEBUG_SWAPS") == "true" && crossPoint1 > 4 {
			fmt.Println("swapped at", crossPoint1, crossPoint2)
			fmt.Println("parent1")
			printTrack(parent1.Track)
			fmt.Println("parent2")
			printTrack(parent2.Track)
			fmt.Println("child1")
			printTrack(c1.Track)
			fmt.Println("child2")
			printTrack(c2.Track)
		}
		return c1, c2
	}
	return parent1, parent2
}

// crossoverOne crosses over the parent tracks. if point is -1, a random point
// is chosen. (point is used for testing)
func crossoverOne(parent1 *Member, parent2 *Member, point int) (*Member, *Member) {
	//	choose a random point between the beginning and the end
	if point == -1 {
		minval := min(len(parent1.Track)-1, len(parent2.Track)-1)
		point = rand.Intn(minval)
	}
	return crossoverAtPoint(parent1, parent2, point)
}

func printTrack(t []tracks.Element) {
	for i, elem := range t {
		fmt.Printf("%d %s\n", i, elem.Segment.String())
	}
}

// Crossover chooses two members of a pool and joins them at random.
func (p *Pool) Crossover() *Pool {
	halfLen := len(p.Members) / 2
	for i := 0; i < halfLen; i++ {
		// select 2 parents at random
		idx1, parent1 := p.Select()
		idx2, parent2 := p.Select()
		if idx1 == -1 || idx2 == -1 {
			continue
		}
		if idx1 == idx2 {
			// No point in crossing over with ourself
			continue
		}
		if rand.Float64() < CROSSOVER_PROBABILITY {
			// XXX delete parents
			child1, child2 := crossoverOne(parent1, parent2, -1)
			p.Members[idx1] = child1
			p.Members[idx2] = child2
		}
	}
	return p
}

// Swap creates two children out of the parents, by crossing over the tracks at
// the given cross points. The sum of the two track lengths may be the same,
// but the tracks themselves will change.
func Swap(parent1 *Member, parent2 *Member, crossPoint1 int, crossPoint2 int) (*Member, *Member) {
	child1len := crossPoint1 + (len(parent2.Track) - crossPoint2)
	child2len := crossPoint2 + (len(parent1.Track) - crossPoint1)

	child1track := make([]tracks.Element, child1len)
	child2track := make([]tracks.Element, child2len)

	copy(child1track[:crossPoint1], parent1.Track[:crossPoint1])
	// if we filled in to child1 from [crossPoint2:] we might have gaps in the
	// track. fill from the end of cross point 1, with the contents of cross
	// point 2.
	copy(child1track[crossPoint1:], parent2.Track[crossPoint2:])

	copy(child2track[:crossPoint2], parent2.Track[:crossPoint2])
	copy(child2track[crossPoint2:], parent1.Track[crossPoint1:])
	child1 := &Member{
		Id:    fmt.Sprintf("iter_%s", uuid.New()),
		Track: child1track,
	}
	child2 := &Member{
		Id:    fmt.Sprintf("iter_%s", uuid.New()),
		Track: child2track,
	}
	return child1, child2
}

func (p *Pool) Spawn(numParents int) *Pool {
	return nil
}
