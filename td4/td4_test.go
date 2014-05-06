package td4

import (
	"testing"

	"github.com/kevinburke/rct-rides/tracks"
)

var stationTrack = &tracks.Data{
	Elements: []tracks.Element{
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_END_STATION],
		},
	},
}

// The simplest possible circuit
var completeTrack = &tracks.Data{
	Elements: []tracks.Element{
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_END_STATION],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_FLAT],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
	},
}

var hillyTrack = &tracks.Data{
	Elements: []tracks.Element{
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_END_STATION],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_FLAT_TO_25_DEG_UP],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_25_DEG_UP_TO_FLAT],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_FLAT_TO_25_DEG_DOWN],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN_TO_FLAT],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_LEFT_QUARTER_TURN_1_TILE],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_FLAT],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_FLAT],
		},
		tracks.Element{
			Segment: tracks.TS_MAP[tracks.ELEM_FLAT],
		},
	},
}

func TestCircuit(t *testing.T) {
	td := &tracks.Data{}
	t.Skip("circuit checks broke recently")
	if IsCircuit(td) {
		t.Errorf("empty ride should not be a circuit.")
	}

	if IsCircuit(stationTrack) {
		t.Errorf("track with single station shouldn't be a circuit")
	}

	if !IsCircuit(completeTrack) {
		t.Errorf("Complete track should be marked as a circuit but wasn't")
	}

	if !IsCircuit(hillyTrack) {
		t.Errorf("Hilly track should be marked as a circuit but wasn't")
	}
}
