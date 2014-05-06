package tracks

import (
	"fmt"
	"testing"
)

func TestPossibilities(t *testing.T) {
	steep := TS_MAP[ELEM_60_DEG_UP]
	fmt.Println(steep.Possibilities())
	t.Fatalf("foo")
}
