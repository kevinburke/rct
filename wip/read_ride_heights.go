package main

import (
	//"encoding/hex"

	"encoding/binary"
	"fmt"
	"os"
)

//"strings"

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

var elementNames = []string{
	"ELEM_FLAT",
	"ELEM_END_STATION",
	"ELEM_BEGIN_STATION",
	"ELEM_MIDDLE_STATION",
	"ELEM_25_DEG_UP",
	"ELEM_60_DEG_UP",
	"ELEM_FLAT_TO_25_DEG_UP",
	"ELEM_25_DEG_UP_TO_60_DEG_UP",
	"ELEM_60_DEG_UP_TO_25_DEG_UP",
	"ELEM_25_DEG_UP_TO_FLAT",
	"ELEM_25_DEG_DOWN",
	"ELEM_60_DEG_DOWN",
	"ELEM_FLAT_TO_25_DEG_DOWN",
	"ELEM_25_DEG_DOWN_TO_60_DEG_DOWN",
	"ELEM_60_DEG_DOWN_TO_25_DEG_DOWN",
	"ELEM_25_DEG_DOWN_TO_FLAT",
	"ELEM_LEFT_QUARTER_TURN_5_TILES",
	"ELEM_RIGHT_QUARTER_TURN_5_TILES",
	"ELEM_FLAT_TO_LEFT_BANK",
	"ELEM_FLAT_TO_RIGHT_BANK",
	"ELEM_LEFT_BANK_TO_FLAT",
	"ELEM_RIGHT_BANK_TO_FLAT",
	"ELEM_BANKED_LEFT_QUARTER_TURN_5_TILES",
	"ELEM_BANKED_RIGHT_QUARTER_TURN_5_TILES",
	"ELEM_LEFT_BANK_TO_25_DEG_UP",
	"ELEM_RIGHT_BANK_TO_25_DEG_UP",
	"ELEM_25_DEG_UP_TO_LEFT_BANK",
	"ELEM_25_DEG_UP_TO_RIGHT_BANK",
	"ELEM_LEFT_BANK_TO_25_DEG_DOWN",
	"ELEM_RIGHT_BANK_TO_25_DEG_DOWN",
	"ELEM_25_DEG_DOWN_TO_LEFT_BANK",
	"ELEM_25_DEG_DOWN_TO_RIGHT_BANK",
	"ELEM_LEFT_BANK",
	"ELEM_RIGHT_BANK",
	"ELEM_LEFT_QUARTER_TURN_5_TILES_25_DEG_UP",
	"ELEM_RIGHT_QUARTER_TURN_5_TILES_25_DEG_UP",
	"ELEM_LEFT_QUARTER_TURN_5_TILES_25_DEG_DOWN",
	"ELEM_RIGHT_QUARTER_TURN_5_TILES_25_DEG_DOWN",
	"ELEM_S_BEND_LEFT",
	"ELEM_S_BEND_RIGHT",
	"ELEM_LEFT_VERTICAL_LOOP",
	"ELEM_RIGHT_VERTICAL_LOOP",
	"ELEM_LEFT_QUARTER_TURN_3_TILES",
	"ELEM_RIGHT_QUARTER_TURN_3_TILES",
	"ELEM_LEFT_QUARTER_TURN_3_TILES_BANK",
	"ELEM_RIGHT_QUARTER_TURN_3_TILES_BANK",
	"ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_UP",
	"ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_UP",
	"ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_DOWN",
	"ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_DOWN",
	"ELEM_LEFT_QUARTER_TURN_1_TILE",
	"ELEM_RIGHT_QUARTER_TURN_1_TILE",
	"ELEM_LEFT_TWIST_DOWN_TO_UP",
	"ELEM_RIGHT_TWIST_DOWN_TO_UP",
	"ELEM_LEFT_TWIST_UP_TO_DOWN",
	"ELEM_RIGHT_TWIST_UP_TO_DOWN",
	"ELEM_HALF_LOOP_UP",
	"ELEM_HALF_LOOP_DOWN",
	"ELEM_LEFT_CORKSCREW_UP",
	"ELEM_RIGHT_CORKSCREW_UP",
	"ELEM_LEFT_CORKSCREW_DOWN",
	"ELEM_RIGHT_CORKSCREW_DOWN",
	"ELEM_FLAT_TO_60_DEG_UP",
	"ELEM_60_DEG_UP_TO_FLAT",
	"ELEM_FLAT_TO_60_DEG_DOWN",
	"ELEM_60_DEG_DOWN_TO_FLAT",
	"ELEM_TOWER_BASE",
	"ELEM_TOWER_SECTION",
	"ELEM_FLAT_COVERED",
	"ELEM_25_DEG_UP_COVERED",
	"ELEM_60_DEG_UP_COVERED",
	"ELEM_FLAT_TO_25_DEG_UP_COVERED",
	"ELEM_25_DEG_UP_TO_60_DEG_UP_COVERED",
	"ELEM_60_DEG_UP_TO_25_DEG_UP_COVERED",
	"ELEM_25_DEG_UP_TO_FLAT_COVERED",
	"ELEM_25_DEG_DOWN_COVERED",
	"ELEM_60_DEG_DOWN_COVERED",
	"ELEM_FLAT_TO_25_DEG_DOWN_COVERED",
	"ELEM_25_DEG_DOWN_TO_60_DEG_DOWN_COVERED",
	"ELEM_60_DEG_DOWN_TO_25_DEG_DOWN_COVERED",
	"ELEM_25_DEG_DOWN_TO_FLAT_COVERED",
	"ELEM_LEFT_QUARTER_TURN_5_TILES_COVERED",
	"ELEM_RIGHT_QUARTER_TURN_5_TILES_COVERED",
	"ELEM_S_BEND_LEFT_COVERED",
	"ELEM_S_BEND_RIGHT_COVERED",
	"ELEM_LEFT_QUARTER_TURN_3_TILES_COVERED",
	"ELEM_RIGHT_QUARTER_TURN_3_TILES_COVERED",
	"ELEM_LEFT_HALF_BANKED_HELIX_UP_SMALL",
	"ELEM_RIGHT_HALF_BANKED_HELIX_UP_SMALL",
	"ELEM_LEFT_HALF_BANKED_HELIX_DOWN_SMALL",
	"ELEM_RIGHT_HALF_BANKED_HELIX_DOWN_SMALL",
	"ELEM_LEFT_HALF_BANKED_HELIX_UP_LARGE",
	"ELEM_RIGHT_HALF_BANKED_HELIX_UP_LARGE",
	"ELEM_LEFT_HALF_BANKED_HELIX_DOWN_LARGE",
	"ELEM_RIGHT_HALF_BANKED_HELIX_DOWN_LARGE",
	"ELEM_LEFT_QUARTER_TURN_1_TILE_60_DEG_UP",
	"ELEM_RIGHT_QUARTER_TURN_1_TILE_60_DEG_UP",
	"ELEM_LEFT_QUARTER_TURN_1_TILE_60_DEG_DOWN",
	"ELEM_RIGHT_QUARTER_TURN_1_TILE_60_DEG_DOWN",
	"ELEM_BRAKES",
	"ELEM_ROTATION_CONTROL_TOGGLE",
	"ELEM_INVERTED_90_DEG_UP_TO_FLAT_QUARTER_LOOP",
	"ELEM_LEFT_QUARTER_BANKED_HELIX_LARGE_UP",
	"ELEM_RIGHT_QUARTER_BANKED_HELIX_LARGE_UP",
	"ELEM_LEFT_QUARTER_BANKED_HELIX_LARGE_DOWN",
	"ELEM_RIGHT_QUARTER_BANKED_HELIX_LARGE_DOWN",
	"ELEM_LEFT_QUARTER_HELIX_LARGE_UP",
	"ELEM_RIGHT_QUARTER_HELIX_LARGE_UP",
	"ELEM_LEFT_QUARTER_HELIX_LARGE_DOWN",
	"ELEM_RIGHT_QUARTER_HELIX_LARGE_DOWN",
	"ELEM_25_DEG_UP_LEFT_BANKED",
	"ELEM_25_DEG_UP_RIGHT_BANKED",
	"ELEM_WATERFALL",
	"ELEM_RAPIDS",
	"ELEM_ON_RIDE_PHOTO",
	"ELEM_25_DEG_DOWN_LEFT_BANKED",
	"ELEM_25_DEG_DOWN_RIGHT_BANKED",
	"ELEM_WATER_SPLASH",
	"ELEM_FLAT_TO_60_DEG_UP_LONG_BASE",
	"ELEM_60_DEG_UP_TO_FLAT_LONG_BASE",
	"ELEM_WHIRLPOOL",
	"ELEM_FLAT_TO_60_DEG_DOWN_LONG_BASE",
	"ELEM_60_DEG_UP_TO_FLAT_LONG_BASE",
	"ELEM_CABLE_LIFT_HILL",
	"ELEM_REVERSE_WHOA_BELLY_SLOPE",
	"ELEM_REVERSE_WHOA_BELLY_VERTICAL",
	"ELEM_90_DEG_UP",
	"ELEM_90_DEG_DOWN",
	"ELEM_60_DEG_UP_TO_90_DEG_UP",
	"ELEM_90_DEG_DOWN_TO_60_DEG_DOWN",
	"ELEM_90_DEG_UP_TO_60_DEG_UP",
	"ELEM_60_DEG_DOWN_TO_90_DEG_DOWN",
	"ELEM_BRAKE_FOR_DROP",
	"ELEM_LEFT_EIGHTH_TO_DIAG",
	"ELEM_RIGHT_EIGHTH_TO_DIAG",
	"ELEM_LEFT_EIGHTH_TO_ORTHOGONAL",
	"ELEM_RIGHT_EIGHTH_TO_ORTHOGONAL",
	"ELEM_LEFT_EIGHTH_BANK_TO_DIAG",
	"ELEM_RIGHT_EIGHTH_BANK_TO_DIAG",
	"ELEM_LEFT_EIGHTH_BANK_TO_ORTHOGONAL",
	"ELEM_RIGHT_EIGHTH_BANK_TO_ORTHOGONAL",
	"ELEM_DIAG_FLAT",
	"ELEM_DIAG_25_DEG_UP",
	"ELEM_DIAG_60_DEG_UP",
	"ELEM_DIAG_FLAT_TO_25_DEG_UP",
	"ELEM_DIAG_25_DEG_UP_TO_60_DEG_UP",
	"ELEM_DIAG_60_DEG_UP_TO_25_DEG_UP",
	"ELEM_DIAG_25_DEG_UP_TO_FLAT",
	"ELEM_DIAG_25_DEG_DOWN",
	"ELEM_DIAG_60_DEG_DOWN",
	"ELEM_DIAG_FLAT_TO_25_DEG_DOWN",
	"ELEM_DIAG_25_DEG_DOWN_TO_60_DEG_DOWN",
	"ELEM_DIAG_60_DEG_DOWN_TO_25_DEG_DOWN",
	"ELEM_DIAG_25_DEG_DOWN_TO_FLAT",
	"ELEM_DIAG_FLAT_TO_60_DEG_UP",
	"ELEM_DIAG_60_DEG_UP_TO_FLAT",
	"ELEM_DIAG_FLAT_TO_60_DEG_DOWN",
	"ELEM_DIAG_60_DEG_DOWN_TO_FLAT",
	"ELEM_DIAG_FLAT_TO_LEFT_BANK",
	"ELEM_DIAG_FLAT_TO_RIGHT_BANK",
	"ELEM_DIAG_LEFT_BANK_TO_FLAT",
	"ELEM_DIAG_RIGHT_BANK_TO_FLAT",
	"ELEM_DIAG_LEFT_BANK_TO_25_DEG_UP",
	"ELEM_DIAG_RIGHT_BANK_TO_25_DEG_UP",
	"ELEM_DIAG_25_DEG_UP_TO_LEFT_BANK",
	"ELEM_DIAG_25_DEG_UP_TO_RIGHT_BANK",
	"ELEM_DIAG_LEFT_BANK_TO_25_DEG_DOWN",
	"ELEM_DIAG_RIGHT_BANK_TO_25_DEG_DOWN",
	"ELEM_DIAG_25_DEG_DOWN_TO_LEFT_BANK",
	"ELEM_DIAG_25_DEG_DOWN_TO_RIGHT_BANK",
	"ELEM_DIAG_LEFT_BANK",
	"ELEM_DIAG_RIGHT_BANK",
	"ELEM_LOG_FLUME_REVERSER",
	"ELEM_SPINNING_TUNNEL",
	"ELEM_LEFT_BARREL_ROLL_UP_TO_DOWN",
	"ELEM_RIGHT_BARREL_ROLL_UP_TO_DOWN",
	"ELEM_LEFT_BARREL_ROLL_DOWN_TO_UP",
	"ELEM_RIGHT_BARREL_ROLL_DOWN_TO_UP",
	"ELEM_LEFT_BANK_TO_LEFT_QUARTER_TURN_3_TILES_25_DEG_UP",
	"ELEM_RIGHT_BANK_TO_RIGHT_QUARTER_TURN_3_TILES_25_DEG_UP",
	"ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_DOWN_TO_LEFT_BANK",
	"ELEM_RIGHT_QUARTER_TURN_3_TILES_25_DEG_DOWN_TO_RIGHT_BANK",
	"ELEM_POWERED_LIFT",
	"ELEM_LEFT_LARGE_HALF_LOOP_UP",
	"ELEM_RIGHT_LARGE_HALF_LOOP_UP",
	"ELEM_RIGHT_LARGE_HALF_LOOP_DOWN",
	"ELEM_LEFT_LARGE_HALF_LOOP_DOWN",
}

