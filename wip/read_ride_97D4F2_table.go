package main

import (
	"fmt"
	"os"

	rct "github.com/kevinburke/rct-rides"
)

func main() {
	f, err := os.Open(os.Getenv("HOME") + "/code/OpenRCT2/openrct2.exe")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var WIDTH = 8
	b := make([]byte, 100*WIDTH)
	// all 4 bytes here have significance on their own, somehow
	addr := 0x0057D4F2
	f.ReadAt(b, int64(addr)) // direction change stored in 2nd bit.

	for i := 0; i < len(rct.RIDENAMES); i++ {
		fmt.Printf("%2x ", i)
		fmt.Printf("%40s ", rct.RIDENAMES[i])
		fmt.Printf("%5d ", b[i*WIDTH])
		fmt.Printf("%t\n", b[i*WIDTH]&0x80 > 0)
	}
}
