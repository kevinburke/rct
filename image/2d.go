// Draw a RCT roller coaster in 2d.
package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"code.google.com/p/draw2d/draw2d"
	rct "github.com/kevinburke/rct-rides"
)

func saveToPngFile(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
}

func Draw(r *rct.Ride) image.Image {
	width := 16 + 8*len(r.TrackData.Elements)
	height := 200
	rect := image.Rect(0, 0, width, height)
	i := image.NewRGBA(rect)
	c := color.RGBA{0xff, 0xff, 0xff, 0xff}
	draw.Draw(i, i.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	gc := draw2d.NewGraphicContext(i)
	gc.MoveTo(8, 4)
	for j := 0; j < len(r.TrackData.Elements); j++ {
		fmt.Println(r.TrackData.Elements[j].Segment.ElevationGain)
	}
	return i
}

func main() {
	r := rct.ReadRide("../rides/mischief.td6")
	img := Draw(r)
	saveToPngFile("test.png", img)
}
