package tracks

import "strings"

type Degree int

const (
	TRACK_NONE    Degree = 0
	TRACK_UP_25          = 1
	TRACK_UP_60          = 4
	TRACK_DOWN_25        = -1
	TRACK_DOWN_60        = -4
	TRACK_UP_90          = 10
	TRACK_DOWN_90        = -10
)

type Bank int

const (
	TRACK_BANK_NONE Bank = iota
	TRACK_BANK_LEFT
	TRACK_BANK_RIGHT
	TRACK_BANK_FLAT
	TRACK_BANK_UPSIDE_DOWN
)

type DirectionDelta int

const (
	// All of these expressed in positives to help with degree calculations
	DIR_STRAIGHT     DirectionDelta = 0
	DIR_45_DEG_RIGHT                = 45
	DIR_90_DEG_RIGHT                = 90
	DIR_180_DEG                     = 180
	DIR_90_DEG_LEFT                 = 270
	DIR_45_DEG_LEFT                 = 315
	DIR_DIAGONAL                    = iota
	// 1/8th direction change left (?)
	DIR_DIAGONAL_LEFT = iota
	// 1/8th direction change (?)
	DIR_DIAGONAL_RIGHT = iota
)

var RCTDirectionKeys = map[int]DirectionDelta{
	0: DIR_STRAIGHT,
	1: DIR_90_DEG_RIGHT,
	2: DIR_180_DEG,
	3: DIR_90_DEG_LEFT,
	// XXX
	4: DIR_DIAGONAL,
	7: DIR_DIAGONAL_LEFT,
}

type Segment struct {
	// Corresponds to the hex number of this element's order, eg 1 is flat,
	// 2 is 25 deg up, etc.
	Type SegmentType

	// Computes the change in direction
	DirectionDelta DirectionDelta

	// The starting slope
	InputDegree Degree
	// The ending slope
	OutputDegree Degree

	// Change in tile position. A positive number indicates after the track
	// segment, we're that many tiles forward. A negative number indicates we
	// moved backwards
	ForwardDelta int

	// The total change in elevation
	ElevationDelta int

	// How far we moved side to side. Negative numbers indicate the track
	// turned left.
	SidewaysDelta int

	// Starting bank
	StartingBank Bank
	// the ending bank
	EndingBank Bank
}

type SegmentType int

