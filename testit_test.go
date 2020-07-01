package testit

import (
	"fmt"
	"testing"

	message "./message"
)

func TestItTest(t *testing.T) {
	a, b := 1, 1
	if err := equal(a, b); err != nil {
		t.Errorf(message.Fail, nil, err)
	} else {
		t.Logf(message.PassEqual, err, nil)
	}

	if err := notEqual(a, b); err == nil {
		t.Errorf(message.Fail, fmt.Errorf(message.FailNotEqual, a, b, a, b), err)
	} else {
		t.Logf(message.PassEqual, err, fmt.Errorf(message.FailNotEqual, a, b, a, b))
	}

	as, bs := "hello", "there"
	if err := notEqual(as, bs); err != nil {
		t.Errorf(message.Fail, nil, err)
	} else {
		t.Logf(message.PassEqual, err, nil)
	}

	if err := equal(as, bs); err == nil {
		t.Errorf(message.Fail, fmt.Errorf(message.FailEqual, as, bs, as, bs), err)
	} else {
		t.Logf(message.PassEqual, err, fmt.Errorf(message.FailEqual, as, bs, as, bs))
	}
}
