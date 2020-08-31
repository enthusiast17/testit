package testit

import (
	"fmt"
	"runtime"
	"testing"

	message "./message"
)

// Equal prints FAIL/PASS (1 interface is expected value, 2 interface is value for compare).
func Equal(t *testing.T, i1, i2 interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file, line = "nil", -1
	}
	if err := equal(i1, i2, file, line); err != nil {
		t.Error(err)
	}
	t.Logf(message.PassEqual, i1, i2, file, line)
}

// NotEqual prints FAIL/PASS (1 interface is expected value, 2 interface is value for compare).
func NotEqual(t *testing.T, i1, i2 interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file, line = "nil", -1
	}
	if err := notEqual(i1, i2, file, line); err != nil {
		t.Error(err)
	}
	t.Logf(message.PassNotEqual, i1, i2, file, line)
}

func equal(i1, i2, file, line interface{}) error {
	if i1 != i2 {
		return fmt.Errorf(message.FailEqual, i1, i2, i1, i2, file, line)
	}
	return nil
}

func notEqual(i1, i2, file, line interface{}) error {
	if i1 == i2 {
		return fmt.Errorf(message.FailNotEqual, i1, i2, i1, i2, file, line)
	}
	return nil
}