// Track piece names
const (
	ELEM_FLAT                              SegmentType = 0
	ELEM_END_STATION                                   = 0x1
	ELEM_BEGIN_STATION                                 = 0x2
	ELEM_MIDDLE_STATION                                = 0x3
	ELEM_25_DEG_UP                                     = 0x4
	ELEM_60_DEG_UP                                     = 0x5
	ELEM_FLAT_TO_25_DEG_UP                             = 0x6
	ELEM_25_DEG_UP_TO_60_DEG_UP                        = 0x7
	ELEM_60_DEG_UP_TO_25_DEG_UP                        = 0x8
	ELEM_25_DEG_UP_TO_FLAT                             = 0x9
	ELEM_25_DEG_DOWN                                   = 0x0a
	ELEM_60_DEG_DOWN                                   = 0x0b
	ELEM_FLAT_TO_25_DEG_DOWN                           = 0x0c
	ELEM_25_DEG_DOWN_TO_60_DEG_DOWN                    = 0x0d
	ELEM_60_DEG_DOWN_TO_25_DEG_DOWN                    = 0x0e
	ELEM_25_DEG_DOWN_TO_FLAT                           = 0x0f
	ELEM_LEFT_QUARTER_TURN_5_TILES                     = 0x10
	ELEM_RIGHT_QUARTER_TURN_5_TILES                    = 0x11
	ELEM_FLAT_TO_LEFT_BANK                             = 0x12
	ELEM_FLAT_TO_RIGHT_BANK                            = 0x13
	ELEM_LEFT_BANK_TO_FLAT                             = 0x14
	ELEM_RIGHT_BANK_TO_FLAT                            = 0x15
	ELEM_BANKED_LEFT_QUARTER_TURN_5_TILES              = 0x16
	ELEM_BANKED_RIGHT_QUARTER_TURN_5_TILES             = 0x17
	ELEM_LEFT_BANK_TO_25_DEG_UP                        = 0x18
	ELEM_RIGHT_BANK_TO_25_DEG_UP                       = 0x19
	ELEM_25_DEG_UP_TO_LEFT_BANK                        = 0x1a
	ELEM_25_DEG_UP_TO_RIGHT_BANK                       = 0x1b
	ELEM_LEFT_BANK_TO_25_DEG_DOWN                      = 0x1c
	ELEM_RIGHT_BANK_TO_25_DEG_DOWN                     = 0x1d
	ELEM_25_DEG_DOWN_TO_LEFT_BANK                      = 0x1e
	ELEM_25_DEG_DOWN_TO_RIGHT_BANK                     = 0x1f
	ELEM_LEFT_BANK                                     = 0x20
	ELEM_RIGHT_BANK                                    = 0x21
	// Not banked, just turning and sliding downward
	ELEM_LEFT_QUARTER_TURN_5_TILES_25_DEG_UP    = 0x22
	ELEM_RIGHT_QUARTER_TURN_5_TILES_25_DEG_UP   = 0x23
	ELEM_LEFT_QUARTER_TURN_5_TILES_25_DEG_DOWN  = 0x24
	ELEM_RIGHT_QUARTER_TURN_5_TILES_25_DEG_DOWN = 0x25
	ELEM_S_BEND_LEFT                            = 0x26
	ELEM_S_BEND_RIGHT                           = 0x27
	ELEM_LEFT_VERTICAL_LOOP                     = 0x28
	ELEM_RIGHT_VERTICAL_LOOP                    = 0x29
	ELEM_LEFT_QUARTER_TURN_3_TILES              = 0x2a
	ELEM_RIGHT_QUARTER_TURN_3_TILES             = 0x2b
	ELEM_LEFT_QUARTER_TURN_3_TILES_BANK         = 0x2c
	ELEM_RIGHT_QUARTER_TURN_3_TILES_BANK        = 0x2d
	ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_UP    = 0x2e
	ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_UP   = 0x2f
	ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_DOWN  = 0x30
	ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_DOWN = 0x31
	ELEM_LEFT_QUARTER_TURN_1_TILE               = 0x32
	ELEM_RIGHT_QUARTER_TURN_1_TILE              = 0x33

	ELEM_LEFT_TWIST_DOWN_TO_UP  = 0x34
	ELEM_RIGHT_TWIST_DOWN_TO_UP = 0x35
	ELEM_LEFT_TWIST_UP_TO_DOWN  = 0x36
	ELEM_RIGHT_TWIST_UP_TO_DOWN = 0x37
	ELEM_HALF_LOOP_UP           = 0x38
	ELEM_HALF_LOOP_DOWN         = 0x39
	ELEM_LEFT_CORKSCREW_UP      = 0x3a
	ELEM_RIGHT_CORKSCREW_UP     = 0x3b
	ELEM_LEFT_CORKSCREW_DOWN    = 0x3c
	ELEM_RIGHT_CORKSCREW_DOWN   = 0x3d

	ELEM_FLAT_TO_60_DEG_UP   = 0x3e
	ELEM_60_DEG_UP_TO_FLAT   = 0x3f
	ELEM_FLAT_TO_60_DEG_DOWN = 0x40
	ELEM_60_DEG_DOWN_TO_FLAT = 0x41

	ELEM_TOWER_BASE    = 0x42
	ELEM_TOWER_SECTION = 0x43

	// covered track pieces
	ELEM_FLAT_COVERED                       = 0x44
	ELEM_25_DEG_UP_COVERED                  = 0x45
	ELEM_60_DEG_UP_COVERED                  = 0x46
	ELEM_FLAT_TO_25_DEG_UP_COVERED          = 0x47
	ELEM_25_DEG_UP_TO_60_DEG_UP_COVERED     = 0x48
	ELEM_60_DEG_UP_TO_25_DEG_UP_COVERED     = 0x49
	ELEM_25_DEG_UP_TO_FLAT_COVERED          = 0x4a
	ELEM_25_DEG_DOWN_COVERED                = 0x4b
	ELEM_60_DEG_DOWN_COVERED                = 0x4c
	ELEM_FLAT_TO_25_DEG_DOWN_COVERED        = 0x4d
	ELEM_25_DEG_DOWN_TO_60_DEG_DOWN_COVERED = 0x4e
	ELEM_60_DEG_DOWN_TO_25_DEG_DOWN_COVERED = 0x4f
	ELEM_25_DEG_DOWN_TO_FLAT_COVERED        = 0x50
	ELEM_LEFT_QUARTER_TURN_5_TILES_COVERED  = 0x51
	ELEM_RIGHT_QUARTER_TURN_5_TILES_COVERED = 0x52
	ELEM_S_BEND_LEFT_COVERED                = 0x53
	ELEM_S_BEND_RIGHT_COVERED               = 0x54
	ELEM_LEFT_QUARTER_TURN_3_TILES_COVERED  = 0x55
	ELEM_RIGHT_QUARTER_TURN_3_TILES_COVERED = 0x56

	ELEM_LEFT_HALF_BANKED_HELIX_UP_SMALL    = 0x57
	ELEM_RIGHT_HALF_BANKED_HELIX_UP_SMALL   = 0x58
	ELEM_LEFT_HALF_BANKED_HELIX_DOWN_SMALL  = 0x59
	ELEM_RIGHT_HALF_BANKED_HELIX_DOWN_SMALL = 0x5a
	ELEM_LEFT_HALF_BANKED_HELIX_UP_LARGE    = 0x5b
	ELEM_RIGHT_HALF_BANKED_HELIX_UP_LARGE   = 0x5c
	ELEM_LEFT_HALF_BANKED_HELIX_DOWN_LARGE  = 0x5d
	ELEM_RIGHT_HALF_BANKED_HELIX_DOWN_LARGE = 0x5e

	// Steeply sloped turns
	ELEM_LEFT_QUARTER_TURN_1_TILE_60_DEG_UP    = 0x5f
	ELEM_RIGHT_QUARTER_TURN_1_TILE_60_DEG_UP   = 0x60
	ELEM_LEFT_QUARTER_TURN_1_TILE_60_DEG_DOWN  = 0x61
	ELEM_RIGHT_QUARTER_TURN_1_TILE_60_DEG_DOWN = 0x62

	ELEM_BRAKES = 0x63

	ELEM_ROTATION_CONTROL_TOGGLE                 = 0x64
	ELEM_INVERTED_90_DEG_UP_TO_FLAT_QUARTER_LOOP = 0x65
	ELEM_LEFT_QUARTER_BANKED_HELIX_LARGE_UP      = 0x66
	ELEM_RIGHT_QUARTER_BANKED_HELIX_LARGE_UP     = 0x67
	ELEM_LEFT_QUARTER_BANKED_HELIX_LARGE_DOWN    = 0x68
	ELEM_RIGHT_QUARTER_BANKED_HELIX_LARGE_DOWN   = 0x69
	ELEM_LEFT_QUARTER_HELIX_LARGE_UP             = 0x6a
	ELEM_RIGHT_QUARTER_HELIX_LARGE_UP            = 0x6b
	ELEM_LEFT_QUARTER_HELIX_LARGE_DOWN           = 0x6c
	ELEM_RIGHT_QUARTER_HELIX_LARGE_DOWN          = 0x6d

	ELEM_25_DEG_UP_LEFT_BANKED  = 0x6e
	ELEM_25_DEG_UP_RIGHT_BANKED = 0x6f

	ELEM_WATERFALL     = 0x70
	ELEM_RAPIDS        = 0x71
	ELEM_ON_RIDE_PHOTO = 0x72

	// Not sure about this piece, the exe indicates that this is in the same
	// category as a water feature/rapids/photo
	ELEM_25_DEG_DOWN_LEFT_BANKED  = 0x73
	ELEM_25_DEG_DOWN_RIGHT_BANKED = 0x74

	ELEM_WATER_SPLASH = 0x75

	ELEM_FLAT_TO_60_DEG_UP_LONG_BASE = 0x76
	ELEM_60_DEG_UP_TO_FLAT_LONG_BASE = 0x77

	ELEM_WHIRLPOOL = 0x78

	ELEM_60_DEG_DOWN_TO_FLAT_LONG_BASE = 0x79
	ELEM_FLAT_TO_60_DEG_DOWN_LONG_BASE = 0x7a

	ELEM_CABLE_LIFT_HILL             = 0x7b
	ELEM_REVERSE_WHOA_BELLY_SLOPE    = 0x7c
	ELEM_REVERSE_WHOA_BELLY_VERTICAL = 0x7d

	ELEM_90_DEG_UP                  = 0x7e
	ELEM_90_DEG_DOWN                = 0x7f
	ELEM_60_DEG_UP_TO_90_DEG_UP     = 0x80
	ELEM_90_DEG_DOWN_TO_60_DEG_DOWN = 0x81
	ELEM_90_DEG_UP_TO_60_DEG_UP     = 0x82
	ELEM_60_DEG_DOWN_TO_90_DEG_DOWN = 0x83
	ELEM_BRAKE_FOR_DROP             = 0x84

	ELEM_LEFT_EIGHTH_TO_DIAG             = 0x85
	ELEM_RIGHT_EIGHTH_TO_DIAG            = 0x86
	ELEM_LEFT_EIGHTH_TO_ORTHOGONAL       = 0x87
	ELEM_RIGHT_EIGHTH_TO_ORTHOGONAL      = 0x88
	ELEM_LEFT_EIGHTH_BANK_TO_DIAG        = 0x89
	ELEM_RIGHT_EIGHTH_BANK_TO_DIAG       = 0x8a
	ELEM_LEFT_EIGHTH_BANK_TO_ORTHOGONAL  = 0x8b
	ELEM_RIGHT_EIGHTH_BANK_TO_ORTHOGONAL = 0x8c
	ELEM_DIAG_FLAT                       = 0x8d
	ELEM_DIAG_25_DEG_UP                  = 0x8e
	ELEM_DIAG_60_DEG_UP                  = 0x8f
	ELEM_DIAG_FLAT_TO_25_DEG_UP          = 0x90
	ELEM_DIAG_25_DEG_UP_TO_60_DEG_UP     = 0x91
	ELEM_DIAG_60_DEG_UP_TO_25_DEG_UP     = 0x92
	ELEM_DIAG_25_DEG_UP_TO_FLAT          = 0x93
	ELEM_DIAG_25_DEG_DOWN                = 0x94
	ELEM_DIAG_60_DEG_DOWN                = 0x95
	ELEM_DIAG_FLAT_TO_25_DEG_DOWN        = 0x96
	ELEM_DIAG_25_DEG_DOWN_TO_60_DEG_DOWN = 0x97
	ELEM_DIAG_60_DEG_DOWN_TO_25_DEG_DOWN = 0x98
	ELEM_DIAG_25_DEG_DOWN_TO_FLAT        = 0x99
	ELEM_DIAG_FLAT_TO_60_DEG_UP          = 0x9a
	ELEM_DIAG_60_DEG_UP_TO_FLAT          = 0x9b
	ELEM_DIAG_FLAT_TO_60_DEG_DOWN        = 0x9c
	ELEM_DIAG_60_DEG_DOWN_TO_FLAT        = 0x9d
	ELEM_DIAG_FLAT_TO_LEFT_BANK          = 0x9e
	ELEM_DIAG_FLAT_TO_RIGHT_BANK         = 0x9f
	ELEM_DIAG_LEFT_BANK_TO_FLAT          = 0xa0
	ELEM_DIAG_RIGHT_BANK_TO_FLAT         = 0xa1
	ELEM_DIAG_LEFT_BANK_TO_25_DEG_UP     = 0xa2
	ELEM_DIAG_RIGHT_BANK_TO_25_DEG_UP    = 0xa3
	ELEM_DIAG_25_DEG_UP_TO_LEFT_BANK     = 0xa4
	ELEM_DIAG_25_DEG_UP_TO_RIGHT_BANK    = 0xa5
	ELEM_DIAG_LEFT_BANK_TO_25_DEG_DOWN   = 0xa6
	ELEM_DIAG_RIGHT_BANK_TO_25_DEG_DOWN  = 0xa7
	ELEM_DIAG_25_DEG_DOWN_TO_LEFT_BANK   = 0xa8
	ELEM_DIAG_25_DEG_DOWN_TO_RIGHT_BANK  = 0xa9
	ELEM_DIAG_LEFT_BANK                  = 0xaa
	ELEM_DIAG_RIGHT_BANK                 = 0xab

	ELEM_LOG_FLUME_REVERSER = 0xac
	ELEM_SPINNING_TUNNEL    = 0xad

	ELEM_LEFT_BARREL_ROLL_UP_TO_DOWN  = 0xae
	ELEM_RIGHT_BARREL_ROLL_UP_TO_DOWN = 0xaf
	ELEM_LEFT_BARREL_ROLL_DOWN_TO_UP  = 0xb0
	ELEM_RIGHT_BARREL_ROLL_DOWN_TO_UP = 0xb1

	ELEM_LEFT_BANK_TO_LEFT_QUARTER_TURN_3_TILES_25_DEG_UP     = 0xb2
	ELEM_RIGHT_BANK_TO_RIGHT_QUARTER_TURN_3_TILES_25_DEG_UP   = 0xb3
	ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_DOWN_TO_LEFT_BANK   = 0xb4
	ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_DOWN_TO_RIGHT_BANK = 0xb5

	ELEM_POWERED_LIFT = 0xb6

	ELEM_LEFT_LARGE_HALF_LOOP_UP    = 0xb7
	ELEM_RIGHT_LARGE_HALF_LOOP_UP   = 0xb8
	ELEM_RIGHT_LARGE_HALF_LOOP_DOWN = 0xb9
	ELEM_LEFT_LARGE_HALF_LOOP_DOWN  = 0xba

	ELEM_BLOCK_BRAKES = 0xd8

	ELEM_END_OF_RIDE = 0xff
)

