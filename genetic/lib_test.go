package genetic

import (
	"testing"

	"github.com/kevinburke/rct/tracks"
)

func TestSwap(t *testing.T) {
	a := &Member{
		Track: []tracks.Element{
			{Segment: tracks.TS_MAP[tracks.ELEM_FLAT]},
			{Segment: tracks.TS_MAP[tracks.ELEM_25_DEG_UP]},
		},
	}
	b := &Member{
		Track: []tracks.Element{
			{Segment: tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN]},
			{Segment: tracks.TS_MAP[tracks.ELEM_60_DEG_DOWN]},
		},
	}
	c, d := Swap(a, b, 1, 1)
	cexpected := []tracks.Element{
		{Segment: tracks.TS_MAP[tracks.ELEM_FLAT]},
		{Segment: tracks.TS_MAP[tracks.ELEM_60_DEG_DOWN]},
	}
	if c.Track[0] != cexpected[0] || c.Track[1] != cexpected[1] {
		t.Errorf("expected c to be %v, was %v", cexpected, c.Track)
	}

	dexpected := []tracks.Element{
		{Segment: tracks.TS_MAP[tracks.ELEM_25_DEG_DOWN]},
		{Segment: tracks.TS_MAP[tracks.ELEM_25_DEG_UP]},
	}
	if d.Track[0] != dexpected[0] || d.Track[1] != dexpected[1] {
		t.Errorf("expected d to be %v, was %v", dexpected, d.Track)
	}
}
