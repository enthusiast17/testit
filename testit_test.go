package testit

import (
	"fmt"
	"testing"
)

func TestItTest(t *testing.T) {
	a, b := 1, 1
	if err := equal(a, b); err != nil {
		t.Errorf(failMessage, nil, err)
	}
	if err := notEqual(a, b); err == nil {
		t.Errorf(failMessage, fmt.Errorf(failMessageNotEqual, a, b, a, b), err)
	}

	as, bs := "hello", "there"
	if err := notEqual(as, bs); err != nil {
		t.Errorf(failMessage, nil, err)
	}
	if err := equal(as, bs); err == nil {
		t.Errorf(failMessage, fmt.Errorf(failMessageEqual, as, bs, as, bs), err)
	}
}
