package gu

import (
	"fmt"
	"reflect"
)

const DefaultSkip = 4 // stack frames to skip in normal assert

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

/*Assertion is convention for test functions.
It is declared as an interface because the function can be of any arity and
argument types so long as it returns an error.

A nil error means the assertion passes.
A non-nil error means the test failed, and the reporter will attempt to parse
the error's message (err.Error()) to control verbosity. See the Reporter struct
for details.
*/
type Assertion interface{}

var errorType = reflect.TypeOf((*error)(nil)).Elem()

// Apply attempts to apply args to an assertion function
func Apply(f Assertion, args ...interface{}) (err error) {
	var inputs, outputs []reflect.Value
	t := reflect.TypeOf(f)
	badFunc := fmt.Errorf("this test library expects an assertion function returning only an error, got %#v (type %T)", f, f)

	// make sure f is a function
	if t.Kind() != reflect.Func {
		return badFunc
	}

	// run the function
	if inputs, err = checkedInputs(t, args); err != nil {
		return // problem with args
	}
	outputs = reflect.ValueOf(f).Call(inputs)
	if len(outputs) != 1 {
		return badFunc
	}
	// check and convert the output to an error
	output := outputs[0].Interface()
	if output == nil {
		return nil
	}
	if result, ok := output.(error); ok {
		return result
	}
	return badFunc
}

// checkedInputs ensures the assertion function is able to accept the args provided
func checkedInputs(t reflect.Type, args []interface{}) (vals []reflect.Value, err error) {
	nParams := t.NumIn()
	nArgs := len(args)
	if t.IsVariadic() {
		nParams--
		if nArgs < nParams {
			return nil, fmt.Errorf("test function expecting at least %d args, got %d", nParams, nArgs)
		}
	} else if nArgs != nParams {
		return nil, fmt.Errorf("test function expecting %d args, got %d", nParams, nArgs)
	}
	vals = make([]reflect.Value, nArgs)
	for i := 0; i < nArgs; i++ {
		vals[i] = reflect.ValueOf(args[i])
		if i >= nParams { // handle variadics
			if !vals[i].Type().AssignableTo(t.In(nParams).Elem()) {
				return nil, fmt.Errorf("arg %d (%#v) not assignable to variadic param %d (%s)", i, args[i], i, t.In(nParams).Elem())
			}
		} else if !vals[i].Type().AssignableTo(t.In(i)) {
			return nil, fmt.Errorf("arg %d (%#v) not assignable to param %d (%s)", i, args[i], i, t.In(i))
		}
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
	if err := Apply(assertion, params...); err != nil {
		assert.Reporter.Report(t, skip, err, params)
		if assert.FailFast {
			assert.Reporter.Log(t, "Skipping remaining assertions for this test because of FailFast.\n")
			t.FailNow()
		}
	}
}

// Assert reports errors, attempting to guess stack depth
func (assert *Asserter) Assert(t T, assertion Assertion, params ...interface{}) {
	assert.AssertSkip(t, DefaultSkip, assertion, params...)
}
