package main

import (
	"strings"
	"testing"
)

func TestCalculateScore1(t *testing.T) {

	Result, _ := Calculate([]string{"{", "}"})
	Want := 1
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateScore2(t *testing.T) {

	Result, _ := Calculate([]string{"{", "{", "{", "}", "}", "}"})
	Want := 6
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateScore3(t *testing.T) {

	Result, _ := Calculate([]string{"{", "{", "}", ",", "{", "}", "}"})
	Want := 5
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateScore4(t *testing.T) {

	Result, _ := Calculate([]string{"{", "{", "{", "}", ",", "{", "}", ",", "{", "{", "}", "}", "}", "}"})
	Want := 16
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateScore5(t *testing.T) {

	Result, _ := Calculate(strings.Split("{<a>,<a>,<a>,<a>}", ""))
	Want := 1
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateScore6(t *testing.T) {

	Result, _ := Calculate(strings.Split("{{<ab>},{<ab>},{<ab>},{<ab>}}", ""))
	Want := 9
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateScore7(t *testing.T) {

	Result, _ := Calculate(strings.Split("{{<!!>},{<!!>},{<!!>},{<!!>}}", ""))
	Want := 9
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateScore8(t *testing.T) {

	Result, _ := Calculate(strings.Split("{{<a!>},{<a!>},{<a!>},{<ab>}}", ""))
	Want := 3
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateGarbage1(t *testing.T) {

	_, Result := Calculate(strings.Split("<>", ""))
	Want := 0
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateGarbage2(t *testing.T) {

	_, Result := Calculate(strings.Split("<random characters>", ""))
	Want := 17
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateGarbage3(t *testing.T) {

	_, Result := Calculate(strings.Split("<<<<>", ""))
	Want := 3
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateGarbage4(t *testing.T) {

	_, Result := Calculate(strings.Split("<{!>}>", ""))
	Want := 2
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateGarbage5(t *testing.T) {

	_, Result := Calculate(strings.Split("<!!>", ""))
	Want := 0
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateGarbage6(t *testing.T) {

	_, Result := Calculate(strings.Split("<!!!>>", ""))
	Want := 0
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestCalculateGarbage7(t *testing.T) {

	_, Result := Calculate(strings.Split("<{o\"i!a,<{i<a>", ""))
	Want := 10
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}
