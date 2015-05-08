package main

import (
	"fmt"

	"github.com/kevinburke/rct/tracks"
)

func main() {
	// Prints out element names and their hex values
	for i, elem := range tracks.ElementNames {
		// these are auto-padded to the correct spots
		fmt.Printf("\"%s\",  // %X\n", elem, i)
	}
}
