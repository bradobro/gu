package should

import (
	"testing"

	"github.com/kindrid/gotest/debug"
)

// HELPERS for self testing

func Passes(t *testing.T, topic string, a Assertion, actual interface{}, expected ...interface{}) {
	fail := a(actual, expected...)
	if fail == "" {
		return
	}
	t.Errorf("%s Expected %#v to pass but got '%s'", topic, a, fail)
	t.Errorf(debug.FormattedCallStack(2, 3))
	// if fail != "" {
	// 	t.Errorf("Expected %v to pass. Instead got '%s'.", a, fail)
	// } else if pass == "" {
	// 	t.Errorf("Expected %v to pass, but got a blank pass messasge.", a)
	// }
}

func Fails(t *testing.T, topic string, a Assertion, actual interface{}, expected ...interface{}) {
	fail := a(actual, expected...)
	if fail == "" {
		t.Errorf("Expected %#v (%s) to fail. Instead it passed.", a, topic)
	}
}

func TestAssertion(t *testing.T) {
	t.Run("Assertion fundamentals", testFundamentals)
}

func testFundamentals(t *testing.T) {
	// Forced passes and failures act correctly
	Passes(t, "Forced Pass", AlwaysPass, "PASS!", nil, nil)
	Fails(t, "Forced Fail", AlwaysFail, "FAIL!", nil, nil)

	// Negations correctly invert forced passes and failures
	Passes(t, "Negated Forced Fail", Not(AlwaysFail), "FAIL!", nil, nil)
	Fails(t, "Negated Forced Pass", Not(AlwaysPass), "PASS!", nil, nil)
}
