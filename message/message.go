package message

const (
	// Fail constant message when got fail.
	Fail = "FAIL: Expected [%[1]v], got [%[2]v] | called from [file: %[3]v, line: %[3]v]"

	// FailEqual constant message when got fail from equal.
	FailEqual = "FAIL: Expected [%[1]v == %[2]v], got [%[1]v != %[2]v] | called from [file: %[3]v, line: %[4]v]"

	// FailNotEqual constant message when got fail from not equal.
	FailNotEqual = "FAIL: Expected [%[1]v != %[2]v], got [%[1]v == %[2]v] | called from [file: %[3]v, line: %[4]v]"

	// PassEqual constant message when got pass from equal.
	PassEqual = "PASS: [%[1]v] == [%[2]v] | called from [file: %[3]v, line: %[4]v]"

	// PassNotEqual constant message when got pass from not equal.
	PassNotEqual = "PASS: [%[1]v] != [%[2]v] | called from [file: %[3]v, line: %[4]v]"
)
