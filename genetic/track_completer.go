package genetic

import (
	"log"
	"math"

	"github.com/kevinburke/rct/geo"
	"github.com/kevinburke/rct/tracks"
)

func rightTurn(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {

	elem := tracks.Element{
		Segment: tracks.TS_MAP[tracks.ELEM_RIGHT_QUARTER_TURN_3_TILES],
	}
	v := geo.AdvanceVector(trackEnd, elem.Segment)
	return append([]tracks.Element{elem}, connect2DTrackPieces(v, stationStart)...)
}

func leftTurn(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {

	elem := tracks.Element{
		Segment: tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_3_TILES],
	}
	v := geo.AdvanceVector(trackEnd, elem.Segment)
	return append([]tracks.Element{elem}, connect2DTrackPieces(v, stationStart)...)
}

func straight(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	elem := tracks.Element{
		Segment: tracks.TS_MAP[tracks.ELEM_FLAT],
	}
	v := geo.AdvanceVector(trackEnd, elem.Segment)
	return append([]tracks.Element{elem}, connect2DTrackPieces(v, stationStart)...)
}

func connect2DOppositeDirTrackPieces(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {

	if trackEnd.Point[1] >= stationStart.Point[1] {
		if trackEnd.Point[1] >= stationStart.Point[1]+3 {
			if trackEnd.Point[0] >= stationStart.Point[0] {

				elems := make([]tracks.Element, round(trackEnd.Point[0]-stationStart.Point[0]+1))
				count := 0
				v := trackEnd
				for i := trackEnd.Point[0]; i >= stationStart.Point[0]; i-- {
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

			return rightTurn(trackEnd, stationStart)
		}
	} else {
		if trackEnd.Point[1] < stationStart.Point[1]-3 {
			if trackEnd.Point[0] >= stationStart.Point[0] {

				elems := make([]tracks.Element, round(trackEnd.Point[0]-stationStart.Point[0]+1))
				v := trackEnd
				count := 0
				for i := trackEnd.Point[0]; i >= stationStart.Point[0]; i-- {
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

	if stationStart.Dir != tracks.DIR_STRAIGHT {
		log.Panic("stationStart direction is not straight: ", stationStart)
	}

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
		} else {

			return rightTurn(trackEnd, stationStart)
		}
	} else if trackEnd.Point[1] < stationStart.Point[1] {

		if trackEnd.Point[0] < stationStart.Point[0]-5 && trackEnd.Point[1] < stationStart.Point[1]-5 {

			return leftTurn(trackEnd, stationStart)
		} else {
			return rightTurn(trackEnd, stationStart)
		}
	} else {

		if trackEnd.Point[0] > stationStart.Point[0]-5 {

			if trackEnd.Point[1] < stationStart.Point[1]-5 {

				return leftTurn(trackEnd, stationStart)
			} else if trackEnd.Point[1] > stationStart.Point[1]+5 {

				return rightTurn(trackEnd, stationStart)
			} else if trackEnd.Point[1] <= stationStart.Point[1] {

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

	if -1 <= trackEnd.Point[1] && trackEnd.Point[1] < 1 {
		return straight(trackEnd, stationStart)
	}

	if trackEnd.Point[1] <= -2 {

		if trackEnd.Point[0] <= stationStart.Point[0]-6 {

			return leftTurn(trackEnd, stationStart)
		} else {

			return rightTurn(trackEnd, stationStart)
		}
	}

	if trackEnd.Point[0] <= stationStart.Point[0]-2 {

		if trackEnd.Point[1] == stationStart.Point[1]+1 {

			return leftTurn(trackEnd, stationStart)
		} else {

			return straight(trackEnd, stationStart)
		}
	} else {

		if trackEnd.Point[1] >= stationStart.Point[1]+4 {

			return rightTurn(trackEnd, stationStart)
		} else {

			return straight(trackEnd, stationStart)
		}
	}
}

// trackEnd is facing up
func connect2DLeftFacingTrackPieces(trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {

	if -1 < trackEnd.Point[1] && trackEnd.Point[1] <= 1 {
		return straight(trackEnd, stationStart)
	}

	if trackEnd.Point[1] >= 2 {

		if trackEnd.Point[0] <= stationStart.Point[0]-6 {

			return rightTurn(trackEnd, stationStart)
		} else {

			return leftTurn(trackEnd, stationStart)
		}
	}

	if trackEnd.Point[0] <= stationStart.Point[0]-2 {

		if trackEnd.Point[1] == stationStart.Point[1]-1 {

			return rightTurn(trackEnd, stationStart)
		} else {

			return straight(trackEnd, stationStart)
		}
	} else {

		if trackEnd.Point[1] <= stationStart.Point[1]-3 {

			return leftTurn(trackEnd, stationStart)
		} else {

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

// Add pieces to the track to remove the slope and the bank
func straightenTrack(trackPiece tracks.Element, trackEnd geo.Vector) ([]tracks.Element, geo.Vector) {
	if trackPiece.Segment.EndingBank == tracks.TRACK_BANK_NONE &&
		trackPiece.Segment.OutputDegree == tracks.TRACK_NONE {
		return []tracks.Element{}, trackEnd
	}

	if trackPiece.Segment.OutputDegree == tracks.TRACK_UP_60 {
		elem := tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_60_DEG_UP_TO_25_DEG_UP],
		}
		v := geo.AdvanceVector(trackEnd, elem.Segment)
		result, finalV := straightenTrack(elem, v)
		return append([]tracks.Element{elem}, result...), finalV
	} else if trackPiece.Segment.OutputDegree == tracks.TRACK_UP_25 {
		elem := tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_25_DEG_UP_TO_FLAT],
		}
		v := geo.AdvanceVector(trackEnd, elem.Segment)
		result, finalV := straightenTrack(elem, v)
		return append([]tracks.Element{elem}, result...), finalV
	} else if trackPiece.Segment.OutputDegree == tracks.TRACK_DOWN_25 {
		elem := tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN_TO_FLAT],
		}
		v := geo.AdvanceVector(trackEnd, elem.Segment)
		result, finalV := straightenTrack(elem, v)
		return append([]tracks.Element{elem}, result...), finalV
	} else if trackPiece.Segment.OutputDegree == tracks.TRACK_DOWN_60 {
		elem := tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_60_DEG_DOWN_TO_25_DEG_DOWN],
		}
		v := geo.AdvanceVector(trackEnd, elem.Segment)
		result, finalV := straightenTrack(elem, v)
		return append([]tracks.Element{elem}, result...), finalV
	}

	if trackPiece.Segment.EndingBank == tracks.TRACK_BANK_LEFT {
		elem := tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_LEFT_BANK_TO_FLAT],
		}
		v := geo.AdvanceVector(trackEnd, elem.Segment)
		result, finalV := straightenTrack(elem, v)
		return append([]tracks.Element{elem}, result...), finalV
	} else if trackPiece.Segment.EndingBank == tracks.TRACK_BANK_RIGHT {
		elem := tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_RIGHT_BANK_TO_FLAT],
		}
		v := geo.AdvanceVector(trackEnd, elem.Segment)
		result, finalV := straightenTrack(elem, v)
		return append([]tracks.Element{elem}, result...), finalV
	}

	panic("unknown output degree")
}

// Returns track pieces to get the trackEnd to the same vertical height as the
// station start.
func descendToLevel(trackEnd geo.Vector, stationStart geo.Vector) ([]tracks.Element, geo.Vector) {
	if trackEnd.Point[2] == stationStart.Point[2] {

		return []tracks.Element{}, trackEnd
	}
	if trackEnd.Point[2] < stationStart.Point[2] {

		elevation := int(stationStart.Point[2] - trackEnd.Point[2])

		elems := make([]tracks.Element, elevation/2+1)

		elems[0] = tracks.Element{
			ChainLift: true,
			Segment:   tracks.TS_MAP[tracks.ELEM_FLAT_TO_25_DEG_UP],
		}
		for i := 1; i < elevation/2; i++ {
			elems[i] = tracks.Element{
				ChainLift: true,
				Segment:   tracks.TS_MAP[tracks.ELEM_25_DEG_UP],
			}
		}
		elems[elevation/2] = tracks.Element{
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
		elems := make([]tracks.Element, elevation/2+1)

		elems[0] = tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_FLAT_TO_25_DEG_DOWN],
		}
		for i := 1; i < elevation/2; i++ {
			elems[i] = tracks.Element{
				Segment: tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN],
			}
		}
		elems[elevation/2] = tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN_TO_FLAT],
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
func CompleteTrack(trackPiece tracks.Element, trackEnd geo.Vector, stationStart geo.Vector) []tracks.Element {
	straighterTrack, v := straightenTrack(trackPiece, trackEnd)
	levelDescend, v := descendToLevel(v, stationStart)
	straighteners := append(straighterTrack, levelDescend...)
	twodtrack := connect2DTrackPieces(v, stationStart)
	return append(straighteners, twodtrack...)
}
