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

func TestRightTurn(t *testing.T) {
	trackEnd := geo.Vector{geo.Point{
		7, 3, 11,
	}, tracks.DIR_90_DEG_RIGHT}
	stationStart := geo.Vector{geo.Point{
		5, 1, 11,
	}, tracks.DIR_180_DEG}
	fmt.Println("at (7, 3) facing up. turning right. should get to 9, 5)")
	elems := rightTurn(trackEnd, stationStart)
	if len(elems) != 1 {
		t.Fatalf("expected a single right turn, got %#v", elems)
	}
	if elems[0].Segment != tracks.TS_MAP[tracks.ELEM_RIGHT_QUARTER_TURN_3_TILES] {
		t.Fatalf("expected a right turn, got %#v", elems[0])
	}
}
