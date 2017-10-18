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

var matrix [100][100][300]bool

// From OpenRCT2: https://github.com/OpenRCT2/OpenRCT2/blob/db438a27b79d3c29b5f7f9100995612e69f798dd/test/testpaint/compat.c#L41
var TILE_DIRECTION_DELTA = [8][2]float64{
	{-1,   0},
	{0,    1},
	{1,    0},
	{0,   -1},
	{-1,  1},
	{1,   1},
	{1,  -1},
	{-1, -1},
};

func Vectors(elems []tracks.Element) []Vector {
	var direction tracks.DirectionDelta
	direction = 0
	current := Vector{Point{0, 0, 0}, direction}
	plist := []Vector{current}
	for _, elem := range elems {
		segm := tracks.TS_MAP[elem.Segment.Type]
		newVector := AdvanceVector(current, segm)
		current.Point[0] = newVector.Point[0]
		current.Point[1] = newVector.Point[1]
		current.Point[2] = newVector.Point[2]
		current.Dir = newVector.Dir
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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Advance all of the values by one track segment.
//
// Positive forward = right
// Negative forward = left
// Positive sideways = up
// Negative sideways = down

// right = up. left = down.

// DEPRECATED
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
		// facing right - (forward, sideways)
		p = Point{
			v.Point[0] + float64(s.ForwardDelta),
			v.Point[1] + float64(s.SidewaysDelta),
			v.Point[2] + float64(s.ElevationDelta),
		}
	} else if v.Dir == tracks.DIR_90_DEG_LEFT {
		// facing up - (-sideways, forward)
		p = Point{
			v.Point[0] - float64(s.SidewaysDelta),
			v.Point[1] + float64(s.ForwardDelta),
			v.Point[2] + float64(s.ElevationDelta),
		}
	} else if v.Dir == tracks.DIR_180_DEG {
		// facing left - (-forward, -sideways)
		p = Point{
			v.Point[0] - float64(s.ForwardDelta),
			v.Point[1] - float64(s.SidewaysDelta),
			v.Point[2] + float64(s.ElevationDelta),
		}
	} else if v.Dir == tracks.DIR_90_DEG_RIGHT {
		// facing down - (sideways, -forward)
		p = Point{
			v.Point[0] + float64(s.SidewaysDelta),
			v.Point[1] - float64(s.ForwardDelta),
			v.Point[2] + float64(s.ElevationDelta),
		}
	} else {
		log.Panic("unsupported direction", v.Dir)
	}
	dir := v.Dir + s.DirectionDelta
	for ; dir >= 360; dir -= 360 {
	}

	if s.DirectionDelta > 0 {
		// Based on OpenRCT2 algorithm: https://github.com/OpenRCT2/OpenRCT2/blob/46de90df8699bc74f1da15fdb16c4a0e2ac1e435/src/openrct2/windows/track_place.c#L547
		// +4 to prevent negative numbers
		// +1 is a magic number
		var dirIndex = int((dir + (dir - v.Dir)) / 90 + 5) % 4

		// Coordinate system here is flipped from OpenRCT2
		p[0] += TILE_DIRECTION_DELTA[dirIndex][1]
		p[1] += TILE_DIRECTION_DELTA[dirIndex][0]
	}

	return Vector{p, dir}
}

// XYSpaceRequired determines how big of a footprint the track has
func XYSpaceRequired(elems []tracks.Element) (int, int) {
	minX, minY, maxX, maxY := float64(0), float64(0), float64(0), float64(0)
	v := Vector{Point{0, 0, 0}, tracks.DIR_STRAIGHT}
	for _, elem := range elems {
		v = AdvanceVector(v, elem.Segment)
		if v.Point[0] < minX {
			minX = v.Point[0]
		}
		if v.Point[0] > maxX {
			maxX = v.Point[0]
		}
		if v.Point[1] < minY {
			minY = v.Point[1]
		}
		if v.Point[1] > maxY {
			maxY = v.Point[1]
		}
	}
	return round(maxX - minX), round(maxY - minY)
}

// reset the matrix, XXX add some kind of lock mechanism here...
func reset() {
	for i := range matrix {
		for j := range matrix[i] {
			for k := range matrix[i][j] {
				matrix[i][j][k] = false
			}
		}
	}
}

// NumCollisions finds the number of collisions the track has with itself.
//
// This isn't perfect - it doesn't nail all of the clearances - but it'll get
// us close enough.
func NumCollisions(t *tracks.Data) int {
	reset()
	v := Vector{Point{50, 50, 150}, tracks.DIR_STRAIGHT}
	closestX, closestY, closestZ := 0, 0, 0
	oldX, oldY, oldZ := 0, 0, 0
	count := 0
	for i := range t.Elements {
		ts := t.Elements[i].Segment
		oldX = closestX
		oldY = closestY
		oldZ = closestZ
		v = AdvanceVector(v, ts)
		closestX = round(v.Point[0])
		closestY = round(v.Point[1])
		closestZ = round(v.Point[2])

		// XXX - take helixes into account
		lowX := min(oldX, closestX)
		highX := max(oldX, closestX)
		lowY := min(oldY, closestY)
		highY := max(oldY, closestY)
		lowZ := min(oldZ, closestZ)
		highZ := max(oldZ, closestZ)

		// Too wide of a track is a "collision"
		if lowX < 0 || highX >= 100 {
			count++
			continue
		}
		if lowY < 0 || highY >= 100 {
			count++
			continue
		}
		// XXX: Car can only go 20 pieces above the ground
		if lowZ < 0 || highZ+2 >= 300 {
			count++
			continue
		}

		// fill in in between points. assume everything is at ending elevation
		// - not great, meh. probably need a third loop
		// Ugh, this might not fill properly based on high/low ordering. might
		// be ok since it's ok to fill space more than once
		collision := false
		for i := lowX; i < highX && !collision; i++ {
			for j := lowY; j < highY && !collision; j++ {
				for k := lowZ; k < highZ+2 && !collision; k++ {
					//fmt.Printf("%d %d %d\n", i, j, k)
					if matrix[i][j][k] {
						count++
						// only count one track piece once.
						collision = true
						break
					}
					matrix[i][j][k] = true
				}
			}
		}
	}
	return count
}
