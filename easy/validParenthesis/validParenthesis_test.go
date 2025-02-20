package main

import "testing"

func TestIsValid(t *testing.T) {
	ret := isValid("((()))")
	if !ret {
		t.Fatalf("([)]: expected true, got false")
	}

	ret = isValid("([)]")
	if ret {
		t.Fatalf("([)]: expected false, got true")
	}

	ret = isValid("[([]])")
	if ret {
		t.Fatalf("[([]]): expected false, got true")
	}
}
