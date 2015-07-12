// RCT2 specific track stuff goes in here
package genetic

import (
	"fmt"
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

func rightTurn(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	// make a right turn & recursive call
	fmt.Println("right turn")
	elem := tracks.Element{
		Segment: tracks.TS_MAP[tracks.ELEM_RIGHT_QUARTER_TURN_3_TILES],
	}
	v := geo.AdvanceVector(trackEnd, elem.Segment)
	return append([]tracks.Element{elem}, connect2DTrackPieces(v, stationStart)...)
}

func leftTurn(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	// make a left turn & recursive call
	fmt.Println("left turn")
	elem := tracks.Element{
		Segment: tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_3_TILES],
	}
	v := geo.AdvanceVector(trackEnd, elem.Segment)
	return append([]tracks.Element{elem}, connect2DTrackPieces(v, stationStart)...)
}

func straight(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	fmt.Println("straight")
	elem := tracks.Element{
		Segment: tracks.TS_MAP[tracks.ELEM_FLAT],
	}
	v := geo.AdvanceVector(trackEnd, elem.Segment)
	return append([]tracks.Element{elem}, connect2DTrackPieces(v, stationStart)...)
}

func connect2DOppositeDirTrackPieces(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	// if is really far to the left
	//	if is really far behind start point
	//   advance to station start, recursive call
	//  else
	//	 return X direction turn + recursive call
	// else if really far to the right
	//	repeat in other direction
	// else if left of center
	//  turn toward the outside
	// else (slightly right of center or on center)
	//  turn toward the outside
	//

	// this should mean that the station start is facing to the right
	// (dir_straight) and the track end is facing to the left.

	// if track end is (x, y), ok points are ( (any x), y + 5)
	if trackEnd.Point[1] >= stationStart.Point[1] {
		if trackEnd.Point[1] >= stationStart.Point[1]+4 {
			if trackEnd.Point[0] > stationStart.Point[0] {
				// go straight until you get in line
				elems := make([]tracks.Element, round(trackEnd.Point[0]-stationStart.Point[0]))
				count := 0
				v := trackEnd
				for i := trackEnd.Point[0]; i > stationStart.Point[0]; i-- {
					elems[count] = tracks.Element{
						Segment: tracks.TS_MAP[tracks.ELEM_FLAT],
					}
					v = geo.AdvanceVector(v, elems[count].Segment)
					count += 1
				}
				return append(elems, connect2DTrackPieces(v, stationStart)...)
			} else {
				return leftTurn(trackEnd, stationStart)
			}
		} else {
			// heading the wrong way and too close to the station. make a turn
			// to the outside
			return rightTurn(trackEnd, stationStart)
		}
	} else {
		if trackEnd.Point[1] < stationStart.Point[1]-4 {
			if trackEnd.Point[0] > stationStart.Point[0] {
				// go straight until you get in line
				elems := make([]tracks.Element, round(trackEnd.Point[0]-stationStart.Point[0]))
				v := trackEnd
				count := 0
				for i := trackEnd.Point[0]; i > stationStart.Point[0]; i-- {
					elems[count] = tracks.Element{
						Segment: tracks.TS_MAP[tracks.ELEM_FLAT],
					}
					v = geo.AdvanceVector(v, elems[count].Segment)
					count += 1
				}
				return append(elems, connect2DTrackPieces(v, stationStart)...)
			} else {
				return rightTurn(trackEnd, stationStart)
			}
		} else {
			// heading the wrong way and too close to the station. make a turn
			// to the outside
			return leftTurn(trackEnd, stationStart)
		}
	}
}

// Round to the nearest integer
func round(f float64) int {
	return int(math.Floor(f + 0.5))
}

// caller is responsible for aligning these
func connect2DSameDirTrackPieces(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	if trackEnd.Dir != stationStart.Dir {
		log.Panic("track end and station start are not the same", trackEnd, stationStart)
	}
	if stationStart.Point[0] == trackEnd.Point[0] && stationStart.Point[1] == trackEnd.Point[1] {
		return []tracks.Element{}
	}
	// in theory, since the station start is always created facing straight, it
	// should always be DIR_STRAIGHT
	if stationStart.Dir != tracks.DIR_STRAIGHT {
		log.Panic("stationStart direction is not straight: ", stationStart)
	}
	// directionDelta: 0, stationStart: (3, 4). ok straight points are ((x < 3), 4)
	if stationStart.Point[1] == trackEnd.Point[1] {
		// points have same Y axis
		if trackEnd.Point[0] < stationStart.Point[0] {
			// trackEnd directly behind stationStart. proceed forward to
			// the station!
			elems := make([]tracks.Element, round(stationStart.Point[0]-trackEnd.Point[0]))
			count := 0
			for i := trackEnd.Point[0]; i < stationStart.Point[0]; i++ {
				elems[count] = tracks.Element{
					Segment: tracks.TS_MAP[tracks.ELEM_FLAT],
				}
				count += 1
			}
			return elems
		} else {
			// track end ahead of station start...not great, turn to the
			// right.
			return rightTurn(trackEnd, stationStart)
		}
	} else if trackEnd.Point[1] < stationStart.Point[1] {
		// trackEnd below the station start
		if trackEnd.Point[0] < stationStart.Point[0]-5 && trackEnd.Point[1] < stationStart.Point[1]-5 {
			// Can get there with a left turn and a right turn
			return leftTurn(trackEnd, stationStart)
		} else {
			return rightTurn(trackEnd, stationStart)
		}
	} else {
		// trackEnd is above the station start

		// if the track is behind the station start by enough
		if trackEnd.Point[0] > stationStart.Point[0]-5 {
			// if it's below the station start by enough
			if trackEnd.Point[1] < stationStart.Point[1]-5 {
				// Can get there with a left turn and a right turn
				return leftTurn(trackEnd, stationStart)
			} else if trackEnd.Point[1] > stationStart.Point[1]+5 {
				// Can get there with a right turn and a left turn
				return rightTurn(trackEnd, stationStart)
			} else if trackEnd.Point[1] <= stationStart.Point[1] {
				// below but not by far enough, turn right
				return rightTurn(trackEnd, stationStart)
			} else {
				return leftTurn(trackEnd, stationStart)
			}
		} else {
			if trackEnd.Point[1] < stationStart.Point[1] {
				return rightTurn(trackEnd, stationStart)
			} else {
				return leftTurn(trackEnd, stationStart)
			}
		}
	}
	panic("shouldn't get here")
}

