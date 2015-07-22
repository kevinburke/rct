package tracks

import (
	"log"
	"strings"
)

func isPossibleChainLift(val *Segment) bool {
	// XXX, check other pieces that can be chain lifts. in theory you can also
	// place chain lifts on straight pieces of track but this seems silly
	// (maybe a super in-future optimization will allow this)
	return val.InputDegree == TRACK_UP_25 &&
		(val.OutputDegree == TRACK_UP_25 || val.OutputDegree == TRACK_NONE) &&
		val.StartingBank == TRACK_BANK_NONE &&
		val.EndingBank == TRACK_BANK_NONE
}

func (e *Element) diagonal() bool {
	// 8d to ab are straight diagonal pieces, probably a better way to do that
	// check
	return e.Segment.DirectionDelta == DIR_DIAGONAL_RIGHT ||
		e.Segment.DirectionDelta == DIR_DIAGONAL_LEFT ||
		(e.Segment.Type < 0xAC && e.Segment.Type > 0x8C)
}

func (s *Segment) diagonal() bool {
	return s.DirectionDelta == DIR_DIAGONAL_RIGHT ||
		s.DirectionDelta == DIR_DIAGONAL_LEFT ||
		(s.Type < 0xAC && s.Type > 0x8C)
}

func (e *Element) valid() bool {
	// Heartline and mini golf
	if e.Segment.Type >= 0xc5 && e.Segment.Type <= 0xd0 {
		return false
	}
	return true
}

