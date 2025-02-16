package main

import "testing"

func TestMyAtoi(t *testing.T) {
	ret := myAtoi("00123")
	if ret != 123 {
		t.Fatalf("00123: expected 123, got: %d", ret)
	}
	ret = myAtoi("aaaa123aaaa")
	if ret != 0 {
		t.Fatalf("aaaa123aaaa: expected 0, got: %d", ret)
	}
	ret = myAtoi("-00123")
	if ret != -123 {
		t.Fatalf("-00123: expected -123, got: %d", ret)
	}
	ret = myAtoi("010123")
	if ret != 10123 {
		t.Fatalf("010123: expected 10234, got: %d", ret)
	}
	ret = myAtoi("   -042")
	if ret != -42 {
		t.Fatalf("   -042: expected -42, got: %d", ret)
	}
	ret = myAtoi("   -04feeee")
	if ret != -4 {
		t.Fatalf("   -04feeee: expected -4, got: %d", ret)
	}
	ret = myAtoi("   -04f")
	if ret != -4 {
		t.Fatalf("   -04f: expected -4, got: %d", ret)
	}
	ret = myAtoi("0-1")
	if ret != 0 {
		t.Fatalf("0-1: expected 0, got: %d", ret)
	}
	ret = myAtoi("-91283472332")
	if ret != -2147483648 {
		t.Fatalf("-91283472332, expected -2147483648, got: %d", ret)
	}
}
