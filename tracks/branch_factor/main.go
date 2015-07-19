package main

import (
	"fmt"

	"github.com/kevinburke/rct/tracks"
)

func main() {
	for i := range tracks.ElementNames {
		seg := tracks.TS_MAP[tracks.SegmentType(i)]
		elem := tracks.Element{
			Segment: seg,
		}
		if !tracks.Valid(seg) {
			continue
		}
		fmt.Println(len(elem.Possibilities()))
	}
}
