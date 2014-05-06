// Draw a RCT roller coaster from above.
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

const PIECE_WIDTH = 8
const IMG_HEIGHT = 400

var RED = color.RGBA{0xff, 0x00, 0x00, 0xff}
var BLUE = color.RGBA{0x00, 0x00, 0xff, 0xff}

func getWhiteRectangle(width, height)

func Draw(r *rct.Ride) image.Image {
	rect := image.Rect(0, 0, 600, 600)
	i := image.NewRGBA(rect)
	c := color.RGBA{0xff, 0xff, 0xff, 0xff}
	draw.Draw(i, i.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	x := 300.0
	y := 300.0
	for j := 0; j < len(r.TrackData.Elements); j++ {

	}
	//width := PIECE_WIDTH * (len(r.TrackData.Elements) + 2)
	//rect := image.Rect(0, 0, width, IMG_HEIGHT)
	//i := image.NewRGBA(rect)
	//c := color.RGBA{0xff, 0xff, 0xff, 0xff}
	//draw.Draw(i, i.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	//gc := draw2d.NewGraphicContext(i)
	//x := float64(PIECE_WIDTH)
	//y := float64(IMG_HEIGHT) - 20.0
	//for j := 0; j < len(r.TrackData.Elements); j++ {
	//gc.MoveTo(x, y)
	//elem := r.TrackData.Elements[j]
	//seg := elem.Segment
	//if elem.ChainLift {
	//gc.SetStrokeColor(BLUE)
	//} else {
	//gc.SetStrokeColor(RED)
	//}
	//y -= 8.0 * float64(seg.ElevationDelta)
	//gc.LineTo(float64(x+PIECE_WIDTH), y)
	//gc.Stroke()
	//x += PIECE_WIDTH
	//}
	//return i
}

func main() {
	r := rct.ReadRide("/Users/kevin/code/go/src/github.com/kevinburke/rct-rides/rides/mischief.td6")
	img := Draw(r)
	saveToPngFile("test.png", img)
}
