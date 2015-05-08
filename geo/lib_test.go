package geo

import (
	"testing"

	"github.com/kevinburke/rct/tracks"
)

func TestDiffFlat(t *testing.T) {
	seg := tracks.TS_MAP[tracks.ELEM_FLAT]
	out := Diff(seg, 0)
	expected := Point{1, 0, 0}
	if *out != expected {
		t.Errorf("expected straight line to be %#v, was %#v", expected, out)
	}
}

func TestDiffCurve(t *testing.T) {
	seg := tracks.TS_MAP[tracks.ELEM_RIGHT_QUARTER_TURN_5_TILES]
	out := Diff(seg, 0)
	expected := Point{3, 3, 0}
	if *out != expected {
		t.Errorf("expected curved 5-diameter track to be %#v, was %#v", expected, out)
	}
}
