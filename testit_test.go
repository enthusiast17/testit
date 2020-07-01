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
	}
	if err := notEqual(a, b); err == nil {
		t.Errorf(message.Fail, fmt.Errorf(message.FailNotEqual, a, b, a, b), err)
	}

	as, bs := "hello", "there"
	if err := notEqual(as, bs); err != nil {
		t.Errorf(message.Fail, nil, err)
	}
	if err := equal(as, bs); err == nil {
		t.Errorf(message.Fail, fmt.Errorf(message.FailEqual, as, bs, as, bs), err)
	}
}
