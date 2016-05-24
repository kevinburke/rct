package genetic

import (
	"fmt"
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

func elem(s tracks.SegmentType) tracks.Element {
	return tracks.Element{
		Segment: tracks.TS_MAP[s],
	}
}

func TestCrossover(t *testing.T) {
	parent1 := &Member{
		Track: []tracks.Element{
			elem(tracks.ELEM_BEGIN_STATION),
			elem(tracks.ELEM_MIDDLE_STATION),
			elem(tracks.ELEM_MIDDLE_STATION),
			elem(tracks.ELEM_MIDDLE_STATION),
			elem(tracks.ELEM_END_STATION),
			elem(tracks.ELEM_FLAT_TO_RIGHT_BANK),
			elem(tracks.ELEM_RIGHT_QUARTER_TURN_3_TILES_BANK),
			elem(tracks.ELEM_RIGHT_QUARTER_TURN_3_TILES_BANK),
			elem(tracks.ELEM_RIGHT_QUARTER_BANKED_HELIX_LARGE_DOWN),
			elem(tracks.ELEM_RIGHT_BANK_TO_25_DEG_DOWN),
			elem(tracks.ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_DOWN),
			elem(tracks.ELEM_25_DEG_DOWN_TO_LEFT_BANK),
			elem(tracks.ELEM_LEFT_BANK_TO_25_DEG_UP),
			elem(tracks.ELEM_25_DEG_UP_TO_FLAT),
			elem(tracks.ELEM_FLAT_TO_25_DEG_UP),
			elem(tracks.ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_UP),
		},
	}
	parent2 := &Member{
		Track: []tracks.Element{
			elem(tracks.ELEM_BEGIN_STATION),
			elem(tracks.ELEM_MIDDLE_STATION),
			elem(tracks.ELEM_MIDDLE_STATION),
			elem(tracks.ELEM_MIDDLE_STATION),
			elem(tracks.ELEM_END_STATION),
			elem(tracks.ELEM_S_BEND_RIGHT),
			elem(tracks.ELEM_LEFT_QUARTER_TURN_5_TILES),
			elem(tracks.ELEM_LEFT_QUARTER_TURN_3_TILES),
			elem(tracks.ELEM_FLAT_TO_RIGHT_BANK),
			elem(tracks.ELEM_RIGHT_BANK_TO_25_DEG_UP),
			elem(tracks.ELEM_25_DEG_UP_TO_LEFT_BANK),
			elem(tracks.ELEM_LEFT_BANK),
			elem(tracks.ELEM_LEFT_HALF_BANKED_HELIX_UP_SMALL),
			elem(tracks.ELEM_LEFT_HALF_BANKED_HELIX_DOWN_SMALL),
			elem(tracks.ELEM_LEFT_HALF_BANKED_HELIX_DOWN_LARGE),
			elem(tracks.ELEM_LEFT_HALF_BANKED_HELIX_UP_SMALL),
		},
	}
	child1, child2 := crossoverAtPoint(parent1, parent2, 11)
	checkCompatible := func(track []tracks.Element) {
		for i, elem := range track[:len(track)-1] {
			fmt.Printf("checking compatible: %s => %s\n", elem.Segment.String(), track[i+1].Segment.String())
			if tracks.Compatible(elem, track[i+1]) == false {
				t.Errorf("Crossover generated incompatible segments: %s, %s", elem.Segment.String(), track[i+1].Segment.String())
			}
		}
	}
	checkCompatible(child1.Track)
	checkCompatible(child2.Track)
	t.Fail()
}
