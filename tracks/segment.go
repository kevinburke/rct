package tracks

type SegmentType int

type Segment struct {
	Type           SegmentType
	ElevationGain  int
	DirectionDelta DirectionDelta

	InputDegree  Degree
	OutputDegree Degree

	// Change in tile position. A positive number indicates after the track
	// segment, we're that many tiles forward. A negative number indicates we
	// moved backwards
	ForwardDelta int

	// How far we moved side to side. Negative numbers indicate the track
	// turned left.
	SidewaysDelta int
}

var TS_MAP = map[SegmentType]*Segment{
	ELEM_FLAT: &Segment{
		InputDegree:   0,
		OutputDegree:  0,
		ElevationGain: 0,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_END_STATION: &Segment{
		InputDegree:   0,
		OutputDegree:  0,
		ElevationGain: 0,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_BEGIN_STATION: &Segment{
		InputDegree:   0,
		OutputDegree:  0,
		ElevationGain: 0,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_MIDDLE_STATION: &Segment{
		Type:          0x3,
		InputDegree:   0,
		OutputDegree:  0,
		ElevationGain: 0,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_25_DEG_UP: &Segment{
		Type:          0x4,
		InputDegree:   0,
		OutputDegree:  DEGREE_25_UP,
		ElevationGain: 1,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_60_DEG_UP: &Segment{
		Type:          0x5,
		InputDegree:   0,
		OutputDegree:  DEGREE_60_UP,
		ElevationGain: 3,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_FLAT_TO_25_DEG_UP: &Segment{
		Type:          0x6,
		InputDegree:   0,
		OutputDegree:  DEGREE_25_UP,
		ElevationGain: 0,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_25_DEG_UP_TO_60_DEG_UP: &Segment{
		Type:          0x7,
		InputDegree:   DEGREE_25_UP,
		OutputDegree:  DEGREE_60_UP,
		ElevationGain: 3,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_60_DEG_UP_TO_25_DEG_UP: &Segment{
		Type:          0x8,
		InputDegree:   DEGREE_60_UP,
		OutputDegree:  DEGREE_25_UP,
		ElevationGain: 2,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_25_DEG_UP_TO_FLAT: &Segment{
		Type:          0x9,
		InputDegree:   DEGREE_25_UP,
		OutputDegree:  DEGREE_FLAT,
		ElevationGain: 1,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_25_DEG_DOWN: &Segment{
		Type: 0x0a,
	},
	ELEM_60_DEG_DOWN: &Segment{
		Type: 0x0b,
	},
	ELEM_FLAT_TO_25_DEG_DOWN: &Segment{
		Type:          0x0c,
		InputDegree:   DEGREE_FLAT,
		OutputDegree:  DEGREE_25_DOWN,
		ElevationGain: -1,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_25_DEG_DOWN_TO_60_DEG_DOWN: &Segment{
		Type: 0x0d,
	},
	ELEM_60_DEG_DOWN_TO_25_DEG_DOWN: &Segment{
		Type: 0x0e,
	},
	ELEM_25_DEG_DOWN_TO_FLAT: &Segment{
		Type:          0x0f,
		InputDegree:   DEGREE_25_DOWN,
		OutputDegree:  DEGREE_FLAT,
		ElevationGain: 0,
		ForwardDelta:  1,
		SidewaysDelta: 0,
	},
	ELEM_LEFT_QUARTER_TURN_5_TILES: &Segment{
		Type: 0x10,
	},
	ELEM_RIGHT_QUARTER_TURN_5_TILES: &Segment{
		Type: 0x11,
	},
	ELEM_FLAT_TO_LEFT_BANK: &Segment{
		Type: 0x12,
	},
	ELEM_FLAT_TO_RIGHT_BANK: &Segment{
		Type: 0x13,
	},
	ELEM_LEFT_BANK_TO_FLAT: &Segment{
		Type: 0x14,
	},
	ELEM_RIGHT_BANK_TO_FLAT: &Segment{
		Type: 0x15,
	},
	ELEM_BANKED_LEFT_QUARTER_TURN_5_TILES: &Segment{
		Type: 0x16,
	},
	ELEM_BANKED_RIGHT_QUARTER_TURN_5_TILES: &Segment{
		Type: 0x17,
	},
	ELEM_LEFT_BANK_TO_25_DEG_UP: &Segment{
		Type: 0x18,
	},
	ELEM_RIGHT_BANK_TO_25_DEG_UP: &Segment{
		Type: 0x19,
	},
	ELEM_25_DEG_UP_TO_LEFT_BANK: &Segment{
		Type: 0x1a,
	},
	ELEM_25_DEG_UP_TO_RIGHT_BANK: &Segment{
		Type: 0x1b,
	},
	ELEM_LEFT_BANK_TO_25_DEG_DOWN: &Segment{
		Type: 0x1c,
	},
	ELEM_RIGHT_BANK_TO_25_DEG_DOWN: &Segment{
		Type: 0x1d,
	},
	ELEM_25_DEG_DOWN_TO_LEFT_BANK: &Segment{
		Type: 0x1e,
	},
	ELEM_25_DEG_DOWN_TO_RIGHT_BANK: &Segment{
		Type: 0x1f,
	},
	ELEM_LEFT_BANK: &Segment{
		Type: 0x20,
	},
	ELEM_RIGHT_BANK: &Segment{
		Type: 0x21,
	},
	// Not banked, just turning and sliding downward
	ELEM_LEFT_QUARTER_TURN_5_TILES_25_DEG_UP: &Segment{
		Type: 0x22,
	},
	ELEM_RIGHT_QUARTER_TURN_5_TILES_25_DEG_UP: &Segment{
		Type: 0x23,
	},
	ELEM_LEFT_QUARTER_TURN_5_TILES_25_DEG_DOWN: &Segment{
		Type: 0x24,
	},
	ELEM_RIGHT_QUARTER_TURN_5_TILES_25_DEG_DOWN: &Segment{
		Type: 0x25,
	},

	ELEM_LEFT_QUARTER_TURN_3_TILES_BANK: &Segment{
		Type: 0x2c,
	},

	ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_DOWN: &Segment{
		Type: 0x30,
	},
	ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_DOWN: &Segment{
		Type: 0x31,
	},
	ELEM_LEFT_QUARTER_TURN_1_TILE: &Segment{
		Type:           0x32,
		InputDegree:    DEGREE_FLAT,
		OutputDegree:   DEGREE_FLAT,
		ElevationGain:  0,
		ForwardDelta:   1,
		SidewaysDelta:  0,
		DirectionDelta: DIRECTION_90_DEG_LEFT,
	},
	ELEM_RIGHT_QUARTER_TURN_1_TILE: &Segment{
		Type:           0x33,
		InputDegree:    DEGREE_FLAT,
		OutputDegree:   DEGREE_FLAT,
		ElevationGain:  0,
		ForwardDelta:   1,
		SidewaysDelta:  0,
		DirectionDelta: DIRECTION_90_DEG_RIGHT,
	},

	ELEM_LEFT_TWIST_DOWN_TO_UP: &Segment{
		Type: 0x34,
	},
	ELEM_RIGHT_TWIST_DOWN_TO_UP: &Segment{
		Type: 0x35,
	},
	ELEM_LEFT_TWIST_UP_TO_DOWN: &Segment{
		Type: 0x36,
	},
	ELEM_RIGHT_TWIST_UP_TO_DOWN: &Segment{
		Type: 0x37,
	},

	ELEM_LEFT_CORKSCREW_UP: &Segment{
		Type: 0x3a,
	},
	ELEM_RIGHT_CORKSCREW_UP: &Segment{
		Type: 0x3b,
	},
	ELEM_LEFT_CORKSCREW_DOWN: &Segment{
		Type: 0x3c,
	},
	ELEM_RIGHT_CORKSCREW_DOWN: &Segment{
		Type: 0x3d,
	},

	ELEM_BRAKES: &Segment{
		Type: 0x63,
	},
	ELEM_BOOSTER: &Segment{
		Type: 0x64,
	},
	// This should only be used in RCT2, I think.
	ELEM_INVERTED_90_DEG_UP_TO_FLAT_QUARTER_LOOP: &Segment{
		Type: 0x65,
	},
	ELEM_LEFT_QUARTER_BANKED_HELIX_LARGE_UP: &Segment{
		Type: 0x66,
	},
	ELEM_RIGHT_QUARTER_BANKED_HELIX_LARGE_UP: &Segment{
		Type: 0x67,
	},
	ELEM_LEFT_QUARTER_BANKED_HELIX_LARGE_DOWN: &Segment{
		Type: 0x68,
	},
	ELEM_RIGHT_QUARTER_BANKED_HELIX_LARGE_DOWN: &Segment{
		Type: 0x69,
	},
	ELEM_WATERFALL: &Segment{
		Type: 0x70,
	},
	ELEM_ON_RIDE_PHOTO: &Segment{
		Type: 0x72,
	},

	ELEM_RIGHT_EIGHTH_BANK_TO_DIAG: &Segment{
		Type: 0x8a,
	},
	ELEM_LEFT_EIGHTH_BANK_TO_ORTHOGONAL: &Segment{
		Type: 0x8b,
	},

	ELEM_DIAG_25_DEG_UP: &Segment{
		Type: 0x8e,
	},

	ELEM_DIAG_25_DEG_UP_TO_FLAT: &Segment{
		Type: 0x93,
	},
	ELEM_DIAG_FLAT_TO_25_DEG_DOWN: &Segment{
		Type: 0x96,
	},
	ELEM_DIAG_25_DEG_DOWN_TO_60_DEG_DOWN: &Segment{
		Type: 0x97,
	},
	ELEM_DIAG_60_DEG_DOWN_TO_25_DEG_DOWN: &Segment{
		Type: 0x98,
	},
	ELEM_DIAG_25_DEG_DOWN_TO_FLAT: &Segment{
		Type: 0x99,
	},

	ELEM_DIAG_FLAT_TO_LEFT_BANK: &Segment{
		Type: 0x9e,
	},

	ELEM_DIAG_LEFT_BANK_TO_FLAT: &Segment{
		Type: 0xa0,
	},
	ELEM_DIAG_RIGHT_BANK_TO_FLAT: &Segment{
		Type: 0xa1,
	},
	ELEM_DIAG_LEFT_BANK_TO_25_DEG_DOWN: &Segment{
		Type: 0xa2,
	},
	ELEM_DIAG_RIGHT_BANK_TO_25_DEG_UP: &Segment{
		Type: 0xa3,
	},

	ELEM_RIGHT_LARGE_HALF_LOOP_UP: &Segment{
		Type: 0xb8,
	},
	ELEM_RIGHT_LARGE_HALF_LOOP_DOWN: &Segment{
		Type: 0xb9,
	},

	ELEM_LEFT_LARGE_HALF_LOOP_DOWN: &Segment{
		Type: 0xba,
	},

	ELEM_END_OF_RIDE: &Segment{
		Type: 0xff,
	},
}

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
	ELEM_WATERFALL                               = 0x70
	ELEM_RAPIDS                                  = 0x71
	ELEM_ON_RIDE_PHOTO                           = 0x72

	ELEM_25_DEG_DOWN_LEFT_BANKED  = 0x73
	ELEM_25_DEG_DOWN_RIGHT_BANKED = 0x74

	ELEM_WATER_SPLASH = 0x75

	ELEM_FLAT_TO_60_DEG_UP_LONG_BASE = 0x76
	ELEM_60_DEG_UP_TO_FLAT_LONG_BASE = 0x77

	ELEM_WHIRLPOOL = 0x78

	ELEM_FLAT_TO_60_DEG_DOWN_LONG_BASE = 0x79
	ELEM_60_DEG_UP_TO_FLAT_LONG_BASE   = 0x7a

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

	ELEM_END_OF_RIDE = 0xff
)