// trackEnd is facing down on conventional grid
func connect2DRightFacingTrackPieces(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	// if you are 8 above:
	//	turn right
	// else if you are 3 below:
	//	turn right
	// else if you are 0-2 below:
	//	go straight
	// else if you are ahead or less than 3 pieces behind:
	//	go straight
	// else if you are 3 pieces behind:
	//	turn left
	// else
	//	go straight
	if trackEnd.Point[1] >= stationStart.Point[1]+8 {
		return rightTurn(trackEnd, stationStart)
	} else if trackEnd.Point[1] <= stationStart.Point[1]-2 {
		return rightTurn(trackEnd, stationStart)
	} else if trackEnd.Point[1] <= stationStart.Point[1] {
		return straight(trackEnd, stationStart)
	} else if trackEnd.Point[0] > stationStart.Point[0]-2 {
		return straight(trackEnd, stationStart)
	} else if trackEnd.Point[0] <= stationStart.Point[0]-2 {
		// golden
		return leftTurn(trackEnd, stationStart)
	} else {
		return straight(trackEnd, stationStart)
	}
}

// trackEnd is facing up
func connect2DLeftFacingTrackPieces(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	// else:

	// if you are 1 above, 1 below, or dead level, only option is to go straight
	if -1 <= trackEnd.Point[1] && trackEnd.Point[1] <= 1 {
		return straight(trackEnd, stationStart)
	}
	// if you are 2 or more above:
	if trackEnd.Point[1] >= 2 {
		// if you are 6 or more behind the station start:
		if trackEnd.Point[0] <= stationStart.Point[0]-6 {
			// right, right, left, back to station
			return rightTurn(trackEnd, stationStart)
		} else {
			// 3 left turns & back to station
			return leftTurn(trackEnd, stationStart)
		}
	}

	// so far we've covered every square in y >= -1
	// if you are 2 behind:
	if trackEnd.Point[0] <= stationStart.Point[0]-2 {
		//	if you are exactly 2 below:
		if trackEnd.Point[1] == stationStart.Point[1]-2 {
			// right turn puts you in line to go straight to station)
			return rightTurn(trackEnd, stationStart)
		} else {
			// will turn right at 2 below & go straight to station
			return straight(trackEnd, stationStart)
		}
	} else { // (you are 2 or more below & in front of station end)
		//	if you are 4 or more below:
		if trackEnd.Point[1] <= stationStart.Point[1]-4 {
			// can get to station with left, right, right
			return leftTurn(trackEnd, stationStart)
		} else {
			// nothing to do but straight & circle around
			return straight(trackEnd, stationStart)
		}
	}

}

// Given two vectors that exist in the same 2D plane, return a list of track
// pieces that can be used to connect them.
func connect2DTrackPieces(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	if trackEnd.Point[0] == stationStart.Point[0] && trackEnd.Point[1] == stationStart.Point[1] && trackEnd.Dir == stationStart.Dir {
		return []tracks.Element{}
	}
	fmt.Printf("from: (%d, %d) to: (%d, %d), initial direction: %v\n", round(trackEnd.Point[0]), round(trackEnd.Point[1]), round(stationStart.Point[0]), round(stationStart.Point[1]), trackEnd.Dir)
	// Okay! This is not a trivial problem. Slightly easier if you remember
	// that the stations are always created facing to the left.
	if trackEnd.Dir == stationStart.Dir {
		return connect2DSameDirTrackPieces(trackEnd, stationStart)
	} else if trackEnd.Dir-stationStart.Dir == 180 || stationStart.Dir-trackEnd.Dir == 180 {
		return connect2DOppositeDirTrackPieces(trackEnd, stationStart)
	} else if stationStart.Dir == tracks.DIR_STRAIGHT && trackEnd.Dir == tracks.DIR_90_DEG_RIGHT {
		return connect2DRightFacingTrackPieces(trackEnd, stationStart)
	} else if stationStart.Dir == tracks.DIR_STRAIGHT && trackEnd.Dir == tracks.DIR_90_DEG_LEFT {
		return connect2DLeftFacingTrackPieces(trackEnd, stationStart)
	} else {
		log.Panicf("trackend dir: %s, stationstart dir: %s", trackEnd.Dir, stationStart.Dir)
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
		// Ascend to station
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
		// Descend to station
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
	fmt.Println("done!")
	fmt.Printf("track length is %d\n", len(twodtrack))
	result := append(levelDescend, twodtrack...)
	return result
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
	vectors := geo.Vectors(t)
	stationStart := geo.Vector{geo.Point{0, 0, 0}, 0}
	trackEnd := vectors[len(t)-1]
	if trackEnd.Dir >= 350 {
		log.Panic("trackend is too high", trackEnd.Dir)
	}
	trackPieces := completeTrack(trackEnd, stationStart)
	return int64(10*1000*1000 - 4000*len(trackPieces))
}
