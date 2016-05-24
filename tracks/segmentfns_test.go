package tracks

import "testing"

func elem(s SegmentType) Element {
	return Element{
		Segment: TS_MAP[s],
	}
}

func TestNotPossible(t *testing.T) {
	testCases := []struct {
		input          Element
		notPossibility Element
	}{
		{elem(ELEM_FLAT_TO_25_DEG_UP), elem(ELEM_LEFT_QUARTER_TURN_5_TILES)},
	}
	for _, tt := range testCases {
		possibilities := tt.input.Possibilities()
		for _, poss := range possibilities {
			if poss.Segment.Type == tt.notPossibility.Segment.Type {
				t.Errorf("expected %s to not be a possibility for %s, but was", poss.Segment.String(), tt.notPossibility.Segment.String())
			}
		}
	}
}
