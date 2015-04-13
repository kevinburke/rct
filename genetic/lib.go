// Code for genetic algorithms
package genetic

import "github.com/kevinburke/rct-rides/tracks"

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
	Tracks [][]tracks.Element
}

// Create an initial pool
func CreatePool(size int) *Pool {
	// 1. Create a station of length 10
	// 2. For the start station piece, generate a list of possible pieces.
	// 3. Choose one at random. Advance a pointer one forward.
	// 4. Repeat for 50 pieces (Woodchip is length 108. Mischief is 123)
	for i := 0; i < POOL_SIZE; i++ {
		track := CreateStation()
	}
	return nil
}

func (p *Pool) Mutate(rate float64) {

}

func (p *Pool) Crossover() {

}

func (p *Pool) Fittest(fitPercentage float64) *Pool {
	return nil
}

func (p *Pool) Spawn(numParents int) *Pool {
	return nil
}
