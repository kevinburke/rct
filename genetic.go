// Code for genetic algorithms
package main

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

type Pool struct{}

// Create an initial pool
func CreatePool(size int) *Pool {
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
