package exe_reader

import "testing"

var sidewaysTests = []struct {
	in       int
	expected int
}{
	{0, 0},
	{64, 3},
	{96, 4},
	{160, -4},
	{192, -3},
	{224, -2},
	{32, 2},
}

func TestSidewaysDelta(t *testing.T) {
	for _, tt := range sidewaysTests {
		out := SidewaysDelta(tt.in)
		if out != tt.expected {
			t.Errorf("expected SidewaysDelta(%d) to be %d, was %d", tt.in, tt.expected, out)
		}
	}
}
