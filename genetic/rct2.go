// RCT2 specific track stuff goes in here
package genetic

import (
	"log"

	"github.com/kevinburke/rct/geo"
	"github.com/kevinburke/rct/physics"
	"github.com/kevinburke/rct/tracks"
)

const STATION_LENGTH = 5
const INITIAL_TRACK_LENGTH = 50

// Create a station of length 10
func CreateStation() []tracks.Element {
	station := make([]tracks.Element, STATION_LENGTH)
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

type scoreData struct {
	Collisions    int
	Distance      int
	NegativeSpeed int
}

// GetScore determines how "good" a track is. Should take into account:
//
// - How close the track is to forming a complete loop
// - How many collisions exist in the track
// - Whether the car can make it all the way around the track
//
// As well as other components that mark a track as "interesting".
func GetScore(t []tracks.Element) (int64, scoreData) {
	if len(t) == 0 {
		return 0, scoreData{}
	}
	vectors := geo.Vectors(t)
	stationStart := geo.Vector{geo.Point{0, 0, 0}, 0}
	trackEnd := vectors[len(t)-1]
	if trackEnd.Dir >= 350 {
		log.Panic("trackend is too high", trackEnd.Dir)
	}
	trackPieces := CompleteTrack(t[len(t)-1], trackEnd, stationStart)
	startingScore := int64(2700 * 1000)

	completedTrack := append(t, trackPieces...)

	data := &tracks.Data{
		Elements:           completedTrack,
		Clearance:          2,
		ClearanceDirection: tracks.CLEARANCE_ABOVE,
	}
	numCollisions := geo.NumCollisions(data)
	numNegativeSpeed := physics.NumNegativeSpeed(data)
	return startingScore - 8000*int64(len(trackPieces)) - 10000*int64(numCollisions) - 5000*int64(numNegativeSpeed), scoreData{
		Collisions:    numCollisions,
		Distance:      len(trackPieces),
		NegativeSpeed: numNegativeSpeed,
	}
}
