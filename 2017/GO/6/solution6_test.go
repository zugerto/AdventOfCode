package main

import (
	"testing"
)

func TestReallocationCount1(t *testing.T) {

	Result, _ := ReallocationCount([]int{0, 2, 7, 0})
	Want := 5
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestReallocationCount2(t *testing.T) {

	_, Result := ReallocationCount([]int{0, 2, 7, 0})
	Want := 4
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}
