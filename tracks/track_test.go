package tracks

import "testing"

func TestPossibilities(t *testing.T) {
	steep := TS_MAP[ELEM_FLAT]
	elem := Element{Segment: steep}
	poss := elem.Possibilities()
	for _, p := range poss {
		if p.Segment.Type == 0xff {
			t.Errorf("should not have contained the end of ride but did")
		}
		if p.Segment.Type == ELEM_END_STATION {
			t.Errorf("should not be possible to build a station")
		}
	}
}
