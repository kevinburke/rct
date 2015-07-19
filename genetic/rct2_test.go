package genetic

import (
	"fmt"
	"testing"

	"github.com/kevinburke/rct/geo"
	"github.com/kevinburke/rct/tracks"
)

func TestStraightConnector(t *testing.T) {
	trackEnd := geo.Vector{geo.Point{
		3, 4, 2,
	}, tracks.DIR_STRAIGHT}
	stationStart := geo.Vector{geo.Point{
		5, 4, 2,
	}, tracks.DIR_STRAIGHT}
	elems := completeTrack(trackEnd, stationStart)
	if elems[0].Segment != tracks.TS_MAP[tracks.ELEM_FLAT] {
		t.Fatalf("expected two straight flat pieces, got %#v", elems[0].Segment)
	}
	if elems[1].Segment != tracks.TS_MAP[tracks.ELEM_FLAT] {
		t.Fatalf("expected two straight flat pieces, got %#v", elems[1].Segment)
	}
}

func TestSimpleDescender(t *testing.T) {
	trackEnd := geo.Vector{geo.Point{
		3, 4, 3,
	}, tracks.DIR_STRAIGHT}
	stationStart := geo.Vector{geo.Point{
		5, 4, 2,
	}, tracks.DIR_STRAIGHT}
	elems := completeTrack(trackEnd, stationStart)
	if elems[0].Segment != tracks.TS_MAP[tracks.ELEM_FLAT_TO_25_DEG_DOWN] {
		t.Fatalf("expected a down piece, got %+v", elems[0])
	}
	if elems[1].Segment != tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN_TO_FLAT] {
		t.Fatalf("expected a down-to-flat piece, got %+v", elems[1])
	}
}

func TestOneDownhillDescender(t *testing.T) {
	trackEnd := geo.Vector{geo.Point{
		2, 4, 4,
	}, tracks.DIR_STRAIGHT}
	stationStart := geo.Vector{geo.Point{
		5, 4, 2,
	}, tracks.DIR_STRAIGHT}
	elems := completeTrack(trackEnd, stationStart)
	if elems[0].Segment != tracks.TS_MAP[tracks.ELEM_FLAT_TO_25_DEG_DOWN] {
		t.Fatalf("expected a down piece, got %+v", elems[0])
	}
	if elems[1].Segment != tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN] {
		t.Fatalf("expected a down-to-flat piece, got %+v", elems[1])
	}
	if elems[2].Segment != tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN_TO_FLAT] {
		t.Fatalf("expected a down-to-flat piece, got %+v", elems[1])
	}
}

func buildTrack(names []tracks.SegmentType) []tracks.Element {
	te := make([]tracks.Element, len(names))
	for i, name := range names {
		te[i] = tracks.Element{Segment: tracks.TS_MAP[name]}
	}
	return te
}

func buildOneTrack(name tracks.SegmentType) tracks.Element {
	return buildTrack([]tracks.SegmentType{name})[0]
}

var trackTests = []struct {
	in       geo.Vector
	expected []tracks.Element
}{
	{geo.Vector{geo.Point{-2, 2, 0}, tracks.DIR_90_DEG_RIGHT},
		buildTrack([]tracks.SegmentType{tracks.ELEM_LEFT_QUARTER_TURN_3_TILES})},

	{geo.Vector{geo.Point{-2, -2, 0}, tracks.DIR_90_DEG_LEFT},
		buildTrack([]tracks.SegmentType{tracks.ELEM_RIGHT_QUARTER_TURN_3_TILES})},

	{geo.Vector{geo.Point{-2, -8, 0}, tracks.DIR_90_DEG_LEFT},
		buildTrack([]tracks.SegmentType{
			tracks.ELEM_FLAT,
			tracks.ELEM_FLAT,
			tracks.ELEM_FLAT,
			tracks.ELEM_FLAT,
			tracks.ELEM_FLAT,
			tracks.ELEM_FLAT,
			tracks.ELEM_RIGHT_QUARTER_TURN_3_TILES,
		})},

	{geo.Vector{geo.Point{-2, 8, 0}, tracks.DIR_90_DEG_RIGHT},
		buildTrack([]tracks.SegmentType{
			tracks.ELEM_FLAT,
			tracks.ELEM_FLAT,
			tracks.ELEM_FLAT,
			tracks.ELEM_FLAT,
			tracks.ELEM_FLAT,
			tracks.ELEM_FLAT,
			tracks.ELEM_LEFT_QUARTER_TURN_3_TILES,
		})},

	{geo.Vector{geo.Point{-1, 0, 0}, tracks.DIR_STRAIGHT},
		buildTrack([]tracks.SegmentType{tracks.ELEM_FLAT})},

	{geo.Vector{geo.Point{-3, 0, 0}, tracks.DIR_STRAIGHT},
		buildTrack([]tracks.SegmentType{tracks.ELEM_FLAT, tracks.ELEM_FLAT, tracks.ELEM_FLAT})},

	{geo.Vector{geo.Point{0, 4, 0}, tracks.DIR_180_DEG},
		buildTrack([]tracks.SegmentType{tracks.ELEM_LEFT_QUARTER_TURN_3_TILES, tracks.ELEM_LEFT_QUARTER_TURN_3_TILES})},

	{geo.Vector{geo.Point{1, 4, 0}, tracks.DIR_180_DEG},
		buildTrack([]tracks.SegmentType{tracks.ELEM_FLAT, tracks.ELEM_LEFT_QUARTER_TURN_3_TILES, tracks.ELEM_LEFT_QUARTER_TURN_3_TILES})},

	{geo.Vector{geo.Point{3, 4, 0}, tracks.DIR_180_DEG},
		buildTrack([]tracks.SegmentType{
			tracks.ELEM_FLAT, tracks.ELEM_FLAT, tracks.ELEM_FLAT, tracks.ELEM_LEFT_QUARTER_TURN_3_TILES, tracks.ELEM_LEFT_QUARTER_TURN_3_TILES,
		})},
}

