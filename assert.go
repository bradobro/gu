package cyu

// T describes the interface provided by Go's *testing.T.
type T interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow() // exit the current test immediately
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

type Namer interface {
	Name() string // need Go 1.8 for this.
}

// Assertion is a generic test function. By convention its params are either
// actual...
// actual, expected...
// actual, configuration...
type Assertion func(params ...interface{}) (fail string)

type Asserter struct {
	t        T
	FailFast bool
	Reporter *Reporter
}

func NewAsserter(t T, failFast bool, maxDepth, verbosity int) (result *Asserter) {
	result = &Asserter{
		t:        t,
		FailFast: failFast,
		Reporter: &Reporter{
			T:         t,
			Verbosity: verbosity,
			MaxDepth:  maxDepth,
		},
	}
	return
}

// AssertSkip wraps any standard Assertion for use with Go's std.testing library
// skipping a given number of stack frames when reporting tracebacks.
func (t *Asserter) AssertSkip(skip int, assert Assertion, params ...interface{}) {

	fail := assert(params...)
	if fail == "" {
		return
	}
	t.Reporter.Report(skip, fail, params)
	if t.FailFast {
		t.Reporter.Log("Skipping remaining assertions for this test because of FailFast.\n")
		t.t.FailNow()
	}
}

// Assert reports errors, attempting to guess stack depth
func (t *Asserter) Assert(assert Assertion, params ...interface{}) {
	t.AssertSkip(4, assert, params...)
}
