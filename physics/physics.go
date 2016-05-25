package physics

import (
	"log"

	"github.com/kevinburke/rct/tracks"
)

func max(a float64, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

const CHAIN_SPEED_MPH = 6

// Map of elevation changes to speed changes. Dumbest possible thing is to make
// speed a function of the change in elevation.
//
// result values are in change in MPH. Came from observing in-game behavior
// with 2-car mine trains.
var accelerationMap = map[int]float64{
	8: 7,
	// hack: geometric average between 1.4 and 7
	4: 2.37,

	2: 1.4,
	1: 0.7,
	0: -0.1,

	// negative values have slightly higher friction coefficients. the game
	// does something fancy I think to mitigate this/ make it always possible
	// to climb within 2 heights of your last height. instead, do the dumb
	// thing
	-1: -0.75,
	-2: -1.45,
	-4: -2.40,
	-8: -7,
}

func NewSpeed(currentSpeed float64, e tracks.Element) float64 {
	if e.ChainLift == true || e.Station == true {
		return max(currentSpeed, CHAIN_SPEED_MPH)
	}
	accel, ok := accelerationMap[e.Segment.ElevationDelta]
	if !ok {
		log.Panicf("no acceleration available for element: %#v", e)
	}
	// XXX handle brakes; currently blacklisted in segmentfns
	return currentSpeed + accel
}

// NumNegativeSpeed returns the number of times the car speed turns negative
func NumNegativeSpeed(t *tracks.Data) int {
	speed := float64(6)
	numNegativeSpeed := 0
	for i := range t.Elements {
		e := t.Elements[i]
		speed = NewSpeed(speed, e)
		if speed < 0 {
			numNegativeSpeed++
			// Reset the speed so the car keeps going forward.
			speed = float64(6)
		}
	}
	return numNegativeSpeed
}
