// RCT2 specific track stuff goes in here
package genetic

import (
	"log"
	"math"

	"github.com/kevinburke/rct/geo"
	"github.com/kevinburke/rct/tracks"
)

const STATION_LENGTH = 10
const INITIAL_TRACK_LENGTH = 60

// Create a station of length 10
func CreateStation() []tracks.Element {
	station := make([]tracks.Element, 10)
	beginning := tracks.TS_MAP[tracks.ELEM_BEGIN_STATION]
	middle := tracks.TS_MAP[tracks.ELEM_MIDDLE_STATION]
	end := tracks.TS_MAP[tracks.ELEM_END_STATION]
	station[0] = tracks.Element{Segment: beginning}
	for i := 1; i < STATION_LENGTH; i++ {
		station[i] = tracks.Element{Segment: middle}
	}
	station[STATION_LENGTH-1] = tracks.Element{Segment: end}
	return station
}

func connect2DOppositeDirTrackPieces(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	return []tracks.Element{}
}

// Round to the nearest integer
func round(f float64) int {
	return int(math.Floor(f + 0.5))
}

// caller is responsible for aligning these
func connect2DSameDirTrackPieces(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	// directionDelta: 0, stationStart: (3, 4). ok straight points are ((x < 3), 4)
	if trackEnd.Dir != stationStart.Dir {
		log.Panic("track end and station start are not the same", trackEnd, stationStart)
	}
	if stationStart.Point[0] == trackEnd.Point[0] && stationStart.Point[1] == trackEnd.Point[1] {
		return []tracks.Element{}
	}
	if stationStart.Dir == tracks.DIR_STRAIGHT {
		if stationStart.Point[1] == trackEnd.Point[1] {
			if trackEnd.Point[0] < stationStart.Point[0] {
				elems := make([]tracks.Element, round(stationStart.Point[0]-trackEnd.Point[0]))
				count := 0
				for i := trackEnd.Point[0]; i < stationStart.Point[0]; i++ {
					elems[count] = tracks.Element{
						Segment: tracks.TS_MAP[tracks.ELEM_FLAT],
					}
					count += 1
				}
				return elems
			}
		}
	} else {
		log.Panic("direction didn't match", stationStart.Dir)
	}
	return []tracks.Element{}
}

// Given two vectors that exist in the same 2D plane, return a list of track
// pieces that can be used to connect them.
func connect2DTrackPieces(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	// Okay! This is not a trivial problem.
	if trackEnd.Dir == stationStart.Dir {
		return connect2DSameDirTrackPieces(trackEnd, stationStart)
	} else if trackEnd.Dir-stationStart.Dir == 180 || stationStart.Dir-trackEnd.Dir == 180 {
		return connect2DOppositeDirTrackPieces(trackEnd, stationStart)
	} else {

	}
	return []tracks.Element{}
}

// Returns track pieces to get the trackEnd to the same vertical height as the
// station start.
func descendToLevel(trackEnd geo.Vector, stationStart geo.Vector) ([]tracks.Element, geo.Vector) {
	if trackEnd.Point[2] == stationStart.Point[2] {
		// Easy case, no ascending/descending
		return []tracks.Element{}, trackEnd
	}
	if trackEnd.Point[2] < stationStart.Point[2] {
		elevation := int(stationStart.Point[2] - trackEnd.Point[2])
		elems := make([]tracks.Element, elevation+2)
		// First point is simple ascender
		elems[0] = tracks.Element{
			ChainLift: true,
			Segment:   tracks.TS_MAP[tracks.ELEM_FLAT_TO_25_DEG_UP],
		}
		for i := 0; i < elevation; i++ {
			elems[i+1] = tracks.Element{
				ChainLift: true,
				Segment:   tracks.TS_MAP[tracks.ELEM_25_DEG_UP],
			}
		}
		elems[elevation-1] = tracks.Element{
			ChainLift: true,
			Segment:   tracks.TS_MAP[tracks.ELEM_25_DEG_UP_TO_FLAT],
		}
		return elems, geo.Vector{geo.Point{
			trackEnd.Point[0],
			trackEnd.Point[1],
			trackEnd.Point[2] + float64(elevation),
		}, trackEnd.Dir}
	} else {
		elevation := round(trackEnd.Point[2] - stationStart.Point[2])
		elems := make([]tracks.Element, elevation+2)
		// First point is simple descender
		elems[0] = tracks.Element{
			ChainLift: true,
			Segment:   tracks.TS_MAP[tracks.ELEM_FLAT_TO_25_DEG_DOWN],
		}
		for i := 0; i < elevation; i++ {
			elems[i+1] = tracks.Element{
				ChainLift: true,
				Segment:   tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN],
			}
		}
		elems[elevation] = tracks.Element{
			ChainLift: true,
			Segment:   tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN_TO_FLAT],
		}
		return elems, geo.Vector{geo.Point{
			trackEnd.Point[0],
			trackEnd.Point[1],
			trackEnd.Point[2] - float64(elevation),
		}, trackEnd.Dir}
	}
}

// completeTrack takes a trackEnd and a stationStart and returns a list of
// track pieces needed to join them.
func completeTrack(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	levelDescend, v := descendToLevel(trackEnd, stationStart)
	twodtrack := connect2DTrackPieces(v, stationStart)
	return append(levelDescend, twodtrack...)
}

// GetScore determines how "good" a track is. Should take into account:
//
// - How close the track is to forming a complete loop
// - How many collisions exist in the track
// - Whether the car can make it all the way around the track
//
// As well as other components that mark a track as "interesting".
func GetScore(t []tracks.Element) int64 {
	if len(t) == 0 {
		return 0
	}
	eΔ, forwardΔ, sidewaysΔ := 0, 0, 0
	direction := tracks.DIR_STRAIGHT
	// XXX, we're probably computing this multiple times.
	for i := range t {
		ts := t[i].Segment
		eΔ, forwardΔ, sidewaysΔ, direction = geo.Advance(
			ts, eΔ, forwardΔ, sidewaysΔ, direction)
	}
	vectors := geo.Vectors(t)
	stationStart := geo.Vector{geo.Point{0, 0, 0}, 0}
	trackEnd := vectors[len(t)-1]
	trackPieces := completeTrack(trackEnd, stationStart)
	return int64(10*1000*1000 - 4000*len(trackPieces))
}
