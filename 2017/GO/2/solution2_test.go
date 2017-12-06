package main

import (
	"testing"
)

func TestChecksumA(t *testing.T) {

	Result := Checksum([]string{"5", "1", "9", "5"})
	Want := 8
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestChecksumB(t *testing.T) {

	Result := Checksum([]string{"7", "5", "3"})
	Want := 4
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestChecksumC(t *testing.T) {

	Result := Checksum([]string{"2", "4", "6", "8"})
	Want := 6
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestEvenlyDivideA(t *testing.T) {

	Result := EvenlyDivide([]string{"5", "9", "2", "8"})
	Want := 4
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestEvenlyDivideB(t *testing.T) {

	Result := EvenlyDivide([]string{"9", "4", "7", "3"})
	Want := 3
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestEvenlyDivideC(t *testing.T) {

	Result := EvenlyDivide([]string{"3", "8", "6", "5"})
	Want := 2
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}
