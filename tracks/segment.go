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

	ELEM_LEFT_QUARTER_TURN_3_TILES_BANK = 0x2c

	ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_DOWN  = 0x30
	ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_DOWN = 0x31
	ELEM_LEFT_QUARTER_TURN_1_TILE               = 0x32
	ELEM_RIGHT_QUARTER_TURN_1_TILE              = 0x33

	ELEM_LEFT_TWIST_DOWN_TO_UP  = 0x34
	ELEM_RIGHT_TWIST_DOWN_TO_UP = 0x35
	ELEM_LEFT_TWIST_UP_TO_DOWN  = 0x36
	ELEM_RIGHT_TWIST_UP_TO_DOWN = 0x37

	ELEM_LEFT_CORKSCREW_UP    = 0x3a
	ELEM_RIGHT_CORKSCREW_UP   = 0x3b
	ELEM_LEFT_CORKSCREW_DOWN  = 0x3c
	ELEM_RIGHT_CORKSCREW_DOWN = 0x3d

	ELEM_BRAKES  = 0x63
	ELEM_BOOSTER = 0x64
	// This should only be used in RCT2, I think.
	ELEM_INVERTED_90_DEG_UP_TO_FLAT_QUARTER_LOOP = 0x65
	ELEM_LEFT_QUARTER_BANKED_HELIX_LARGE_UP      = 0x66
	ELEM_RIGHT_QUARTER_BANKED_HELIX_LARGE_UP     = 0x67
	ELEM_LEFT_QUARTER_BANKED_HELIX_LARGE_DOWN    = 0x68
	ELEM_RIGHT_QUARTER_BANKED_HELIX_LARGE_DOWN   = 0x69
	ELEM_WATERFALL                               = 0x70
	ELEM_ON_RIDE_PHOTO                           = 0x72

	ELEM_RIGHT_EIGHTH_BANK_TO_DIAG      = 0x8a
	ELEM_LEFT_EIGHTH_BANK_TO_ORTHOGONAL = 0x8b

	ELEM_DIAG_25_DEG_UP = 0x8e
	ELEM_DIAG_60_DEG_UP = 0x8e

	ELEM_DIAG_25_DEG_UP_TO_FLAT          = 0x93
	ELEM_DIAG_25_DEG_DOWN                = 0x94
	ELEM_DIAG_60_DEG_DOWN                = 0x95
	ELEM_DIAG_FLAT_TO_25_DEG_DOWN        = 0x96
	ELEM_DIAG_25_DEG_DOWN_TO_60_DEG_DOWN = 0x97
	ELEM_DIAG_60_DEG_DOWN_TO_25_DEG_DOWN = 0x98
	ELEM_DIAG_25_DEG_DOWN_TO_FLAT        = 0x99

	ELEM_DIAG_FLAT_TO_LEFT_BANK = 0x9e

	ELEM_DIAG_LEFT_BANK_TO_FLAT        = 0xa0
	ELEM_DIAG_RIGHT_BANK_TO_FLAT       = 0xa1
	ELEM_DIAG_LEFT_BANK_TO_25_DEG_DOWN = 0xa2
	ELEM_DIAG_RIGHT_BANK_TO_25_DEG_UP  = 0xa3

	ELEM_RIGHT_LARGE_HALF_LOOP_UP   = 0xb8
	ELEM_RIGHT_LARGE_HALF_LOOP_DOWN = 0xb9

	ELEM_LEFT_LARGE_HALF_LOOP_DOWN = 0xba

	ELEM_END_OF_RIDE = 0xff
)
