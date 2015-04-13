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

}

// Assign scores for every member of the pool
func (p *Pool) Evaluate() {
	for i := 0; i < POOL_SIZE; i++ {
		member := p.Members[i]
		member.Score = GetScore(member.Track)
	}
}

func (p *Pool) Crossover() {
}

func (p *Pool) Fittest(fitPercentage float64) *Pool {
	return nil
}

func (p *Pool) Spawn(numParents int) *Pool {
	return nil
}
