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

var advanceVectorTests = []struct {
	v   Vector
	seg *tracks.Segment
	out Vector
}{
	{
		Vector{Point{0, 0, 0}, tracks.DIR_STRAIGHT},
		tracks.TS_MAP[tracks.ELEM_FLAT],
		Vector{Point{1, 0, 0}, tracks.DIR_STRAIGHT},
	},
	{
		Vector{Point{0, 0, 0}, tracks.DIR_90_DEG_LEFT},
		tracks.TS_MAP[tracks.ELEM_FLAT],
		Vector{Point{0, 1, 0}, tracks.DIR_90_DEG_LEFT},
	},
	{
		Vector{Point{0, 0, 0}, tracks.DIR_90_DEG_RIGHT},
		tracks.TS_MAP[tracks.ELEM_FLAT],
		Vector{Point{0, -1, 0}, tracks.DIR_90_DEG_RIGHT},
	},
	{
		Vector{Point{0, 0, 0}, tracks.DIR_180_DEG},
		tracks.TS_MAP[tracks.ELEM_FLAT],
		Vector{Point{-1, 0, 0}, tracks.DIR_180_DEG},
	},
	{
		Vector{Point{5, 7, 11}, tracks.DIR_STRAIGHT},
		tracks.TS_MAP[tracks.ELEM_FLAT],
		Vector{Point{6, 7, 11}, tracks.DIR_STRAIGHT},
	},
	// simple up tests
	{
		Vector{Point{0, 0, 0}, tracks.DIR_STRAIGHT},
		tracks.TS_MAP[tracks.ELEM_25_DEG_UP],
		Vector{Point{1, 0, 2}, tracks.DIR_STRAIGHT},
	},
	{
		Vector{Point{0, 0, 0}, tracks.DIR_STRAIGHT},
		tracks.TS_MAP[tracks.ELEM_60_DEG_UP],
		Vector{Point{1, 0, 8}, tracks.DIR_STRAIGHT},
	},
	{
		Vector{Point{0, 0, 0}, tracks.DIR_90_DEG_LEFT},
		tracks.TS_MAP[tracks.ELEM_25_DEG_UP],
		Vector{Point{0, 1, 2}, tracks.DIR_90_DEG_LEFT},
	},
	// down tests
	{
		Vector{Point{0, 0, 0}, tracks.DIR_STRAIGHT},
		tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN],
		Vector{Point{1, 0, -2}, tracks.DIR_STRAIGHT},
	},

	// left tests
	{
		Vector{Point{0, 0, 8}, tracks.DIR_STRAIGHT},
		tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_5_TILES],
		Vector{Point{3, 3, 8}, tracks.DIR_90_DEG_LEFT},
	},
	{
		Vector{Point{0, 0, 8}, tracks.DIR_90_DEG_LEFT},
		tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_5_TILES],
		Vector{Point{-3, 3, 8}, tracks.DIR_180_DEG},
	},
	{
		Vector{Point{0, 0, 8}, tracks.DIR_180_DEG},
		tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_5_TILES],
		Vector{Point{-3, -3, 8}, tracks.DIR_90_DEG_RIGHT},
	},
	{
		Vector{Point{0, 0, 8}, tracks.DIR_90_DEG_RIGHT},
		tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_5_TILES],
		Vector{Point{3, -3, 8}, tracks.DIR_STRAIGHT},
	},

	{
		Vector{Point{3, 4, 8}, tracks.DIR_STRAIGHT},
		tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_5_TILES],
		Vector{Point{6, 7, 8}, tracks.DIR_90_DEG_LEFT},
	},
}

func TestAdvanceVector(t *testing.T) {
	for _, tt := range advanceVectorTests {
		out := AdvanceVector(tt.v, tt.seg)
		helper := fmt.Sprintf("%s (%d, %d, %d) %v deg:", tt.seg,
			round(tt.v.Point[0]), round(tt.v.Point[1]), round(tt.v.Point[2]), tt.v.Dir)
		if out.Point != tt.out.Point {
			t.Errorf("%s expected point to be %v, was %v", helper, tt.out.Point, out.Point)
		}
		if out.Dir != tt.out.Dir {
			t.Errorf("%s expected direction to be %v, was %v", helper, tt.out.Dir, out.Dir)
		}
	}
}
