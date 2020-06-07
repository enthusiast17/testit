package testit

import (
	"fmt"
	"testing"
)

var (
	failMessage         = "FAIL: Expected [%v], got [%v]"
	failMessageEqual    = "FAIL: Expected [%v == %v], got [%v != %v]"
	failMessageNotEqual = "FAIL: Expected [%v != %v], got [%v == %v]"
	passMessageEqual    = "PASS: [%v] is equal to [%v]"
	passMessageNotEqual = "PASS: [%v] is not equal to [%v]"
)

// Equal prints FAIL/PASS (1 interface is expected value, 2 interface is value for compare)
func Equal(t *testing.T, i1, i2 interface{}) {
	if err := equal(i1, i2); err != nil {
		t.Error(err)
	}
	t.Logf(passMessageEqual, i1, i2)
}

// NotEqual prints FAIL/PASS (1 interface is expected value, 2 interface is value for compare)
func NotEqual(t *testing.T, i1, i2 interface{}) {
	if err := notEqual(i1, i2); err != nil {
		t.Error(err)
	}
	t.Logf(passMessageNotEqual, i1, i2)
}

func equal(i1, i2 interface{}) error {
	if i1 != i2 {
		return fmt.Errorf(failMessageEqual, i1, i2, i1, i2)
	}
	return nil
}

func notEqual(i1, i2 interface{}) error {
	if i1 == i2 {
		return fmt.Errorf(failMessageNotEqual, i1, i2, i1, i2)
	}
	return nil
}
