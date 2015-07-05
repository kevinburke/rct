// RCT2 specific track stuff goes in here
package genetic

import "github.com/kevinburke/rct/tracks"

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

func distanceFromStart() {

}

// GetScore determines how "good" a track is. Should take into account:
//
// - How close the track is to forming a complete loop
// - How many collisions exist in the track
// - Whether the car can make it all the way around the track
//
// As well as other components that mark a track as "interesting"
func GetScore(t []tracks.Element) int64 {
	return int64(len(t))
}