var rides = []string{
	"Spiral Roller coaster",
	"Stand Up Coaster",
	"Suspended Swinging",
	"Inverted",
	"Steel Mini Coaster",
	"Mini Railroad",
	"Monorail",
	"Mini Suspended Coaster",
	"Bumper Boats",
	"Wooden Wild Mine/Mouse",
	"Steeplechase/Motorbike/Soap Box Derby",
	"Car Ride",
	"Launched Freefall",
	"Bobsleigh Coaster",
	"Flying Turns",
	"Observation Tower",
	"Looping Roller Coaster",
	"Dinghy Slide",
	"Mine Train Coaster",
	"Chairlift",
	"Corkscrew Roller Coaster",
	"Maze",
	"Spiral Slide",
	"Go Karts",
}

func main() {
	f, err := os.Open(os.Getenv("HOME") + "/code/OpenRCT2/openrct2.exe")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b := make([]byte, 1800)
	//f.ReadAt(b, 0x00594638) // ride height data ?
	f.ReadAt(b, 0x00594A38) // ride height data ?

	var WIDTH = 11
	var colHeaderFmt = "%55s "
	// header row
	fmt.Printf(colHeaderFmt, "Number")
	for j := 0; j < WIDTH; j++ {
		fmt.Printf("%4d", j)
	}
	fmt.Printf("\n")

	for i := 0; i < 256; i++ {
		if i < len(elementNames) {
			fmt.Printf(colHeaderFmt, elementNames[i])
		} else {
			fmt.Printf(colHeaderFmt, "unknown")
		}
		dataBit := binary.LittleEndian.Uint32(b[i*4 : i*4+4])
		exeAddress := dataBit - 0x400000
		c := make([]byte, WIDTH)
		f.ReadAt(c, int64(exeAddress))
		for j := 0; j < WIDTH; j++ {
			bijint := int(c[j])
			fmt.Printf("%4d", bijint)
		}
		fmt.Printf("\n")
	}
}

var foo = map[int]map[int]string{
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
