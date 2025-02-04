package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	ret := isPalindrome(1234321)
	if !ret {
		t.Fatalf("expected true, got false")
	}
}
