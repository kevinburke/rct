package tracks

import (
	"fmt"
	"testing"
)

// Somehow got ON_RIDE_PHOTO after FLAT_TO_25_DEG_DOWN, trying to figure out
// how
func TestContiguousTracks(t *testing.T) {
	elem := Element{
		Segment: TS_MAP[ELEM_FLAT_TO_25_DEG_DOWN],
	}
	pos := elem.Possibilities()
	fmt.Println(pos)
}
