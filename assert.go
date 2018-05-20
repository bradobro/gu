package gu

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

// Assertion is a generic test function. By convention its params are either
// actual...
// actual, expected...
// actual, configuration...
type Assertion func(params ...interface{}) (fail string)

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
func (assert *Asserter) AssertSkip(t T, skip int, assertf Assertion, params ...interface{}) {

	fail := assertf(params...)
	if fail == "" {
		return
	}
	assert.Reporter.Report(t, skip, fail, params)
	if assert.FailFast {
		assert.Reporter.Log(t, "Skipping remaining assertions for this test because of FailFast.\n")
		t.FailNow()
	}
}

// Assert reports errors, attempting to guess stack depth
func (assert *Asserter) Assert(t T, assertf Assertion, params ...interface{}) {
	assert.AssertSkip(t, 5, assertf, params...)
}