var TS_MAP = map[SegmentType]*Segment{
	ELEM_FLAT: &Segment{
		Type:           0x0,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_END_STATION: &Segment{
		Type:           0x1,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_BEGIN_STATION: &Segment{
		Type:           0x2,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_MIDDLE_STATION: &Segment{
		Type:           0x3,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_25_DEG_UP: &Segment{
		Type:           0x4,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 2,
	},
	ELEM_60_DEG_UP: &Segment{
		Type:           0x5,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 8,
	},
	ELEM_FLAT_TO_25_DEG_UP: &Segment{
		Type:           0x6,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_25_DEG_UP_TO_60_DEG_UP: &Segment{
		Type:           0x7,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 4,
	},
	ELEM_60_DEG_UP_TO_25_DEG_UP: &Segment{
		Type:           0x8,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 4,
	},
	ELEM_25_DEG_UP_TO_FLAT: &Segment{
		Type:           0x9,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_25_DEG_DOWN: &Segment{
		Type:           0xa,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -2,
	},
	ELEM_60_DEG_DOWN: &Segment{
		Type:           0xb,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -8,
	},
	ELEM_FLAT_TO_25_DEG_DOWN: &Segment{
		Type:           0xc,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_25_DEG_DOWN_TO_60_DEG_DOWN: &Segment{
		Type:           0xd,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -4,
	},
	ELEM_60_DEG_DOWN_TO_25_DEG_DOWN: &Segment{
		Type:           0xe,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -4,
	},
	ELEM_25_DEG_DOWN_TO_FLAT: &Segment{
		Type:           0xf,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_LEFT_QUARTER_TURN_5_TILES: &Segment{
		Type:           0x10,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -3,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_QUARTER_TURN_5_TILES: &Segment{
		Type:           0x11,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  3,
		ElevationDelta: 0,
	},
	ELEM_FLAT_TO_LEFT_BANK: &Segment{
		Type:           0x12,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_FLAT_TO_RIGHT_BANK: &Segment{
		Type:           0x13,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_BANK_TO_FLAT: &Segment{
		Type:           0x14,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_BANK_TO_FLAT: &Segment{
		Type:           0x15,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_BANKED_LEFT_QUARTER_TURN_5_TILES: &Segment{
		Type:           0x16,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -3,
		ElevationDelta: 0,
	},
	ELEM_BANKED_RIGHT_QUARTER_TURN_5_TILES: &Segment{
		Type:           0x17,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  3,
		ElevationDelta: 0,
	},
	ELEM_LEFT_BANK_TO_25_DEG_UP: &Segment{
		Type:           0x18,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_RIGHT_BANK_TO_25_DEG_UP: &Segment{
		Type:           0x19,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_25_DEG_UP_TO_LEFT_BANK: &Segment{
		Type:           0x1a,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_25_DEG_UP_TO_RIGHT_BANK: &Segment{
		Type:           0x1b,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_LEFT_BANK_TO_25_DEG_DOWN: &Segment{
		Type:           0x1c,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_RIGHT_BANK_TO_25_DEG_DOWN: &Segment{
		Type:           0x1d,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_25_DEG_DOWN_TO_LEFT_BANK: &Segment{
		Type:           0x1e,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_25_DEG_DOWN_TO_RIGHT_BANK: &Segment{
		Type:           0x1f,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_LEFT_BANK: &Segment{
		Type:           0x20,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_BANK: &Segment{
		Type:           0x21,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_QUARTER_TURN_5_TILES_25_DEG_UP: &Segment{
		Type:           0x22,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -3,
		ElevationDelta: 8,
	},
	ELEM_RIGHT_QUARTER_TURN_5_TILES_25_DEG_UP: &Segment{
		Type:           0x23,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  3,
		ElevationDelta: 8,
	},
	ELEM_LEFT_QUARTER_TURN_5_TILES_25_DEG_DOWN: &Segment{
		Type:           0x24,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -3,
		ElevationDelta: -8,
	},
	ELEM_RIGHT_QUARTER_TURN_5_TILES_25_DEG_DOWN: &Segment{
		Type:           0x25,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  3,
		ElevationDelta: -8,
	},
	ELEM_S_BEND_LEFT: &Segment{
		Type:           0x26,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  -2,
		ElevationDelta: 0,
	},
	ELEM_S_BEND_RIGHT: &Segment{
		Type:           0x27,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_VERTICAL_LOOP: &Segment{
		Type:           0x28,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  -2,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_VERTICAL_LOOP: &Segment{
		Type:           0x29,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_QUARTER_TURN_3_TILES: &Segment{
		Type:           0x2a,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -2,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_QUARTER_TURN_3_TILES: &Segment{
		Type:           0x2b,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_QUARTER_TURN_3_TILES_BANK: &Segment{
		Type:           0x2c,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -2,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_QUARTER_TURN_3_TILES_BANK: &Segment{
		Type:           0x2d,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_UP: &Segment{
		Type:           0x2e,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -2,
		ElevationDelta: 4,
	},
	ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_UP: &Segment{
		Type:           0x2f,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 4,
	},
	ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_DOWN: &Segment{
		Type:           0x30,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -2,
		ElevationDelta: -4,
	},
	ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_DOWN: &Segment{
		Type:           0x31,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -4,
	},
	ELEM_LEFT_QUARTER_TURN_1_TILE: &Segment{
		Type:           0x32,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_QUARTER_TURN_1_TILE: &Segment{
		Type:           0x33,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_TWIST_DOWN_TO_UP: &Segment{
		Type:           0x34,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_UPSIDE_DOWN,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 2,
	},
	ELEM_RIGHT_TWIST_DOWN_TO_UP: &Segment{
		Type:           0x35,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_UPSIDE_DOWN,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 2,
	},
	ELEM_LEFT_TWIST_UP_TO_DOWN: &Segment{
		Type:           0x36,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_UPSIDE_DOWN,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 30,
	},
	ELEM_RIGHT_TWIST_UP_TO_DOWN: &Segment{
		Type:           0x37,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_UPSIDE_DOWN,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 30,
	},
	ELEM_HALF_LOOP_UP: &Segment{
		Type:           0x38,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_UPSIDE_DOWN,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 19,
	},
	ELEM_HALF_LOOP_DOWN: &Segment{
		Type:           0x39,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_UPSIDE_DOWN,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 13,
	},
	ELEM_LEFT_CORKSCREW_UP: &Segment{
		Type:           0x3a,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_UPSIDE_DOWN,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -2,
		ElevationDelta: 10,
	},
	ELEM_RIGHT_CORKSCREW_UP: &Segment{
		Type:           0x3b,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_UPSIDE_DOWN,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 10,
	},
	ELEM_LEFT_CORKSCREW_DOWN: &Segment{
		Type:           0x3c,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_UPSIDE_DOWN,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -2,
		ElevationDelta: 22,
	},
	ELEM_RIGHT_CORKSCREW_DOWN: &Segment{
		Type:           0x3d,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_UPSIDE_DOWN,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 22,
	},
	ELEM_FLAT_TO_60_DEG_UP: &Segment{
		Type:           0x3e,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 3,
	},
	ELEM_60_DEG_UP_TO_FLAT: &Segment{
		Type:           0x3f,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 3,
	},
	ELEM_FLAT_TO_60_DEG_DOWN: &Segment{
		Type:           0x40,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -3,
	},
	ELEM_60_DEG_DOWN_TO_FLAT: &Segment{
		Type:           0x41,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -3,
	},
	ELEM_TOWER_BASE: &Segment{
		Type:           0x42,
		InputDegree:    TRACK_UP_90,
		OutputDegree:   TRACK_UP_90,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 12,
	},
	ELEM_TOWER_SECTION: &Segment{
		Type:           0x43,
		InputDegree:    TRACK_UP_90,
		OutputDegree:   TRACK_UP_90,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 4,
	},
	ELEM_FLAT_COVERED: &Segment{
		Type:           0x44,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_25_DEG_UP_COVERED: &Segment{
		Type:           0x45,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 2,
	},
	ELEM_60_DEG_UP_COVERED: &Segment{
		Type:           0x46,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 8,
	},
	ELEM_FLAT_TO_25_DEG_UP_COVERED: &Segment{
		Type:           0x47,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_25_DEG_UP_TO_60_DEG_UP_COVERED: &Segment{
		Type:           0x48,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 4,
	},
	ELEM_60_DEG_UP_TO_25_DEG_UP_COVERED: &Segment{
		Type:           0x49,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 4,
	},
	ELEM_25_DEG_UP_TO_FLAT_COVERED: &Segment{
		Type:           0x4a,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_25_DEG_DOWN_COVERED: &Segment{
		Type:           0x4b,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -2,
	},
	ELEM_60_DEG_DOWN_COVERED: &Segment{
		Type:           0x4c,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -8,
	},
	ELEM_FLAT_TO_25_DEG_DOWN_COVERED: &Segment{
		Type:           0x4d,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_25_DEG_DOWN_TO_60_DEG_DOWN_COVERED: &Segment{
		Type:           0x4e,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -4,
	},
	ELEM_60_DEG_DOWN_TO_25_DEG_DOWN_COVERED: &Segment{
		Type:           0x4f,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -4,
	},
	ELEM_25_DEG_DOWN_TO_FLAT_COVERED: &Segment{
		Type:           0x50,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_LEFT_QUARTER_TURN_5_TILES_COVERED: &Segment{
		Type:           0x51,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -3,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_QUARTER_TURN_5_TILES_COVERED: &Segment{
		Type:           0x52,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  3,
		ElevationDelta: 0,
	},
	ELEM_S_BEND_LEFT_COVERED: &Segment{
		Type:           0x53,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  -2,
		ElevationDelta: 0,
	},
	ELEM_S_BEND_RIGHT_COVERED: &Segment{
		Type:           0x54,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_QUARTER_TURN_3_TILES_COVERED: &Segment{
		Type:           0x55,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -2,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_QUARTER_TURN_3_TILES_COVERED: &Segment{
		Type:           0x56,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_HALF_BANKED_HELIX_UP_SMALL: &Segment{
		Type:           0x57,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  -4,
		ElevationDelta: 2,
	},
	ELEM_RIGHT_HALF_BANKED_HELIX_UP_SMALL: &Segment{
		Type:           0x58,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  4,
		ElevationDelta: 2,
	},
	ELEM_LEFT_HALF_BANKED_HELIX_DOWN_SMALL: &Segment{
		Type:           0x59,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  -4,
		ElevationDelta: -2,
	},
	ELEM_RIGHT_HALF_BANKED_HELIX_DOWN_SMALL: &Segment{
		Type:           0x5a,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  4,
		ElevationDelta: -2,
	},
	ELEM_LEFT_HALF_BANKED_HELIX_UP_LARGE: &Segment{
		Type:           0x5b,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  4,
		ElevationDelta: 2,
	},
	ELEM_RIGHT_HALF_BANKED_HELIX_UP_LARGE: &Segment{
		Type:           0x5c,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  -4,
		ElevationDelta: 2,
	},
	ELEM_LEFT_HALF_BANKED_HELIX_DOWN_LARGE: &Segment{
		Type:           0x5d,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  4,
		ElevationDelta: -2,
	},
	ELEM_RIGHT_HALF_BANKED_HELIX_DOWN_LARGE: &Segment{
		Type:           0x5e,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  -4,
		ElevationDelta: -2,
	},
	ELEM_LEFT_QUARTER_TURN_1_TILE_60_DEG_UP: &Segment{
		Type:           0x5f,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_QUARTER_TURN_1_TILE_60_DEG_UP: &Segment{
		Type:           0x60,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 8,
	},
	ELEM_LEFT_QUARTER_TURN_1_TILE_60_DEG_DOWN: &Segment{
		Type:           0x61,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  0,
		ElevationDelta: -8,
	},
	ELEM_RIGHT_QUARTER_TURN_1_TILE_60_DEG_DOWN: &Segment{
		Type:           0x62,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -8,
	},
	ELEM_BRAKES: &Segment{
		Type:           0x63,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_ROTATION_CONTROL_TOGGLE: &Segment{
		Type:           0x64,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_INVERTED_90_DEG_UP_TO_FLAT_QUARTER_LOOP: &Segment{
		Type:           0x65,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_QUARTER_BANKED_HELIX_LARGE_UP: &Segment{
		Type:           0x66,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -3,
		ElevationDelta: 2,
	},
	ELEM_RIGHT_QUARTER_BANKED_HELIX_LARGE_UP: &Segment{
		Type:           0x67,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  3,
		ElevationDelta: 2,
	},
	ELEM_LEFT_QUARTER_BANKED_HELIX_LARGE_DOWN: &Segment{
		Type:           0x68,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -3,
		ElevationDelta: -2,
	},
	ELEM_RIGHT_QUARTER_BANKED_HELIX_LARGE_DOWN: &Segment{
		Type:           0x69,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  3,
		ElevationDelta: -2,
	},
	ELEM_LEFT_QUARTER_HELIX_LARGE_UP: &Segment{
		Type:           0x6a,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -3,
		ElevationDelta: 2,
	},
	ELEM_RIGHT_QUARTER_HELIX_LARGE_UP: &Segment{
		Type:           0x6b,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  3,
		ElevationDelta: 2,
	},
	ELEM_LEFT_QUARTER_HELIX_LARGE_DOWN: &Segment{
		Type:           0x6c,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -3,
		ElevationDelta: -2,
	},
	ELEM_RIGHT_QUARTER_HELIX_LARGE_DOWN: &Segment{
		Type:           0x6d,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  3,
		ElevationDelta: -2,
	},
	ELEM_25_DEG_UP_LEFT_BANKED: &Segment{
		Type:           0x6e,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_25_DEG_UP_RIGHT_BANKED: &Segment{
		Type:           0x6f,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_WATERFALL: &Segment{
		Type:           0x70,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_RAPIDS: &Segment{
		Type:           0x71,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_ON_RIDE_PHOTO: &Segment{
		Type:           0x72,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_25_DEG_DOWN_LEFT_BANKED: &Segment{
		Type:           0x73,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_25_DEG_DOWN_RIGHT_BANKED: &Segment{
		Type:           0x74,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_WATER_SPLASH: &Segment{
		Type:           0x75,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 2,
	},
	ELEM_FLAT_TO_60_DEG_UP_LONG_BASE: &Segment{
		Type:           0x76,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_60_DEG_UP_TO_FLAT_LONG_BASE: &Segment{
		Type:           0x77,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_WHIRLPOOL: &Segment{
		Type:           0x78,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_60_DEG_DOWN_TO_FLAT_LONG_BASE: &Segment{
		Type:           0x79,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_FLAT_TO_60_DEG_DOWN_LONG_BASE: &Segment{
		Type:           0x7a,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_CABLE_LIFT_HILL: &Segment{
		Type:           0x7b,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 12,
	},
	ELEM_REVERSE_WHOA_BELLY_SLOPE: &Segment{
		Type:           0x7c,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_90,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 30,
	},
	ELEM_REVERSE_WHOA_BELLY_VERTICAL: &Segment{
		Type:           0x7d,
		InputDegree:    TRACK_UP_90,
		OutputDegree:   TRACK_UP_90,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 10,
	},
	ELEM_90_DEG_UP: &Segment{
		Type:           0x7e,
		InputDegree:    TRACK_UP_90,
		OutputDegree:   TRACK_UP_90,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 4,
	},
	ELEM_90_DEG_DOWN: &Segment{
		Type:           0x7f,
		InputDegree:    TRACK_DOWN_90,
		OutputDegree:   TRACK_DOWN_90,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -4,
	},
	ELEM_60_DEG_UP_TO_90_DEG_UP: &Segment{
		Type:           0x80,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_UP_90,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 7,
	},
	ELEM_90_DEG_DOWN_TO_60_DEG_DOWN: &Segment{
		Type:           0x81,
		InputDegree:    TRACK_DOWN_90,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -7,
	},
	ELEM_90_DEG_UP_TO_60_DEG_UP: &Segment{
		Type:           0x82,
		InputDegree:    TRACK_UP_90,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 7,
	},
	ELEM_60_DEG_DOWN_TO_90_DEG_DOWN: &Segment{
		Type:           0x83,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_DOWN_90,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -7,
	},
	ELEM_BRAKE_FOR_DROP: &Segment{
		Type:           0x84,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -3,
	},
	ELEM_LEFT_EIGHTH_TO_DIAG: &Segment{
		Type:           0x85,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL_LEFT,
		SidewaysDelta:  -2,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_EIGHTH_TO_DIAG: &Segment{
		Type:           0x86,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_EIGHTH_TO_ORTHOGONAL: &Segment{
		Type:           0x87,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL_LEFT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_EIGHTH_TO_ORTHOGONAL: &Segment{
		Type:           0x88,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL_RIGHT,
		SidewaysDelta:  3,
		ElevationDelta: 0,
	},
	ELEM_LEFT_EIGHTH_BANK_TO_DIAG: &Segment{
		Type:           0x89,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_DIAGONAL_LEFT,
		SidewaysDelta:  -2,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_EIGHTH_BANK_TO_DIAG: &Segment{
		Type:           0x8a,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_DIAGONAL_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_EIGHTH_BANK_TO_ORTHOGONAL: &Segment{
		Type:           0x8b,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_DIAGONAL_LEFT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_RIGHT_EIGHTH_BANK_TO_ORTHOGONAL: &Segment{
		Type:           0x8c,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_DIAGONAL_RIGHT,
		SidewaysDelta:  3,
		ElevationDelta: 0,
	},
	ELEM_DIAG_FLAT: &Segment{
		Type:           0x8d,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_DIAG_25_DEG_UP: &Segment{
		Type:           0x8e,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 2,
	},
	ELEM_DIAG_60_DEG_UP: &Segment{
		Type:           0x8f,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 8,
	},
	ELEM_DIAG_FLAT_TO_25_DEG_UP: &Segment{
		Type:           0x90,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_DIAG_25_DEG_UP_TO_60_DEG_UP: &Segment{
		Type:           0x91,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 4,
	},
	ELEM_DIAG_60_DEG_UP_TO_25_DEG_UP: &Segment{
		Type:           0x92,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 4,
	},
	ELEM_DIAG_25_DEG_UP_TO_FLAT: &Segment{
		Type:           0x93,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_DIAG_25_DEG_DOWN: &Segment{
		Type:           0x94,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -2,
	},
	ELEM_DIAG_60_DEG_DOWN: &Segment{
		Type:           0x95,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -8,
	},
	ELEM_DIAG_FLAT_TO_25_DEG_DOWN: &Segment{
		Type:           0x96,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_DIAG_25_DEG_DOWN_TO_60_DEG_DOWN: &Segment{
		Type:           0x97,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -4,
	},
	ELEM_DIAG_60_DEG_DOWN_TO_25_DEG_DOWN: &Segment{
		Type:           0x98,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -4,
	},
	ELEM_DIAG_25_DEG_DOWN_TO_FLAT: &Segment{
		Type:           0x99,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_DIAG_FLAT_TO_60_DEG_UP: &Segment{
		Type:           0x9a,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 3,
	},
	ELEM_DIAG_60_DEG_UP_TO_FLAT: &Segment{
		Type:           0x9b,
		InputDegree:    TRACK_UP_60,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 3,
	},
	ELEM_DIAG_FLAT_TO_60_DEG_DOWN: &Segment{
		Type:           0x9c,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_60,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -3,
	},
	ELEM_DIAG_60_DEG_DOWN_TO_FLAT: &Segment{
		Type:           0x9d,
		InputDegree:    TRACK_DOWN_60,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -3,
	},
	ELEM_DIAG_FLAT_TO_LEFT_BANK: &Segment{
		Type:           0x9e,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_DIAG_FLAT_TO_RIGHT_BANK: &Segment{
		Type:           0x9f,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_DIAG_LEFT_BANK_TO_FLAT: &Segment{
		Type:           0xa0,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_DIAG_RIGHT_BANK_TO_FLAT: &Segment{
		Type:           0xa1,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_DIAG_LEFT_BANK_TO_25_DEG_UP: &Segment{
		Type:           0xa2,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_DIAG_RIGHT_BANK_TO_25_DEG_UP: &Segment{
		Type:           0xa3,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_DIAG_25_DEG_UP_TO_LEFT_BANK: &Segment{
		Type:           0xa4,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_DIAG_25_DEG_UP_TO_RIGHT_BANK: &Segment{
		Type:           0xa5,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 1,
	},
	ELEM_DIAG_LEFT_BANK_TO_25_DEG_DOWN: &Segment{
		Type:           0xa6,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_DIAG_RIGHT_BANK_TO_25_DEG_DOWN: &Segment{
		Type:           0xa7,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_DIAG_25_DEG_DOWN_TO_LEFT_BANK: &Segment{
		Type:           0xa8,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_DIAG_25_DEG_DOWN_TO_RIGHT_BANK: &Segment{
		Type:           0xa9,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: -1,
	},
	ELEM_DIAG_LEFT_BANK: &Segment{
		Type:           0xaa,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_DIAG_RIGHT_BANK: &Segment{
		Type:           0xab,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_DIAGONAL,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LOG_FLUME_REVERSER: &Segment{
		Type:           0xac,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_SPINNING_TUNNEL: &Segment{
		Type:           0xad,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 0,
	},
	ELEM_LEFT_BARREL_ROLL_UP_TO_DOWN: &Segment{
		Type:           0xae,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_UPSIDE_DOWN,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 4,
	},
	ELEM_RIGHT_BARREL_ROLL_UP_TO_DOWN: &Segment{
		Type:           0xaf,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_UPSIDE_DOWN,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 4,
	},
	ELEM_LEFT_BARREL_ROLL_DOWN_TO_UP: &Segment{
		Type:           0xb0,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_UPSIDE_DOWN,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 28,
	},
	ELEM_RIGHT_BARREL_ROLL_DOWN_TO_UP: &Segment{
		Type:           0xb1,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_UPSIDE_DOWN,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 28,
	},
	ELEM_LEFT_BANK_TO_LEFT_QUARTER_TURN_3_TILES_25_DEG_UP: &Segment{
		Type:           0xb2,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_LEFT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -2,
		ElevationDelta: 3,
	},
	ELEM_RIGHT_BANK_TO_RIGHT_QUARTER_TURN_3_TILES_25_DEG_UP: &Segment{
		Type:           0xb3,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_RIGHT,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 3,
	},
	ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_DOWN_TO_LEFT_BANK: &Segment{
		Type:           0xb4,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_LEFT,
		DirectionDelta: DIR_90_DEG_LEFT,
		SidewaysDelta:  -2,
		ElevationDelta: -3,
	},
	ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_DOWN_TO_RIGHT_BANK: &Segment{
		Type:           0xb5,
		InputDegree:    TRACK_DOWN_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_RIGHT,
		DirectionDelta: DIR_90_DEG_RIGHT,
		SidewaysDelta:  0,
		ElevationDelta: -3,
	},
	ELEM_POWERED_LIFT: &Segment{
		Type:           0xb6,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_UP_25,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_STRAIGHT,
		SidewaysDelta:  0,
		ElevationDelta: 2,
	},
	ELEM_LEFT_LARGE_HALF_LOOP_UP: &Segment{
		Type:           0xb7,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_UPSIDE_DOWN,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  -2,
		ElevationDelta: 3,
	},
	ELEM_RIGHT_LARGE_HALF_LOOP_UP: &Segment{
		Type:           0xb8,
		InputDegree:    TRACK_UP_25,
		OutputDegree:   TRACK_NONE,
		StartingBank:   TRACK_BANK_NONE,
		EndingBank:     TRACK_BANK_UPSIDE_DOWN,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 3,
	},
	ELEM_RIGHT_LARGE_HALF_LOOP_DOWN: &Segment{
		Type:           0xb9,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_UPSIDE_DOWN,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  -2,
		ElevationDelta: 29,
	},
	ELEM_LEFT_LARGE_HALF_LOOP_DOWN: &Segment{
		Type:           0xba,
		InputDegree:    TRACK_NONE,
		OutputDegree:   TRACK_DOWN_25,
		StartingBank:   TRACK_BANK_UPSIDE_DOWN,
		EndingBank:     TRACK_BANK_NONE,
		DirectionDelta: DIR_180_DEG,
		SidewaysDelta:  0,
		ElevationDelta: 29,
	},

	ELEM_END_OF_RIDE: &Segment{
		Type: 0xff,
	},
}

func (s Segment) String() string {
	return strings.Replace(ElementNames[s.Type], "ELEM_", "", 1)
}
