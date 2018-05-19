package cyu

// T describes the interface provided by Go's *testing.T.
type T interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow() // exit the current test immediately
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

type Tester struct {
	t        T
	FailFast bool
	Reporter *Reporter
}

func NewTester(t T, failFast bool, maxDepth, verbosity int) (result *Tester) {
	result = &Tester{
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
func (t *Tester) AssertSkip(skip int, assert Assertion, params ...interface{}) {

	fail := assert(actual, params...)
	t.Reporter.Report(skip, fail, params) // debug levels report even if there isn't a failure
	if fail != "" && t.FailFast {
		t.Reporter.Writef("Skipping remaining assertions for this test because of FailFast.\n")
		t.t.FailNow()
	}
}

// Assert reports errors, attempting to guess stack depth
func (t *Tester) Assert(actual interface{}, f Assertion, expected ...interface{}) {
	t.AssertSkip(4, actual, f, expected)
}
