package main

import (
	"testing"
)

func TestCalculate1(t *testing.T) {

	Result := CalculateNrOfStepsToChild([]string{"ne", "ne", "ne"})
	Want := 3
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculate2(t *testing.T) {

	Result := CalculateNrOfStepsToChild([]string{"ne", "ne", "sw", "sw"})
	Want := 0
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculate3(t *testing.T) {

	Result := CalculateNrOfStepsToChild([]string{"ne", "ne", "s", "s"})
	Want := 2
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}
