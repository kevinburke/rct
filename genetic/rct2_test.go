package genetic

import (
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
