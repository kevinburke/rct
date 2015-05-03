// RCT2 specific track stuff goes in here
package genetic

import "github.com/kevinburke/rct-rides/tracks"

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

func Compatible(a tracks.Element, b tracks.Element) {

}

func GetScore(t []tracks.Element) int64 {
	// Incredibly sophisticated scoring algorithm.
	return int64(len(t))
}
