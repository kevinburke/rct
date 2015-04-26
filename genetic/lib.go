// Code for genetic algorithms
package genetic

import (
	"math/rand"

	"github.com/kevinburke/rct-rides/tracks"
)

// constants which may be altered to affect the ride runtime
const PARENTS = 3
const FIT_PERCENTAGE = 0.2
const MUTATION_RATE = 0.05
const POOL_SIZE = 500
const ITERATIONS = 500

func Run() {
	pool := CreatePool(POOL_SIZE)
	for i := 0; i < ITERATIONS; i++ {
		pool.Mutate(MUTATION_RATE)
		pool.Crossover()
		fittest := pool.Fittest(FIT_PERCENTAGE)
		pool = fittest.Spawn(PARENTS)
	}
}

type Pool struct {
	Members []Member
}

type Member struct {
	Track []tracks.Element
	Score int64
	// Advantage or disadvantage in reproducing
	Fitness float64
}

// Create an initial pool
func CreatePool(size int) *Pool {
	// 1. Create a station of length 10
	// 2. For the start station piece, generate a list of possible pieces.
	// 3. Choose one at random. Advance a pointer one forward.
	// 4. Repeat for 50 pieces (Woodchip is length 108. Mischief is 123)
	members := make([]Member, POOL_SIZE)
	for i := 0; i < POOL_SIZE; i++ {
		track := CreateStation()
		idx := STATION_LENGTH - 1
		for j := 0; j < INITIAL_TRACK_LENGTH-STATION_LENGTH; j++ {
			poss := track[idx].Possibilities()
			track = append(track, poss[rand.Intn(len(poss))])
		}
		members[i] = Member{Track: track, Score: 0}
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
		member := p.Members[i]
		member.Score = GetScore(member.Track)
	}

	// Assign fitness for every member.
}

// Select chooses a member of the population at random
func (p *Pool) Select() Member {
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
			return p.Members[index]
		}
	}
	return Member{}
}

// for starters, we will use one point crossover.
func (p *Pool) Crossover() {
	// select 2 parents at random
	// parent1 := p.Select()
	// parent2 := p.Select()
	// crossover with a probability of 0.6 (taken from the book & De Jong 1975)
	// if crossover:
	//	choose a random point between the beginning and the end
	//  while not compatible:
	//		increment one space on track A
	//		check compatibility
	//		increment one space on track B
	//	swap the track pieces at the chosen point on track A and track B
	// return both children
	// (repeat how many times?)
}

func (p *Pool) Fittest(fitPercentage float64) *Pool {
	return nil
}

func (p *Pool) Spawn(numParents int) *Pool {
	return nil
}
