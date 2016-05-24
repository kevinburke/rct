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

func TestCompatible(t *testing.T) {
	testCases := []struct {
		first      Element
		second     Element
		compatible bool
	}{
		{elem(ELEM_FLAT_TO_25_DEG_UP), elem(ELEM_LEFT_QUARTER_TURN_5_TILES), false},
		{elem(ELEM_LEFT_QUARTER_TURN_3_TILES_25_DEG_DOWN), elem(ELEM_LEFT_BANK), false},
	}
	for _, tt := range testCases {
		compat := Compatible(tt.first, tt.second)
		if compat != tt.compatible {
			t.Errorf("Compatible: first: %s, next: %s, want %t, got %t", tt.first.Segment.String(), tt.second.Segment.String(), tt.compatible, compat)
		}
	}
}
