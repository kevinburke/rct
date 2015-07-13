package main

import (
	"fmt"

	"github.com/kevinburke/rct/td6"
)

func main() {
	ride := td6.ReadRide("rides/mine-train-gold-rush.td6")
	fmt.Printf("%#v\n", ride)
}
