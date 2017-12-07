package main

import (
	"testing"
)

func TestSpiralStepsA1(t *testing.T) {

	Result := SpiralSteps(1)
	Want := 0
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestSpiralStepsA12(t *testing.T) {

	Result := SpiralSteps(12)
	Want := 3
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestSpiralStepsA23(t *testing.T) {

	Result := SpiralSteps(23)
	Want := 2
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestSpiralStepsA1024(t *testing.T) {

	Result := SpiralSteps(1024)
	Want := 31
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestNeighbourSum1(t *testing.T) {

	Result := NeighbourSum(1)
	Want := 1
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestNeighbourSum2(t *testing.T) {

	Result := NeighbourSum(1)
	Want := 1
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestNeighbourSum3(t *testing.T) {

	Result := NeighbourSum(3)
	Want := 2
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestNeighbourSum4(t *testing.T) {

	Result := NeighbourSum(4)
	Want := 4
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestNeighbourSum5(t *testing.T) {

	Result := NeighbourSum(5)
	Want := 5
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestNeighbourSum6(t *testing.T) {

	Result := NeighbourSum(6)
	Want := 10
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}
