package exe_reader

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/kevinburke/rct/bits"
	"github.com/kevinburke/rct/tracks"
)

// XXX, this doesn't correctly handle s-bends, which only move sideways by 1
// piece, I think.
func SidewaysDelta(sidewaysDeltaByte int) int {
	if bits.On(sidewaysDeltaByte, 7) {
		return -(1 + (256-sidewaysDeltaByte)>>5)
	}
	if bits.On(sidewaysDeltaByte, 6) {
		return 1 + sidewaysDeltaByte>>5
	}
	return 0
}

var reverseMap = map[tracks.DirectionDelta]string{
	tracks.DIR_STRAIGHT:     "DIR_STRAIGHT",
	tracks.DIR_45_DEG_RIGHT: "DIR_45_DEG_RIGHT",
	tracks.DIR_90_DEG_RIGHT: "DIR_90_DEG_RIGHT",
	tracks.DIR_180_DEG:      "DIR_180_DEG",
	tracks.DIR_90_DEG_LEFT:  "DIR_90_DEG_LEFT",
	tracks.DIR_45_DEG_LEFT:  "DIR_45_DEG_LEFT",
	// XXX - fix when we allow diagonals.
	//tracks.DIR_DIAGONAL:       "DIR_DIAGONAL",
	tracks.DIR_DIAGONAL_LEFT:  "DIR_DIAGONAL_LEFT",
	tracks.DIR_DIAGONAL_RIGHT: "DIR_DIAGONAL_RIGHT",
}

func GetDiagonalFromRCTStruct(b []byte) tracks.DirectionDelta {
	startingDirectionInt := int(b[0])
	startingDirection := tracks.RCTDirectionKeys[startingDirectionInt]
	endingDirectionInt := int(b[1])
	endingDirection := tracks.RCTDirectionKeys[endingDirectionInt]

	if startingDirection == tracks.DIR_DIAGONAL {
		if endingDirection == tracks.DIR_DIAGONAL {
			return tracks.DIR_DIAGONAL
		} else if endingDirection == tracks.DIR_90_DEG_RIGHT {
			return tracks.DIR_DIAGONAL_RIGHT
		} else {
			return tracks.DIR_DIAGONAL_LEFT
		}
	} else if endingDirection == tracks.DIR_DIAGONAL {
		return tracks.DIR_DIAGONAL_RIGHT
	} else {
		return endingDirection
	}
}

func ElevationDelta(positiveHeightBit int, negativeHeightBit int) int {
	if positiveHeightBit > 0 {
		return positiveHeightBit >> 3
	}
	if negativeHeightBit > 0 {
		return -(negativeHeightBit >> 3)
	}
	return 0
}

var PRINT = false

func pt(format string, val ...interface{}) {
	if PRINT {
		fmt.Printf(format, val...)
	}
}

// Print out the Go code to make up a track segment
// XXX actually build the Go datatypes instead of printing/needing to copy
// paste
func PrintValues(typ int, elementName string, diagonalByte []byte, bankByte []byte, forwardByte []byte) {

	dir := GetDiagonalFromRCTStruct(diagonalByte)
	sidewaysDelta := SidewaysDelta(int(diagonalByte[8]))
	if int(diagonalByte[8]) == 32 {
		fmt.Printf("%s: sideways(%d) = %d\n", elementName, int(diagonalByte[8]), sidewaysDelta)
	}
	negativeHeightBit := int(diagonalByte[2])
	positiveHeightBit := int(diagonalByte[4])
	elevationDelta := ElevationDelta(positiveHeightBit, negativeHeightBit)

	pt("%s: &Segment{\n", elementName)

	pt("\tType: %#x,\n", typ)
	bitVal := int(bankByte[2])
	pt("\tInputDegree: %s,\n", configTable1Map[2][bitVal])
	bitVal = int(bankByte[1])
	pt("\tOutputDegree: %s,\n", configTable1Map[1][bitVal])
	bitVal = int(bankByte[4])
	pt("\tStartingBank: %s,\n", configTable1Map[4][bitVal])
	bitVal = int(bankByte[3])
	pt("\tEndingBank: %s,\n", configTable1Map[3][bitVal])

	var forward int16
	buf := bytes.NewReader(forwardByte[6:8])
	err := binary.Read(buf, binary.LittleEndian, &forward)
	if err != nil {
		panic(err)
	}

	pt("\tForwardDelta: %d,\n", forward/-32+1)

	pt("\tDirectionDelta: %s,\n", reverseMap[dir])
	pt("\tSidewaysDelta: %d,\n", sidewaysDelta)
	pt("\tElevationDelta: %d,\n", elevationDelta)
	pt("},\n")
}

