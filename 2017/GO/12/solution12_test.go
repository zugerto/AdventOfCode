package main

import (
	"testing"
)

func TestPartA(t *testing.T) {

	programs := make(map[int][]int)

	Connect(programs, 0, []int{2})
	Connect(programs, 1, []int{1})
	Connect(programs, 2, []int{0, 3, 4})
	Connect(programs, 3, []int{2, 4})
	Connect(programs, 4, []int{2, 3, 6})
	Connect(programs, 5, []int{6})
	Connect(programs, 6, []int{4, 5})

	Result := CountA(programs)
	Want := 6
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestPartB(t *testing.T) {

	programs := make(map[int][]int)

	Connect(programs, 0, []int{2})
	Connect(programs, 1, []int{1})
	Connect(programs, 2, []int{0, 3, 4})
	Connect(programs, 3, []int{2, 4})
	Connect(programs, 4, []int{2, 3, 6})
	Connect(programs, 5, []int{6})
	Connect(programs, 6, []int{4, 5})

	Result := CountB(programs)
	Want := 2
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}
