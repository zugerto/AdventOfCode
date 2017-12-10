package main

import (
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {

	Result, _, _ := process([]int{0, 1, 2, 3, 4}, []int{3, 4, 1, 5}, 0, 0)
	Want := []int{3, 4, 2, 1, 0}
	if !reflect.DeepEqual(Result, Want) {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestReverse1(t *testing.T) {

	Result := reverse([]int{0, 1, 2, 3, 4}, 0, 3)
	Want := []int{2, 1, 0, 3, 4}
	if !reflect.DeepEqual(Result, Want) {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestReverse2(t *testing.T) {

	Result := reverse([]int{2, 1, 0, 3, 4}, 3, 4)
	Want := []int{4, 3, 0, 1, 2}
	if !reflect.DeepEqual(Result, Want) {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestSalt(t *testing.T) {

	Result := salt("1,2,3")
	Want := []int{49, 44, 50, 44, 51, 17, 31, 73, 47, 23}
	if !reflect.DeepEqual(Result, Want) {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestAll1(t *testing.T) {

	Result := hexadecimal(dense(hash(createList(256), salt(""), 64)))
	Want := "a2582a3a0e66e6e86e3812dcb672a272"
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestAll2(t *testing.T) {

	Result := hexadecimal(dense(hash(createList(256), salt("AoC 2017"), 64)))
	Want := "33efeb34ea91902bb2f59c9920caa6cd"
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestAll3(t *testing.T) {

	Result := hexadecimal(dense(hash(createList(256), salt("1,2,3"), 64)))
	Want := "3efbe78a8d82f29979031a4aa0b16a9d"
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestAll4(t *testing.T) {

	Result := hexadecimal(dense(hash(createList(256), salt("1,2,4"), 64)))
	Want := "63960835bcdc130f0b66d7ff4f6a5a8e"
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}
