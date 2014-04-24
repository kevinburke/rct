package td4

import (
	"testing"
)

func TestSin(t *testing.T) {
	if sindeg(630) != -1 {
		t.Errorf("sin(630) should equal -1 but was %d", sindeg(630))
	}
	if sindeg(0) != 0 {
		t.Errorf("sin(0) should equal 0 but was %d", sindeg(0))
	}
	if sindeg(180) != 0 {
		t.Errorf("sin(180) should equal 0 but was %d", sindeg(180))
	}
}
