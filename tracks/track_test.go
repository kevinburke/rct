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
	}
}
