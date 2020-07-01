package message

// Fail constant message when got fail.
const Fail = "FAIL: Expected [%v], got [%v]"

// FailEqual constant message when got fail from equal.
const FailEqual = "FAIL: Expected [%v == %v], got [%v != %v]"

// FailNotEqual constant message when got fail from not equal.
const FailNotEqual = "FAIL: Expected [%v != %v], got [%v == %v]"

// PassEqual constant message when got pass from equal.
const PassEqual = "PASS: [%v] == [%v]"

// PassNotEqual constant message when got pass from not equal.
const PassNotEqual = "PASS: [%v] != [%v]"
