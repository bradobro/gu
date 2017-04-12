package should

import "fmt"

// Ok is what an Assertion returns when its condition is true.
const Ok = ""

/*Assertion provides a generic assertion signature.

actual:   value under test (if nil, failMessage CAN document the assertion)
expected: optional matcher params

This concept defived from https://github.com/smartystreets/assertions
*/
type Assertion func(actual interface{}, expected ...interface{}) (failMessage string)

// Not negates the assertion it wraps.
// BUG: it doesn't properly rename the Function
func Not(a Assertion) Assertion {
	return func(actual interface{}, expected ...interface{}) (fail string) {
		failureMessage := a(actual, expected...)
		if failureMessage == "" {
			return "Expected a failure"
		}
		return ""
	}
}

// AlwaysPass succeeds no matter what.
func AlwaysPass(actual interface{}, expected ...interface{}) (fail string) {
	return ""
}

// AlwaysFail succeeds no matter what.
func AlwaysFail(actual interface{}, expected ...interface{}) (fail string) {
	return fmt.Sprintf("Forced Fail: %#v, %#v", actual, expected)
}

// Assertion Parameter Checking derived from SmartyStreets/assertions'
func exactly(count int, expected []interface{}) string {
	if len(expected) != count {
		return fmt.Sprintf("This assertion expects %d right-side params, not %d.", count, len(expected))
	}
	return Ok
}

func minimally(min int, expected []interface{}) string {
	if len(expected) < min {
		return fmt.Sprintf("This assertion expects at least %d right-side params, not %d.", min, len(expected))
	}
	return Ok
}

// FailFirst returns the first non-blank failure string in a list of
// assertion returns.
func FailFirst(msgs ...string) string {
	for _, m := range msgs {
		if m != "" {
			return m
		}
	}
	return ""
}