var blacklist = map[SegmentType]bool{
	ELEM_LEFT_TWIST_DOWN_TO_UP:  true,
	ELEM_RIGHT_TWIST_DOWN_TO_UP: true,
	ELEM_LEFT_TWIST_UP_TO_DOWN:  true,
	ELEM_RIGHT_TWIST_UP_TO_DOWN: true,
	ELEM_HALF_LOOP_UP:           true,
	ELEM_HALF_LOOP_DOWN:         true,
	ELEM_LEFT_CORKSCREW_UP:      true,
	ELEM_RIGHT_CORKSCREW_UP:     true,
	ELEM_LEFT_CORKSCREW_DOWN:    true,
	ELEM_RIGHT_CORKSCREW_DOWN:   true,

	ELEM_FLAT_TO_60_DEG_UP:   true,
	ELEM_60_DEG_UP_TO_FLAT:   true,
	ELEM_FLAT_TO_60_DEG_DOWN: true,
	ELEM_60_DEG_DOWN_TO_FLAT: true,

	ELEM_TOWER_BASE:    true,
	ELEM_TOWER_SECTION: true,

	ELEM_FLAT_COVERED:                       true,
	ELEM_25_DEG_UP_COVERED:                  true,
	ELEM_60_DEG_UP_COVERED:                  true,
	ELEM_FLAT_TO_25_DEG_UP_COVERED:          true,
	ELEM_25_DEG_UP_TO_60_DEG_UP_COVERED:     true,
	ELEM_60_DEG_UP_TO_25_DEG_UP_COVERED:     true,
	ELEM_25_DEG_UP_TO_FLAT_COVERED:          true,
	ELEM_25_DEG_DOWN_COVERED:                true,
	ELEM_60_DEG_DOWN_COVERED:                true,
	ELEM_FLAT_TO_25_DEG_DOWN_COVERED:        true,
	ELEM_25_DEG_DOWN_TO_60_DEG_DOWN_COVERED: true,
	ELEM_60_DEG_DOWN_TO_25_DEG_DOWN_COVERED: true,
	ELEM_25_DEG_DOWN_TO_FLAT_COVERED:        true,
	ELEM_LEFT_QUARTER_TURN_5_TILES_COVERED:  true,
	ELEM_RIGHT_QUARTER_TURN_5_TILES_COVERED: true,
	ELEM_S_BEND_LEFT_COVERED:                true,
	ELEM_S_BEND_RIGHT_COVERED:               true,
	ELEM_LEFT_QUARTER_TURN_3_TILES_COVERED:  true,
	ELEM_RIGHT_QUARTER_TURN_3_TILES_COVERED: true,

	ELEM_LEFT_QUARTER_TURN_1_TILE_60_DEG_UP:    true,
	ELEM_RIGHT_QUARTER_TURN_1_TILE_60_DEG_UP:   true,
	ELEM_LEFT_QUARTER_TURN_1_TILE_60_DEG_DOWN:  true,
	ELEM_RIGHT_QUARTER_TURN_1_TILE_60_DEG_DOWN: true,

	ELEM_ROTATION_CONTROL_TOGGLE:                 true,
	ELEM_INVERTED_90_DEG_UP_TO_FLAT_QUARTER_LOOP: true,

	ELEM_WATERFALL: true,
	ELEM_RAPIDS:    true,

	ELEM_25_DEG_DOWN_LEFT_BANKED:  true,
	ELEM_25_DEG_DOWN_RIGHT_BANKED: true,

	ELEM_WATER_SPLASH: true,

	ELEM_FLAT_TO_60_DEG_UP_LONG_BASE: true,
	ELEM_60_DEG_UP_TO_FLAT_LONG_BASE: true,

	ELEM_WHIRLPOOL: true,

	ELEM_60_DEG_DOWN_TO_FLAT_LONG_BASE: true,
	ELEM_FLAT_TO_60_DEG_DOWN_LONG_BASE: true,

	ELEM_CABLE_LIFT_HILL:             true,
	ELEM_REVERSE_WHOA_BELLY_SLOPE:    true,
	ELEM_REVERSE_WHOA_BELLY_VERTICAL: true,

	ELEM_90_DEG_UP:                  true,
	ELEM_90_DEG_DOWN:                true,
	ELEM_60_DEG_UP_TO_90_DEG_UP:     true,
	ELEM_90_DEG_DOWN_TO_60_DEG_DOWN: true,
	ELEM_90_DEG_UP_TO_60_DEG_UP:     true,
	ELEM_60_DEG_DOWN_TO_90_DEG_DOWN: true,
	ELEM_BRAKE_FOR_DROP:             true,

	ELEM_LOG_FLUME_REVERSER: true,
	ELEM_SPINNING_TUNNEL:    true,

	ELEM_LEFT_BARREL_ROLL_UP_TO_DOWN:  true,
	ELEM_RIGHT_BARREL_ROLL_UP_TO_DOWN: true,
	ELEM_LEFT_BARREL_ROLL_DOWN_TO_UP:  true,
	ELEM_RIGHT_BARREL_ROLL_DOWN_TO_UP: true,

	ELEM_LEFT_BANK_TO_LEFT_QUARTER_TURN_3_TILES_25_DEG_UP:     true,
	ELEM_RIGHT_BANK_TO_RIGHT_QUARTER_TURN_3_TILES_25_DEG_UP:   true,
	ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_DOWN_TO_LEFT_BANK:   true,
	ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_DOWN_TO_RIGHT_BANK: true,

	ELEM_POWERED_LIFT: true,

	ELEM_LEFT_LARGE_HALF_LOOP_UP:    true,
	ELEM_RIGHT_LARGE_HALF_LOOP_UP:   true,
	ELEM_RIGHT_LARGE_HALF_LOOP_DOWN: true,
	ELEM_LEFT_LARGE_HALF_LOOP_DOWN:  true,

	ELEM_LEFT_FLYER_TWIST_UP_TO_DOWN:    true,
	ELEM_RIGHT_FLYER_TWIST_UP_TO_DOWN:   true,
	ELEM_LEFT_FLYER_TWIST_DOWN_TO_UP:    true,
	ELEM_RIGHT_FLYER_TWIST_DOWN_TO_UP:   true,
	ELEM_FLYER_HALF_LOOP_UP:             true,
	ELEM_FLYER_HALF_LOOP_DOWN:           true,
	ELEM_LEFT_FLY_CORKSCREW_UP_TO_DOWN:  true,
	ELEM_RIGHT_FLY_CORKSCREW_UP_TO_DOWN: true,
	ELEM_LEFT_FLY_CORKSCREW_DOWN_TO_UP:  true,
	ELEM_RIGHT_FLY_CORKSCREW_DOWN_TO_UP: true,
	ELEM_HEARTLINE_TRANSFER_UP:          true,
	ELEM_HEARTLINE_TRANSFER_DOWN:        true,
	ELEM_LEFT_HEARTLINE_ROLL:            true,
	ELEM_RIGHT_HEARTLINE_ROLL:           true,
	ELEM_MINI_GOLF_HOLE_A:               true,
	ELEM_MINI_GOLF_HOLE_B:               true,
	ELEM_MINI_GOLF_HOLE_C:               true,
	ELEM_MINI_GOLF_HOLE_D:               true,
	ELEM_MINI_GOLF_HOLE_E:               true,

	ELEM_INVERTED_FLAT_TO_90_DEG_DOWN_QUARTER_LOOP: true,
	ELEM_90_DEG_UP_QUARTER_LOOP_TO_INVERTED:        true,
	ELEM_QUARTER_LOOP_INVERT_TO_90_DEG_DOWN:        true,
	ELEM_LEFT_CURVED_LIFT_HILL:                     true,
	ELEM_RIGHT_CURVED_LIFT_HILL:                    true,
	ELEM_LEFT_REVERSER:                             true,
	ELEM_RIGHT_REVERSER:                            true,
	ELEM_AIR_THRUST_TOP_CAP:                        true,
	ELEM_AIR_THRUST_VERTICAL_DOWN:                  true,
	ELEM_AIR_THRUST_VERTICAL_DOWN_TO_LEVEL:         true,

	// XXX - would like to allow these at some point but they mess with
	// acceleration
	ELEM_BRAKES:       true,
	ELEM_BLOCK_BRAKES: true,

	ELEM_BANKED_LEFT_QUARTER_TURN_3_TILES_25_DEG_UP:  true,
	ELEM_BANKED_RIGHT_QUARTER_TURN_3_TILES_25_DEG_UP: true,
	ELEM_END_OF_RIDE:                                 true,

	ELEM_BEGIN_STATION:  true,
	ELEM_END_STATION:    true,
	ELEM_MIDDLE_STATION: true,
}

