package testit

import (
	"fmt"
	"runtime"
	"testing"

	message "./message"
)

func TestItTest(t *testing.T) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		file = "nil"
	}
	a, b := 1, 1
	if err := equal(a, b, file, 18); err != nil {
		t.Errorf(message.Fail, nil, err, file, 18)
	} else {
		t.Logf(message.PassEqual, err, nil, file, 18)
	}

	if err := notEqual(a, b, file, 24); err == nil {
		t.Errorf(message.Fail, fmt.Errorf(message.FailNotEqual, a, b, a, b, file, 24), err, file, 24)
	} else {
		t.Logf(message.PassEqual, err, fmt.Errorf(message.FailNotEqual, a, b, a, b, file, 24), file, 24)
	}

	as, bs := "hello", "there"
	if err := notEqual(as, bs, file, 31); err != nil {
		t.Errorf(message.Fail, nil, err, file, 31)
	} else {
		t.Logf(message.PassEqual, err, nil, file, 31)
	}

	if err := equal(as, bs, file, 37); err == nil {
		t.Errorf(message.Fail, fmt.Errorf(message.FailEqual, as, bs, as, bs, file, 37), err, file, 37)
	} else {
		t.Logf(message.PassEqual, err, fmt.Errorf(message.FailEqual, as, bs, as, bs, file, 37), file, 37)
	}
}
