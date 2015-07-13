package main

import (
	"log"
	"os"

	"github.com/kevinburke/rct/rle"
	"github.com/kevinburke/rct/td6"
)

func main() {
	ride := td6.ReadRide("rides/mine-train-gold-rush.td6")
	bits, err := td6.Marshal(ride)
	if err != nil {
		log.Fatal(err)
	}
	writer := rle.NewWriter(os.Stdout)
	writer.Write(bits)
}
