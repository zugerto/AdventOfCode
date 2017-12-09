package main

import (
	"testing"
)

func TestParseInputRow1(t *testing.T) {

	Result, _ := ParseInputRow("ktlj (57)")
	Want := "ktlj"
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}
