package main

import "github.com/kevinburke/rct-rides/td6"
import "github.com/kevinburke/rct-rides/image"

func main() {
	r := td6.ReadRide("/Users/kevin/code/go/src/github.com/kevinburke/rct-rides/rides/mischief.td6")
	img := image.Draw(r)
	image.SaveToPngFile("test.png", img)
}
