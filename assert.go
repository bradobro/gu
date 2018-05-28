package gu

import (
	"fmt"
	"reflect"
)

// T describes the interface provided by Go's *testing.T.
type T interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow() // exit the current test immediately
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

// Namer exposes testing.T.Name(), new in Go 1.8
type Namer interface {
	Name() string // need Go 1.8 for this.
}

/* Assertion is convention for test functions.
It is declared as an interface because the function can be of any arity and
argument types so long as it returns an error.

A nil error means the assertion passes.
A non-nil error means the test failed, and the reporter will attempt to parse
the error's message (err.Error()) to control verbosity. See the Reporter struct
for details.
*/
type Assertion interface{}

// Apply attempts to apply args to an assertion function
func Apply(f Assertion, args ...interface{}) (err error) {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func || // it's a function
		(t.NumOut() != 1) { //|| // with a single return
		// (t.Out(0).Implements(reflect.TypeOf(err))) { // returning an error
		badFunc := "expecting a function returning an error, got %#v (type %T)"
		return fmt.Errorf(badFunc, f, f)
	}
	return
}

// Asserter tests assertions and reports on failures.
type Asserter struct {
	FailFast bool
	Reporter *Reporter
}

// NewAsserter creates an assertion reporter with reporting settings
func NewAsserter(failFast bool, maxDepth, verbosity int) (result *Asserter) {
	result = &Asserter{
		FailFast: failFast,
		Reporter: &Reporter{
			Verbosity: verbosity,
			MaxDepth:  maxDepth,
		},
	}
	return
}

// AssertSkip wraps any standard Assertion for use with Go's std.testing library
// skipping a given number of stack frames when reporting tracebacks.
func (assert *Asserter) AssertSkip(t T, skip int, assertion Assertion, params ...interface{}) {

	err := Apply(assertion, params...)
	if err == nil {
		return
	}
	assert.Reporter.Report(t, skip, err, params)
	if assert.FailFast {
		assert.Reporter.Log(t, "Skipping remaining assertions for this test because of FailFast.\n")
		t.FailNow()
	}
}

// Assert reports errors, attempting to guess stack depth
func (assert *Asserter) Assert(t T, assertion Assertion, params ...interface{}) {
	assert.AssertSkip(t, 5, assertion, params...)
}
