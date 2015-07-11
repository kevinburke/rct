package main

import (
	"os"

	"github.com/kevinburke/rct/exe_reader"
	"github.com/kevinburke/rct/tracks"
)

const RCT_DIRECTION_ADDR = 0x005972bb
const RCT_DIRECTION_WIDTH = 10

const RCT_BANK_SLOPE_ADDR = 0x00597c9d
const RCT_BANK_SLOPE_WIDTH = 8

// Follows the format in TrackCoordinates
/*
	sint8 rotation_negative;	// 0x00
	sint8 rotation_positive;	// 0x01
	sint16 z_negative;			// 0x02
	sint16 z_positive;			// 0x04
	sint16 x;					// 0x06
	sint16 y;					// 0x08
*/
const RCT_FORWARD_ADDR = 0x005968bb
const RCT_FORWARD_WIDTH = 0x0A

func main() {
	f, err := os.Open(os.Getenv("HOME") + "/code/OpenRCT2/openrct2.exe")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	b := make([]byte, 256*RCT_DIRECTION_WIDTH)
	f.ReadAt(b, int64(RCT_DIRECTION_ADDR))

	c := make([]byte, 256*RCT_BANK_SLOPE_WIDTH)
	f.ReadAt(c, RCT_BANK_SLOPE_ADDR)

	d := make([]byte, 256*RCT_FORWARD_WIDTH)
	f.ReadAt(d, RCT_FORWARD_ADDR)

	for i := 0; i < len(tracks.ElementNames); i++ {
		//fmt.Println(i)
		//fmt.Printf("%55s ", tracks.ElementNames[i])
		//fmt.Printf("%4d ", b[i*WIDTH])
		//fmt.Printf("\n")
		idx := i * RCT_DIRECTION_WIDTH
		bitSubset := b[idx : idx+RCT_DIRECTION_WIDTH]

		bankIdx := i * RCT_BANK_SLOPE_WIDTH
		bankBitSubset := c[bankIdx : bankIdx+RCT_BANK_SLOPE_WIDTH]

		forwardIdx := i * RCT_FORWARD_WIDTH
		forwardBitSubset := d[forwardIdx : forwardIdx+RCT_FORWARD_WIDTH]

		exe_reader.PrintValues(i, tracks.ElementNames[i], bitSubset, bankBitSubset, forwardBitSubset)
	}

	//fmt.Printf("%#v\n", tracks.TS_MAP)
	//fmt.Printf("%T\n", tracks.TS_MAP)
}
