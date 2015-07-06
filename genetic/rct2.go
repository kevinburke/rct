// RCT2 specific track stuff goes in here
package genetic

import (
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

// Given two vectors that exist in the same 2D plane, return a list of track
// pieces that can be used to connect them.
func connect2DTrackPieces(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	return []tracks.Element{}
}

// Returns track pieces to get the trackEnd to the same vertical height as the
// station start.
func descendToLevel(trackEnd geo.Vector, stationStart geo.Vector) ([]tracks.Element, geo.Vector) {
	return []tracks.Element{}, geo.Vector{geo.Point{0, 0, 0}, 0}
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