// Follows the format in TrackCoordinates
/*
	sint8 rotation_negative;	// 0x00
	sint8 rotation_positive;	// 0x01
	sint16 z_negative;			// 0x02
	sint16 z_positive;			// 0x04
	sint16 x;					// 0x06
	sint16 y;					// 0x08
*/
var configTable1Map = map[int]map[int]string{
	0: map[int]string{
		0:   "TRACK_FLAT",
		2:   "TRACK_STATION_END",
		7:   "TRACK_VERTICAL_LOOP",
		13:  "TRACK_S_BEND",
		17:  "TRACK_TWIST",
		18:  "TRACK_HALF_LOOP",
		19:  "TRACK_CORKSCREW",
		20:  "TRACK_TOWER_BASE",
		21:  "TRACK_HELIX_SMALL",
		22:  "TRACK_HELIX_LARGE",
		23:  "TRACK_HELIX_LARGE_UNBANKED",
		24:  "TRACK_BRAKES",
		26:  "TRACK_ON_RIDE_PHOTO",
		27:  "TRACK_WATER_SPLASH",
		29:  "TRACK_BARREL_ROLL",
		30:  "TRACK_POWERED_LIFT",
		31:  "TRACK_HALF_LOOP_2", // ?
		33:  "TRACK_LOG_FLUME_REVERSER",
		36:  "TRACK_WHOA_BELLY",
		43:  "TRACK_LIFT_HILL",
		46:  "TRACK_SPINNING_TUNNEL",
		47:  "TRACK_ROTATION_CONTROL_TOGGLE",
		52:  "TRACK_RAPIDS",
		152: "TRACK_WATERFALL",
		//152: "TRACK_WHIRLPOOL",
		172: "TRACK_BRAKE_FOR_DROP",
	},
	1: map[int]string{
		0:  "TRACK_NONE",
		2:  "TRACK_UP_25",
		4:  "TRACK_UP_60",
		6:  "TRACK_DOWN_25",
		8:  "TRACK_DOWN_60",
		10: "TRACK_UP_90",
		18: "TRACK_DOWN_90",
	},
	2: map[int]string{
		0:  "TRACK_NONE",
		2:  "TRACK_UP_25",
		4:  "TRACK_UP_60",
		6:  "TRACK_DOWN_25",
		8:  "TRACK_DOWN_60",
		10: "TRACK_UP_90",
		18: "TRACK_DOWN_90",
	},
	3: map[int]string{
		0:  "TRACK_BANK_NONE",
		2:  "TRACK_BANK_LEFT",
		4:  "TRACK_BANK_RIGHT",
		15: "TRACK_BANK_UPSIDE_DOWN",
	},
	4: map[int]string{
		0:  "TRACK_BANK_NONE",
		2:  "TRACK_BANK_LEFT",
		4:  "TRACK_BANK_RIGHT",
		15: "TRACK_BANK_UPSIDE_DOWN",
	},
	5: map[int]string{
		0:   "TRACK_NONE",
		64:  "TRACK_HALF_LOOP_UP",
		192: "TRACK_HALF_LOOP_DOWN",
		208: "TRACK_UNKNOWN_VERTICAL_LOOP",
		224: "TRACK_CORKSCREW_DOWN",
	},
}
