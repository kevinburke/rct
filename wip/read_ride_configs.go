package main

import (
	//"encoding/hex"
	"fmt"
	"os"

	//"strings"
)

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

var table1 = []string{
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
	"ELEM_POWERED_LIFT", // b6
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
	//f.ReadAt(b, 0x0057d218) // supplementary ride data ?
	//f.ReadAt(b, 0x005acfa4) // all 0xff, maybe this gets writtento
	f.ReadAt(b, 0x0057d4f2)

	//rptrs := make([]byte, 30)
	//for i := int64(0); i < 70; i++ {
	//rptrs = append(rptrs, b[0])
	//}
	//for i := 0; i < len(rides)+70; i++ {
	//if i < len(rides) {
	//fmt.Printf("%50s ", rides[i])
	//fmt.Println(rptrs[i])
	//} else {
	//fmt.Printf("%50s ", "unknown")
	//fmt.Println(rptrs[i])
	//}
	//}

	var WIDTH = 8

	for i := 0; i < len(rides)+30; i++ {
		if i < len(rides) {
			fmt.Printf("%50s ", rides[i])
		} else {
			fmt.Printf("%50s ", "unknown")
		}
		//for i := 0; i < len(table1)+30; i++ {
		//if i < len(table1) {
		//fmt.Printf("%50s ", table1[i])
		//} else {
		//fmt.Printf("%50s ", "unknown")
		//}

		//fmt.Printf("%4d", int(b[i*WIDTH]))
		for j := 0; j < WIDTH; j++ {
			bijint := int(b[i*WIDTH+j])
			//fmt.Printf("%08b ", bijint)
			fmt.Printf("%4d", bijint)
			//fmt.Printf("%s ", hex.EncodeToString([]byte{byte(i)}))
			//fmt.Println(table1[i])
			//fmt.Println(strings.Repeat("=", len(table1[i])))
			//fmt.Println(" 0 1 2 3 4 5 6 7 ")
			//fmt.Println(b[i*8 : i*8+8])
			//fmt.Println(c[i*8 : i*8+8])
			//printValues(i, b)
			//fmt.Println("")
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

func printValues(i int, b []byte) {
	fmt.Printf("%s: &Segment{\n", table1[i])
	fmt.Printf("\tType: %#x,\n", i)
	bitVal := int(b[i*8+2])
	fmt.Printf("\tInputDegree: %s,\n", foo[2][bitVal])
	bitVal = int(b[i*8+1])
	fmt.Printf("\tOutputDegree: %s,\n", foo[1][bitVal])
	bitVal = int(b[i*8+4])
	fmt.Printf("\tStartingBank: %s,\n", foo[4][bitVal])
	bitVal = int(b[i*8+3])
	fmt.Printf("\tEndingBank: %s,\n", foo[3][bitVal])
	fmt.Printf("},\n")
	//fmt.Printf("\t{ ")

	//for j := 0; j < 6; j++ {
	//bitVal := int(b[i*8+j])
	//fmt.Printf("%s,\t\t", foo[j][bitVal])
	//}

	//fmt.Printf(" },\t\t// %s\n", table1[i])

	//if b[i*8+2] == 2 {
	//fmt.Println("- starts at 25 degrees up")
	//}
	//if b[i*8+2] == 4 {
	//fmt.Println("- starts at 60 degrees up")
	//}
	//if b[i*8+2] == 6 {
	//fmt.Println("- starts at 25 degrees down")
	//}
	//if b[i*8+2] == 8 {
	//fmt.Println("- starts at 60 degrees down")
	//}
	//if b[i*8+1] == 0 {
	//fmt.Println("- ends flat")
	//}
	//if b[i*8+1] == 2 {
	//fmt.Println("- ends at 25 degrees up")
	//}
	//if b[i*8+1] == 4 {
	//fmt.Println("- ends at 60 degrees up")
	//}
	//if b[i*8+1] == 6 {
	//fmt.Println("- ends at 25 degrees down")
	//}
	//if b[i*8+1] == 8 {
	//fmt.Println("- ends at 60 degrees down")
	//}
	//if b[i*8+4] == 2 {
	//fmt.Println("- starts with left bank")
	//}
	//if b[i*8+4] == 4 {
	//fmt.Println("- starts with right bank")
	//}
	//if b[i*8+3] == 2 {
	//fmt.Println("- ends with left bank")
	//}
	//if b[i*8+3] == 4 {
	//fmt.Println("- ends with right bank")
	//}
}

// keys
// byte 0:
// value 2: end station
// value 7: vertical loop
// value 13: s bend
// 17: twist
// 18: half loop
// 19: corkscrew
// 20: tower base
// 21: small helix
// 22: large helix
// 23: unbanked large helix
// 24: brakes
// 26: on ride photo
// 27: water splash
// 29: barrel roll
// 30: powered lift
// 31: half loop
// 33: log flume reverser
// 36: whoa belly
// 43: lift hill
// 46: spinning tunnel
// 47: rotation control toggle
// 52: rapids (rct2 only)
// 152: waterfall/whirlpool
// 172: brake for drop

// byte 1:
// 2: ends at 25 degree up
// 4: ends at 60 degree up
// 6: ends at 25 degree down
// 8: ends at 60 degree down
// 10: 90 degree up (tower, whoa belly)
// 18: 90 degree down

// byte 2:
// 2: starts at 25 degree up
// 4: starts at 60 degree up
// 6: starts at 25 degree down
// 8: starts at 60 degree down
// 10: 90 degree up (also tower, whoa belly)
// 18: 90 degree down

// byte 3:
// 2: ends with left bank
// 4: ends with right bank
// 15: ends upside down

// byte 4:
// 2: starts with L bank
// 4: starts with R bank
// 15: starts upside down

// byte 5:
// 64: half loop up
// 192: half loop down
// 208: something relating to vertical loops. same for both L and R
// 224: corkscrew down

// odd things:
// - no instruction for diagonal
// - no instruction for direction change
// - no instruction for diameter
