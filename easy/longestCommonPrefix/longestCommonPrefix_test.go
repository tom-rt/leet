package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	input := make([]string, 3)
	input[0] = "toto"
	input[1] = "toto"
	input[2] = "tot"
	ret := longestCommonPrefix(input)
	if ret != "tot" {
		t.Fatalf("expected tot, got %s", ret)
	}
}
