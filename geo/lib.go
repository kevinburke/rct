package geo

import (
	"math"

	"github.com/kevinburke/rct/tracks"
)

// A Point is a representation of (x, y, z) coordinates
type Point [3]float64

func Render(elems []tracks.Element) ([]*Point, []*Point) {
	left := []*Point{&Point{0, 0, 0}}
	direction := 0
	for _, elem := range elems {
		p := Diff(elem.Segment, direction)
		left = append(left, p)
		direction += int(elem.Segment.DirectionDelta)
	}
	return left, []*Point{}
}

// Round to the nearest integer
func round(f float64) int {
	return int(math.Floor(f + 0.5))
}

// Advance all of the values by one track segment.
func advanceTrack(ts *tracks.Segment, ΔE int, ΔForward int, ΔSideways int,
	direction tracks.DirectionDelta) (int, int, int, tracks.DirectionDelta) {

	// XXX
	ΔE += 0

	ΔForward += round(cosdeg(int(direction)) * float64(ts.ForwardDelta))
	ΔForward += round(sindeg(int(direction)) * float64(ts.SidewaysDelta))

	ΔSideways += round(sindeg(int(direction)) * float64(ts.ForwardDelta))
	ΔSideways += round(cosdeg(int(direction)) * float64(ts.SidewaysDelta))

	direction += ts.DirectionDelta
	for ; direction >= 360; direction -= 360 {
	}

	return ΔE, ΔForward, ΔSideways, direction
}

// sindeg computes sines in degrees
func sindeg(deg int) float64 {
	for ; deg >= 360; deg -= 360 {
	}
	if deg%180 == 0 {
		return 0
	} else if deg == 90 {
		return 1
	} else if deg == 270 {
		return -1
	} else {
		return math.Sin(float64(deg) * math.Pi / 180)
	}
}

// computes sines in degrees
func cosdeg(deg int) float64 {
	for ; deg >= 360; deg -= 360 {
	}
	if deg == 0 {
		return 1
	} else if deg == 90 || deg == 270 {
		return 0
	} else if deg == 180 {
		return -1
	} else {
		return math.Sin(float64(deg) * math.Pi / 180)
	}
}

// PositionChange takes a travel direction in degrees and a track segment and
// returns the distance traveled in the X, Y, and Z directions.
func Diff(ts *tracks.Segment, direction int) *Point {
	//rotate around the Z axis: http://stackoverflow.com/a/14609567/329700
	x := float64(ts.ForwardDelta)*cosdeg(direction) - float64(ts.SidewaysDelta)*sindeg(direction)
	y := float64(ts.ForwardDelta)*sindeg(direction) + float64(ts.SidewaysDelta)*cosdeg(direction)
	return &Point{x, y, float64(ts.ElevationDelta)}
}

// Test whether the ride's track forms a continuous circuit. Does not test
// whether the ride collides with itself.
func IsCircuit(t *tracks.Data) bool {
	// X and Y don't really make sense as variable names, easier to think about
	// relative changes
	eΔ, forwardΔ, sidewaysΔ := 0, 0, 0
	direction := tracks.DIR_STRAIGHT
	if len(t.Elements) == 0 {
		return false
	}
	for i := range t.Elements {
		ts := t.Elements[i].Segment
		eΔ, forwardΔ, sidewaysΔ, direction = advanceTrack(
			ts, eΔ, forwardΔ, sidewaysΔ, direction)
	}
	return forwardΔ == 0 && sidewaysΔ == 0 && eΔ == 0
}

// Detect whether the track collides with itself.
func HasCollision(t *tracks.Data) bool {
	matrix := make([][][]bool, 100)
	for i := range matrix {
		matrix[i] = make([][]bool, 100)
		for j := range matrix[i] {
			matrix[i][j] = make([]bool, 100)
		}
	}
	eΔ, forwardΔ, sidewaysΔ := 0, 0, 0
	direction := tracks.DIR_STRAIGHT
	for i := range t.Elements {
		ts := t.Elements[i].Segment
		eΔ, forwardΔ, sidewaysΔ, direction = advanceTrack(
			ts, eΔ, forwardΔ, sidewaysΔ, direction)
		// if there already exists a piece there, we can't build.
		if matrix[forwardΔ][sidewaysΔ][eΔ] {
			return true
		}
		matrix[forwardΔ][sidewaysΔ][eΔ] = true
	}
	return false
}
