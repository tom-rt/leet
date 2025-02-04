package main

import (
	"testing"
)

func TestConvertZigzag(t *testing.T) {
	// zigzag
	zigzag := convertZigzag("BONJOURATOUS", 2)
	if zigzag != "BNORTUOJUAOS" {
		t.Fatalf("expected BNORTUOJUAOS, got %s", zigzag)
	}
}
