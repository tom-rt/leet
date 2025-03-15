package main

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	ret := romanToInt("IV")
	if ret != 4 {
		fmt.Printf("IV: expected 4, got %d\n", ret)
		t.Fatal()
	}
}
