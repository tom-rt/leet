package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var ret int

	ret = reverse(123)
	if ret != 321 {
		t.Fatalf("123: expected 321, got %d", ret)
	}

	ret = reverse(-123)
	if ret != -321 {
		t.Fatalf("-123: expected -321, got %d", ret)
	}

	ret = reverse(1534236469)
	if ret != 0 {
		t.Fatalf("1534236469: expected 0, got %d", ret)
	}

	ret = reverse(-9999999999)
	if ret != 0 {
		t.Fatalf("-9999999999: expected 0, got %d", ret)
	}

}