func Test2DTrack(t *testing.T) {
	stationStart := geo.Vector{geo.Point{0, 0, 0}, tracks.DIR_STRAIGHT}
	for _, tt := range trackTests {
		out := connect2DTrackPieces(tt.in, stationStart)
		helper := fmt.Sprintf("trackEnd: (%d, %d, %d) %d", round(tt.in.Point[0]), round(tt.in.Point[1]), round(tt.in.Point[2]), tt.in.Dir)
		if len(out) != len(tt.expected) {
			t.Errorf("%s expected track to be %v, was %v", helper, tt.expected, out)
			break
		}
		for i := range out {
			if tt.expected[i] != out[i] {
				t.Errorf("%s expected track to be %v, was %v", helper, tt.expected, out)
			}
		}
		fmt.Println("one test down")
	}
}

func TestRightTurn(t *testing.T) {
	trackEnd := geo.Vector{geo.Point{
		7, 3, 11,
	}, tracks.DIR_90_DEG_RIGHT}
	stationStart := geo.Vector{geo.Point{
		5, 1, 11,
	}, tracks.DIR_180_DEG}
	elems := rightTurn(trackEnd, stationStart)
	if len(elems) != 1 {
		t.Fatalf("expected a single right turn, got %#v", elems)
	}
	if elems[0].Segment != tracks.TS_MAP[tracks.ELEM_RIGHT_QUARTER_TURN_3_TILES] {
		t.Fatalf("expected a right turn, got %#v", elems[0])
	}
}

var straightenTests = []struct {
	in       tracks.Element
	inv      geo.Vector
	expected []tracks.Element
}{
	{buildOneTrack(tracks.ELEM_FLAT),
		geo.Vector{geo.Point{-2, 2, 0}, tracks.DIR_90_DEG_RIGHT},
		[]tracks.Element{}},

	{buildOneTrack(tracks.ELEM_60_DEG_UP),
		geo.Vector{geo.Point{-2, 2, 0}, tracks.DIR_90_DEG_RIGHT},
		buildTrack([]tracks.SegmentType{
			tracks.ELEM_60_DEG_UP_TO_25_DEG_UP,
			tracks.ELEM_25_DEG_UP_TO_FLAT,
		})},

	{buildOneTrack(tracks.ELEM_25_DEG_UP),
		geo.Vector{geo.Point{-2, 2, 0}, tracks.DIR_90_DEG_RIGHT},
		[]tracks.Element{buildOneTrack(tracks.ELEM_25_DEG_UP_TO_FLAT)}},

	{buildOneTrack(tracks.ELEM_25_DEG_DOWN),
		geo.Vector{geo.Point{-2, 2, 0}, tracks.DIR_90_DEG_LEFT},
		[]tracks.Element{buildOneTrack(tracks.ELEM_25_DEG_DOWN_TO_FLAT)}},

	{buildOneTrack(tracks.ELEM_60_DEG_DOWN),
		geo.Vector{geo.Point{-2, 2, 0}, tracks.DIR_90_DEG_LEFT},
		buildTrack([]tracks.SegmentType{
			tracks.ELEM_60_DEG_DOWN_TO_25_DEG_DOWN,
			tracks.ELEM_25_DEG_DOWN_TO_FLAT,
		})},

	{buildOneTrack(tracks.ELEM_BANKED_LEFT_QUARTER_TURN_5_TILES),
		geo.Vector{geo.Point{-2, 2, 0}, tracks.DIR_90_DEG_LEFT},
		[]tracks.Element{buildOneTrack(tracks.ELEM_LEFT_BANK_TO_FLAT)}},

	{buildOneTrack(tracks.ELEM_BANKED_RIGHT_QUARTER_TURN_5_TILES),
		geo.Vector{geo.Point{-2, 2, 0}, tracks.DIR_90_DEG_RIGHT},
		[]tracks.Element{buildOneTrack(tracks.ELEM_RIGHT_BANK_TO_FLAT)}},
}

func TestStraighten(t *testing.T) {
	for _, tt := range straightenTests {
		pieces, _ := straightenTrack(tt.in, tt.inv)
		if len(pieces) != len(tt.expected) {
			t.Errorf("%v expected track to be %v, was %v", tt.in, tt.expected, pieces)
			break
		}
		for i := range pieces {
			if tt.expected[i] != pieces[i] {
				t.Errorf("%v expected track to be %v, was %v", tt.in, tt.expected, pieces)
			}
		}
	}
}
