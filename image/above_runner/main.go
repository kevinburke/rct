package main

import "github.com/kevinburke/rct/td6"
import "github.com/kevinburke/rct/image"

func main() {
	r := td6.ReadRide("/Users/kevin/code/go/src/github.com/kevinburke/rct/rides/mischief.td6")
	img := image.Draw(r)
	image.SaveToPngFile("test.png", img)
}
