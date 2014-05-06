package main

import (
	"testing"
)

func testMarshalControlFlags(t *testing.T) {
	t.Parallel()
	cFlags := ControlFlags{
		Load:           3,
		UseMaximumTime: true,
	}
	n := marshalControlFlags(cFlags)
	if !hasBit(n, 7) {
		t.Errorf("Maximum time bit should have been set, but wasn't, n is %d", n)
	}
}