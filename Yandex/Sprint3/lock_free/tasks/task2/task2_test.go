package main

import (
	"testing"
)

func TestAtomicSwap(t *testing.T) {
	var value1 int32 = 42
	var value2 int32 = 24

	AtomicSwap(&value1, &value2)

	if value1 != 24 || value2 != 42 {
		t.Errorf("Expected values after swap: 24, 42; got: %v, %v", value1, value2)
	}
}
