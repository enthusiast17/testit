package testit

import (
	"fmt"
	"testing"
)

// Equal prints FAIL/PASS (1 interface is expected value, 2 interface is value for compare)
func Equal(t *testing.T, i1, i2 interface{}) {
	if err := equal(i1, i2); err != nil {
		t.Error(err)
	}
	t.Logf("PASS: [%v] is equal to [%v]", i1, i2)
}

// NotEqual prints FAIL/PASS (1 interface is expected value, 2 interface is value for compare)
func NotEqual(t *testing.T, i1, i2 interface{}) {
	if err := notEqual(i1, i2); err != nil {
		t.Error(err)
	}
	t.Logf("PASS: [%v] is not equal to [%v]", i1, i2)
}

func equal(i1, i2 interface{}) error {
	if i1 != i2 {
		return fmt.Errorf("FAIL: Expected [%v], got [%v]", i1, i2)
	}
	return nil
}

func notEqual(i1, i2 interface{}) error {
	if i1 == i2 {
		return fmt.Errorf("FAIL: Expected [%v], got [%v]", i1, i2)
	}
	return nil
}
