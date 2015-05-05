// Code for genetic algorithms
package genetic

import (
	"fmt"
	"math/rand"
	"os"
	"path"

	"github.com/kevinburke/rct/tracks"
)

// constants which may be altered to affect the ride runtime
const PARENTS = 2
const FIT_PERCENTAGE = 0.2
const MUTATION_RATE = 0.05

// crossover with a probability of 0.6 (taken from the book & De Jong 1975)
const CROSSOVER_PROBABILITY = 0.6
const POOL_SIZE = 500
const ITERATIONS = 200
const PRINT_RESULTS_EVERY = 1

func Run(outputDirectory string) error {
	pool := CreatePool(POOL_SIZE)
	expDir := path.Join(outputDirectory, "experiments")
	err := os.Mkdir(expDir, 0755)
	// ugh
	if err != nil && err.(*os.PathError).Err.Error() != "file exists" {
		return err
	}
	for i := 0; i < ITERATIONS; i++ {
		pool = pool.Crossover()
		pool.Mutate(MUTATION_RATE)
		pool.Evaluate()
		pool.Statistics(i)
	}
	return nil
}

type Pool struct {
	Members []*Member
}

type Member struct {
	Track []tracks.Element
	Score int64
	// Advantage or disadvantage in reproducing
	Fitness float64
}

func (p *Pool) Statistics(iteration int) {
	if iteration%PRINT_RESULTS_EVERY == 0 {
		var highestScore int64 = -1
		var worstScore int64 = 7000
		bestMember := new(Member)
		for i := 0; i < len(p.Members); i++ {
			if p.Members[i].Score > highestScore {
				highestScore = p.Members[i].Score
				bestMember = p.Members[i]
			}
			if p.Members[i].Score < worstScore {
				worstScore = p.Members[i].Score
			}
		}
		fmt.Printf("Iteration %d: %d members, best member has score %d, worst has score %d\n", iteration, len(p.Members), bestMember.Score, worstScore)
	}
}

// Create an initial pool
func CreatePool(size int) *Pool {
	// 1. Create a station of length 10
	// 2. For the start station piece, generate a list of possible pieces.
	// 3. Choose one at random. Advance a pointer one forward.
	// 4. Repeat for 50 pieces (Woodchip is length 108. Mischief is 123)
	members := make([]*Member, POOL_SIZE)
	for i := 0; i < POOL_SIZE; i++ {
		track := CreateStation()
		idx := STATION_LENGTH - 1
		for j := 0; j < INITIAL_TRACK_LENGTH-STATION_LENGTH; j++ {
			poss := track[idx].Possibilities()
			track = append(track, poss[rand.Intn(len(poss))])
		}
		score := GetScore(track)
		members[i] = &Member{Track: track, Score: score}
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
		p.Members[i].Score = GetScore(p.Members[i].Track)
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
	child2len := crossPoint2 + (len(parent1.Track) - crossPoint1)
	// XXX, probably something fancy you can do with xor, or a temporary array.
	child1track := make([]tracks.Element, child1len)
	child2track := make([]tracks.Element, child2len)
	copy(child1track[:crossPoint1], parent1.Track[:crossPoint1])
	copy(child1track[crossPoint1:], parent2.Track[crossPoint1:])
	copy(child2track[:crossPoint2], parent2.Track[:crossPoint2])
	copy(child2track[crossPoint2:], parent1.Track[crossPoint2:])
	child1 := &Member{Track: child1track}
	child2 := &Member{Track: child2track}
	return child1, child2
}

func (p *Pool) Spawn(numParents int) *Pool {
	return nil
}
