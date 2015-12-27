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
	"time"

	"github.com/pborman/uuid"

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
const ITERATIONS = 50
const PRINT_RESULTS_EVERY = 1

// create a directory and ignore "directory exists" errors
func mkdir(name string) error {
	err := os.Mkdir(name, 0755)
	// ugh
	if err != nil && err.(*os.PathError).Err.Error() != "file exists" {
		return err
	}
	return nil
}

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
	err := mkdir(expDir)
	if err != nil {
		return err
	}
	id := fmt.Sprintf("exp_%s", uuid.New())
	expIdDir := path.Join(expDir, id)
	err = mkdir(expIdDir)
	if err != nil {
		return err
	}
	iterationsDir := path.Join(expIdDir, "iterations")
	err = mkdir(iterationsDir)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = packageRoot
	hashb, err := cmd.Output()
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
		err = mkdir(iterationDir)
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
	if iteration%PRINT_RESULTS_EVERY == 0 {
		var scores scoresArray
		var highestScore int64 = -1
		var worstScore int64 = 100 * 1000 * 1000
		bestMember := new(Member)
		for i := 0; i < len(p.Members); i++ {
			scores[i] = p.Members[i].Score
			if p.Members[i].Score > highestScore {
				highestScore = p.Members[i].Score
				bestMember = p.Members[i]
			}
			if p.Members[i].Score < worstScore {
				worstScore = p.Members[i].Score
			}
		}
		middle := len(scores) / 2
		median := (scores[middle] + scores[middle-1]) / 2
		sort.Sort(&scores)
		bestScorer := fmt.Sprintf("\t(collisions: %d, to completion: %d, negative speed points: %d)\n\n",
			bestMember.ScoreData.Collisions, bestMember.ScoreData.Distance,
			bestMember.ScoreData.NegativeSpeed)
		fmt.Printf("Iteration %d: %d members, best member %s has score %d, "+
			"median %d, worst has score %d\n%s",
			iteration, len(p.Members), bestMember.Id, bestMember.Score, median,
			worstScore, bestScorer)
	}
	// XXX, move offline to a goroutine
	for i := 0; i < len(p.Members); i++ {
		pth := path.Join(outputDirectory, "experiments", p.Id,
			"iterations", strconv.Itoa(iteration), fmt.Sprintf("%s.json", p.Members[i].Id))
		err := encode(pth, p.Members[i])
		if err != nil {
			log.Print(err)
		}
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
	for i := 0; i < POOL_SIZE; i++ {
		p.Members[i].Score, p.Members[i].ScoreData = GetScore(p.Members[i].Track)
	}

	// Assign fitness for every member. For now, every member gets a fitness of
	// 1. In the future, consider sorting/giving higher score members a better
	// chance of reproducing.
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
		weightedTotal += p.Members[i].Fitness
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

func crossoverOne(parent1 *Member, parent2 *Member) (*Member, *Member) {
	//	choose a random point between the beginning and the end
	minval := min(len(parent1.Track), len(parent2.Track))
	crossPoint1 := rand.Intn(minval)
	crossPoint2 := crossPoint1
	foundMatch := false
	for {
		if tracks.Compatible(parent1.Track[crossPoint1], parent2.Track[crossPoint2]) {
			foundMatch = true
			break
		}
		crossPoint1++
		if crossPoint1 >= len(parent1.Track) {
			break
		}
		if tracks.Compatible(parent1.Track[crossPoint1], parent2.Track[crossPoint2]) {
			foundMatch = true
			break
		}
		crossPoint2++
		if crossPoint2 >= len(parent2.Track) {
			break
		}
	}
	//	swap the track pieces at the chosen point on track A and track B
	if foundMatch {
		return Swap(parent1, parent2, crossPoint1, crossPoint2)
	}
	return parent1, parent2
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
		if rand.Float64() < CROSSOVER_PROBABILITY {
			// XXX delete parents
			child1, child2 := crossoverOne(parent1, parent2)
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
	child1len := crossPoint1 + (len(parent2.Track) - crossPoint1)
	child2len := crossPoint2 + (len(parent1.Track) - crossPoint2)
	// XXX, probably something fancy you can do with xor, or a temporary array.
	child1track := make([]tracks.Element, child1len)
	child2track := make([]tracks.Element, child2len)
	copy(child1track[:crossPoint1], parent1.Track[:crossPoint1])
	copy(child1track[crossPoint1:], parent2.Track[crossPoint1:])
	copy(child2track[:crossPoint2], parent2.Track[:crossPoint2])
	copy(child2track[crossPoint2:], parent1.Track[crossPoint2:])
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
