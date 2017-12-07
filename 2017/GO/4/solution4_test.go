package main

import (
	"testing"
)

func TestValidPassprase1(t *testing.T) {

	Result := ValidPassphrase([]string{"aa", "bb", "cc", "dd", "ee"})
	Want := 1
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestValidPassprase2(t *testing.T) {

	Result := ValidPassphrase([]string{"aa", "bb", "cc", "dd", "aa"})
	Want := 0
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestValidPassprase3(t *testing.T) {

	Result := ValidPassphrase([]string{"aa", "bb", "cc", "dd", "aaa"})
	Want := 1
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestValidPasspraseAlsoConsiderAnagrams1(t *testing.T) {

	Result := ValidPassphraseAlsoConsiderAnagrams([]string{"abcde", "fghij"})
	Want := 1
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestValidPasspraseAlsoConsiderAnagrams2(t *testing.T) {

	Result := ValidPassphraseAlsoConsiderAnagrams([]string{"abcde", "xyz", "ecdab"})
	Want := 0
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestValidPasspraseAlsoConsiderAnagrams3(t *testing.T) {

	Result := ValidPassphraseAlsoConsiderAnagrams([]string{"a", "ab", "abc", "abd", "abf", "abj"})
	Want := 1
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestValidPasspraseAlsoConsiderAnagrams4(t *testing.T) {

	Result := ValidPassphraseAlsoConsiderAnagrams([]string{"iiii", "oiii", "ooii", "oooi", "oooo"})
	Want := 1
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}

func TestValidPasspraseAlsoConsiderAnagrams5(t *testing.T) {

	Result := ValidPassphraseAlsoConsiderAnagrams([]string{"oiii", "ioii", "iioi", "iiio"})
	Want := 0
	if Result != Want {
		t.Errorf("Result: %d, want: %d", Result, Want)
	}
}
