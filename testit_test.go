package testit

import "testing"

func TestItTest(t *testing.T) {
	if a, b := 1, 1; equal(a, b) != nil {
		t.Errorf("Expected [%v], got [%v]", a, b)
	}

	if a, b := 1, 1; notEqual(a, b) == nil {
		t.Errorf("Expected [%v], got [%v]", a, b)
	}

	if a, b := "hello", "there"; notEqual(a, b) != nil {
		t.Errorf("Expected [%v], got [%v]", a, b)
	}

	if a, b := "hello", "there"; equal(a, b) == nil {
		t.Errorf("Expected [%v], got [%v]", a, b)
	}
}
