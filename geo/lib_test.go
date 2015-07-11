package geo

import (
	"fmt"
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
	expected := Point{3, -3, 0}
	if *out != expected {
		t.Errorf("expected curved 5-diameter track to be %#v, was %#v", expected, out)
	}
}

func TestDiffFlatRotated(t *testing.T) {
	seg := tracks.TS_MAP[tracks.ELEM_FLAT]
	out := Diff(seg, 90)
	expected := Point{0, 1, 0}
	if *out != expected {
		t.Errorf("expected flat track to be %#v, was %#v", expected, out)
	}

	out = Diff(seg, 180)
	expected = Point{-1, 0, 0}
	if *out != expected {
		t.Errorf("expected flat track to be %#v, was %#v", expected, out)
	}

	out = Diff(seg, 270)
	expected = Point{0, -1, 0}
	if *out != expected {
		t.Errorf("expected flat track to be %#v, was %#v", expected, out)
	}
}

var advanceTests = []struct {
	seg       *tracks.Segment
	e         int
	forward   int
	sideways  int
	direction tracks.DirectionDelta
	e_out     int
	f_out     int
	s_out     int
	d_out     tracks.DirectionDelta
}{
	{tracks.TS_MAP[tracks.ELEM_FLAT], 0, 0, 0, tracks.DIR_STRAIGHT, 0, 1, 0, tracks.DIR_STRAIGHT},
	{tracks.TS_MAP[tracks.ELEM_FLAT], 0, 0, 0, tracks.DIR_90_DEG_RIGHT, 0, 0, -1, tracks.DIR_90_DEG_RIGHT},
	{tracks.TS_MAP[tracks.ELEM_FLAT], 0, 0, 0, tracks.DIR_180_DEG, 0, -1, 0, tracks.DIR_180_DEG},
	{tracks.TS_MAP[tracks.ELEM_60_DEG_UP], 0, 0, 0, tracks.DIR_STRAIGHT, 8, 1, 0, tracks.DIR_STRAIGHT},
	{tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_5_TILES], 0, 0, 0, tracks.DIR_STRAIGHT, 0, 3, 3, tracks.DIR_90_DEG_LEFT},
	{tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_5_TILES], 0, 0, 0, tracks.DIR_90_DEG_LEFT, 0, 3, -3, tracks.DIR_180_DEG},
	{tracks.TS_MAP[tracks.ELEM_RIGHT_QUARTER_TURN_5_TILES], 0, 0, 0, tracks.DIR_STRAIGHT, 0, 3, -3, tracks.DIR_90_DEG_RIGHT},
}

func TestAdvance(t *testing.T) {
	t.Skip("These tests don't work for turns...")
	for _, tt := range advanceTests {
		eout, fout, sout, dout := Advance(tt.seg, tt.e, tt.forward, tt.sideways, tt.direction)
		header := fmt.Sprintf("%s (init %d, %d, %d, %s):", tt.seg, tt.e, tt.forward, tt.sideways, tt.direction)
		if eout != tt.e_out {
			t.Errorf("%s expected %d elevation change, got %d", header, tt.e_out, eout)
		}
		if fout != tt.f_out {
			t.Errorf("%s expected %d forward delta, got %d", header, tt.f_out, fout)
		}
		if sout != tt.s_out {
			t.Errorf("%s expected %d sideways delta, got %d", header, tt.s_out, sout)
		}
		if dout != tt.d_out {
			t.Errorf("%s expected to go %s, got %d", header, tt.d_out, dout)
		}
	}
}

func TestFlat90DegRotated(t *testing.T) {
	v := Vector{Point{3, 7, 0}, tracks.DIR_90_DEG_RIGHT}
	seg := tracks.TS_MAP[tracks.ELEM_FLAT]
	out := AdvanceVector(v, seg)
	expected := Point{3, 8, 0}
	if out.Point != expected {
		t.Errorf("expected upward facing track to advance to %#v, got %#v", expected, out.Point)
	}
}

func TestFlat180DegRotated(t *testing.T) {
	v := Vector{Point{3, 7, 0}, tracks.DIR_180_DEG}
	seg := tracks.TS_MAP[tracks.ELEM_FLAT]
	out := AdvanceVector(v, seg)
	expected := Point{2, 7, 0}
	if out.Point != expected {
		t.Errorf("expected upward facing track to advance to %#v, got %#v", expected, out.Point)
	}
}

func TestFlat270DegRotated(t *testing.T) {
	v := Vector{Point{3, 7, 0}, tracks.DIR_90_DEG_LEFT}
	seg := tracks.TS_MAP[tracks.ELEM_FLAT]
	out := AdvanceVector(v, seg)
	expected := Point{3, 6, 0}
	if out.Point != expected {
		t.Errorf("expected upward facing track to advance to %#v, got %#v", expected, out.Point)
	}
}