func Valid(val *Segment) bool {
	if _, ok := blacklist[val.Type]; ok {
		return false
	}
	// XXX, figure out how to handle diagonal, different types of coasters
	// etc.
	name := ElementNames[val.Type]
	if strings.Contains(string(name), "DIAG") {
		return false
	}
	if strings.Contains(string(name), "EIGHTH") {
		return false
	}
	if strings.Contains(string(name), "GOLF") {
		return false
	}
	if strings.Contains(string(name), "COVERED") {
		return false
	}
	if strings.Contains(string(name), "BRAKES") {
		return false
	}
	if strings.Contains(string(name), "LOOP") {
		return false
	}
	return true
}

// Possibilities computes all of the possible track pieces that can be built
func (s *Element) Possibilities() (o []Element) {
	for _, val := range TS_MAP {
		if !Valid(val) {
			continue
		}
		if val.InputDegree == s.Segment.OutputDegree && val.StartingBank == s.Segment.EndingBank && !s.diagonal() && !val.diagonal() && s.valid() {
			if val.DirectionDelta != 0 && val.DirectionDelta != 90 && val.DirectionDelta != 180 && val.DirectionDelta != 270 {
				log.Panicf("%#v", val)
			}
			o = append(o, Element{Segment: val, ChainLift: false})
			if isPossibleChainLift(val) {
				o = append(o, Element{Segment: val, ChainLift: true})
			}
			// XXX, cache possibilities, since they're not going to change, or
			// just build this once
		}
	}
	return o
}

// Compatible returns true if the end of the first element is compatible with
// the beginning of the second element.
func Compatible(first Element, second Element) bool {
	return first.Segment.OutputDegree == second.Segment.InputDegree &&
		first.Segment.EndingBank == second.Segment.StartingBank
}
