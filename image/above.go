// Draw a RCT roller coaster from above.
package image

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/kevinburke/rct/td6"
)

func getWhiteRectangle(width int, height int) *image.RGBA {
	rect := image.Rect(0, 0, 600, 600)
	i := image.NewRGBA(rect)
	c := color.RGBA{0xff, 0xff, 0xff, 0xff}
	draw.Draw(i, i.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	return i
}

func DrawAbove(r *td6.Ride) image.Image {
	i := getWhiteRectangle(600, 600)
	x := 300.0
	y := 300.0
	fmt.Println(x)
	fmt.Println(y)
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

	// not correct.
	return i
}
