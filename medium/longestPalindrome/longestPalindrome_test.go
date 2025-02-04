package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {

	longestPal := longestPalindrome("bsssssbb")
	if longestPal != "bsssssb" {
		t.Fatalf("expected bsssssb got %s", longestPal)
	}
}
