package main

import (
	"fmt"
	"os"

	rct "github.com/kevinburke/rct"
)

func main() {
	f, err := os.Open(os.Getenv("HOME") + "/code/OpenRCT2/openrct2.exe")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var WIDTH = 0x12
	b := make([]byte, 100*WIDTH)
	addr := 0x0057E3B6
	f.ReadAt(b, int64(addr))

	for i := 0; i < len(rct.RIDENAMES); i++ {
		if b[i*WIDTH] == 3 {
			fmt.Printf("\ttrue,\t")
		} else {
			fmt.Printf("\tfalse,\t")
		}
		fmt.Printf("// %02x ", i)
		fmt.Printf("%s\n", rct.RIDENAMES[i])
	}
}
