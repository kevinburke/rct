package geo

import (
	"log"
	"math"

	"github.com/kevinburke/rct/tracks"
)

type Vector struct {
	Point Point
	Dir   tracks.DirectionDelta
}

// A Point is a representation of (x, y, z) coordinates
type Point [3]float64

func Vectors(elems []tracks.Element) []Vector {
	var direction tracks.DirectionDelta
	direction = 0
	current := Vector{Point{0, 0, 0}, direction}
	plist := []Vector{current}
	for _, elem := range elems {
		segm := tracks.TS_MAP[elem.Segment.Type]
		p := Diff(segm, direction)
		current.Point[0] += p[0]
		current.Point[1] += p[1]
		current.Point[2] += p[2]
		direction += segm.DirectionDelta
		for direction >= 360 {
			direction -= 360
		}
		current.Dir = direction
		plist = append(plist, current)
	}
	return plist
}

func Render(elems []tracks.Element) ([]Vector, []Vector) {
	plist := Vectors(elems)
	return plist, []Vector{}
}

// Round to the nearest integer
func round(f float64) int {
	return int(math.Floor(f + 0.5))
}

// Advance all of the values by one track segment.
//
// Positive forward = right
// Negative forward = left
// Positive sideways = up
// Negative sideways = down

// right = up. left = down.
func Advance(ts *tracks.Segment, ΔE int, ΔForward int, ΔSideways int,
	direction tracks.DirectionDelta) (int, int, int, tracks.DirectionDelta) {
	panic("these functions need to be fixed, check the tests")

	ΔE += ts.ElevationDelta

	fdirection := float64(direction)
	ΔForward += round(cosdeg(fdirection) * float64(ts.ForwardDelta))
	ΔForward += round(sindeg(fdirection) * float64(ts.SidewaysDelta))

	ΔSideways += round(sindeg(fdirection) * float64(ts.ForwardDelta))
	ΔSideways += round(cosdeg(fdirection) * float64(ts.SidewaysDelta))

	direction += ts.DirectionDelta
	for ; direction >= 360; direction -= 360 {
	}

	return ΔE, ΔForward, ΔSideways, direction
}

// sindeg computes sines in degrees
func sindeg(deg float64) float64 {
	for ; deg >= 360; deg -= 360 {
	}
	if round(deg)%180 == 0 {
		return 0
	} else if round(deg) == 90 {
		return 1
	} else if round(deg) == 270 {
		return -1
	} else {
		return math.Sin(deg * math.Pi / 180)
	}
}

// computes sines in degrees
func cosdeg(deg float64) float64 {
	for ; deg >= 360; deg -= 360 {
	}
	if round(deg) == 0 {
		return 1
	} else if round(deg) == 90 || round(deg) == 270 {
		return 0
	} else if round(deg) == 180 {
		return -1
	} else {
		return math.Sin(deg * math.Pi / 180)
	}
}

// PositionChange takes a travel direction in degrees and a track segment and
// returns the distance traveled in the X, Y, and Z directions.
func Diff(ts *tracks.Segment, direction tracks.DirectionDelta) *Point {
	//rotate around the Z axis: http://stackoverflow.com/a/14609567/329700
	x := float64(ts.ForwardDelta)*cosdeg(float64(direction)) - float64(ts.SidewaysDelta)*sindeg(float64(direction))
	y := float64(ts.ForwardDelta)*sindeg(float64(direction)) + float64(ts.SidewaysDelta)*cosdeg(float64(direction))
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
		eΔ, forwardΔ, sidewaysΔ, direction = Advance(
			ts, eΔ, forwardΔ, sidewaysΔ, direction)
	}
	return forwardΔ == 0 && sidewaysΔ == 0 && eΔ == 0
}

func AdvanceVector(v Vector, s *tracks.Segment) Vector {
	var p Point
	if v.Dir == tracks.DIR_STRAIGHT {
		// facing right - (forward, -sideways)
		p = Point{
			v.Point[0] + float64(s.ForwardDelta),
			v.Point[1] - float64(s.SidewaysDelta),
			v.Point[2] + float64(s.ElevationDelta),
		}
	} else if v.Dir == tracks.DIR_90_DEG_RIGHT {
		// facing up - (sideways, forward)
		p = Point{
			v.Point[0] + float64(s.SidewaysDelta),
			v.Point[1] + float64(s.ForwardDelta),
			v.Point[2] + float64(s.ElevationDelta),
		}
	} else if v.Dir == tracks.DIR_180_DEG {
		// facing left - (-forward, sideways)
		p = Point{
			v.Point[0] - float64(s.ForwardDelta),
			v.Point[1] + float64(s.SidewaysDelta),
			v.Point[2] + float64(s.ElevationDelta),
		}
	} else if v.Dir == tracks.DIR_90_DEG_LEFT {
		// facing down - (-sideways, -forward)
		p = Point{
			v.Point[0] - float64(s.SidewaysDelta),
			v.Point[1] - float64(s.ForwardDelta),
			v.Point[2] + float64(s.ElevationDelta),
		}
	} else {
		log.Panic("unsupported direction", v.Dir)
	}
	dir := v.Dir + s.DirectionDelta
	for ; dir >= 360; dir -= 360 {
	}
	return Vector{p, dir}
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
		eΔ, forwardΔ, sidewaysΔ, direction = Advance(
			ts, eΔ, forwardΔ, sidewaysΔ, direction)
		// if there already exists a piece there, we can't build.
		if matrix[forwardΔ][sidewaysΔ][eΔ] {
			return true
		}
		matrix[forwardΔ][sidewaysΔ][eΔ] = true
	}
	return false
}
