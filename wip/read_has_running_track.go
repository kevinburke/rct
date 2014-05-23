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

	var WIDTH = 0x12
	b := make([]byte, 100*WIDTH)
	addr := 0x0057E3AC
	f.ReadAt(b, int64(addr)) // direction change stored in 2nd bit.

	for i := 0; i < len(rct.RIDENAMES); i++ {
		//fmt.Printf("%2x ", i)
		//fmt.Printf("%40s ", rct.RIDENAMES[i])
		var val bool
		if b[i*WIDTH] == 20 {
			val = true
		} else {
			val = false
		}
		fmt.Printf("\t%t,\t", bool(val))
		fmt.Printf("// %x %s\n", i, rct.RIDENAMES[i])

		//fmt.Printf("%t\n", b[i*WIDTH]&0x80 > 0)
	}
}
