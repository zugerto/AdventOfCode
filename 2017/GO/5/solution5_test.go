package main

import (
	"testing"
)

func TestEscapeMaze(t *testing.T) {

	Result := EscapeMaze([]int{0, 3, 0, 1, -3})
	Want := 5
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestEscapeMazeStrangeOffset(t *testing.T) {

	Result := EscapeMazeStrangeOffset([]int{0, 3, 0, 1, -3})
	Want := 10
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}
