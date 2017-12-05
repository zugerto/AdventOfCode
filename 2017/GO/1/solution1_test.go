package main

import (
	"testing"
)

func TestSolutionA1122(t *testing.T) {

	Result := SolutionA([]string{"1", "1", "2", "2"})
	Want := 3
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestSolutionA1111(t *testing.T) {

	Result := SolutionA([]string{"1", "1", "1", "1"})
	Want := 4
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestSolutionA1234(t *testing.T) {

	Result := SolutionA([]string{"1", "2", "3", "4"})
	Want := 0
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}

}

func TestSolutionA91212129(t *testing.T) {

	Result := SolutionA([]string{"9", "1", "2", "1", "2", "1", "2", "9"})
	Want := 9
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}

}

func TestSolutionB1212(t *testing.T) {

	Result := SolutionB([]string{"1", "2", "1", "2"})
	Want := 6
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}

}

func TestSolutionB1221(t *testing.T) {

	Result := SolutionB([]string{"1", "2", "2", "1"})
	Want := 0
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}

}

func TestSolutionB123425(t *testing.T) {

	Result := SolutionB([]string{"1", "2", "3", "4", "2", "5"})
	Want := 4
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}

}

func TestSolutionB123123(t *testing.T) {

	Result := SolutionB([]string{"1", "2", "3", "1", "2", "3"})
	Want := 12
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}

}

func TestSolutionB12131415(t *testing.T) {

	Result := SolutionB([]string{"1", "2", "1", "3", "1", "4", "1", "5"})
	Want := 4
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}

}
