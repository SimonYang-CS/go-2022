package main

import "testing"

func Test_Add1(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}

func Test_Add2(t *testing.T) {
	result := add(2, -1)
	if result != 1 {
		t.Error("incorrect result: expected 1, got", result)
	}
}
