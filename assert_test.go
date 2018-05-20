package cyu_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/bradobro/cyu"
)

func assertTrue(t cyu.T, actual bool, msg string) {
	if !actual {
		t.Errorf("Expected true: %s", msg)
	}
}

func assertEquals(t cyu.T, actual, expected string) {
	if actual != expected {
		t.Errorf("Actual %#v\nExpected %#v", actual, expected)
	}
}

func testStringContains(t cyu.T, actual, contains string) {
	if !strings.Contains(actual, contains) {
		t.Errorf("%#v\ndoes not contain %#v", actual, contains)
	}
}

/*CustomT wraps a *testing.T in ways that test packages can be tested.
 */
type CustomT struct {
	t          cyu.T
	NoFail     bool // if true, t.Fail* and t.Error* only mark a flag
	Failed     bool
	FailedFast bool
	Writer     io.Writer
}

func NewT(t cyu.T, noFail bool, writer io.Writer) *CustomT {
	ct := &CustomT{t: t, NoFail: noFail, Writer: writer}
	return ct
}

func newTestT(t cyu.T) (*CustomT, *bytes.Buffer) {
	var buf bytes.Buffer
	ct := NewT(t, true, &buf)
	return ct, &buf
}

func (t *CustomT) Error(args ...interface{}) {
	t.Log(args...)
	t.Fail()
}

func (t *CustomT) Errorf(format string, args ...interface{}) {
	t.Logf(format, args...)
	t.Fail()
}

func (t *CustomT) Fail() {
	if t.NoFail {
		t.Failed = true
	} else {
		t.t.Fail()
	}
}

func (t *CustomT) FailNow() {
	if t.NoFail {
		t.Failed = true
		t.FailedFast = true
	} else {
		t.t.FailNow()
	}
}

func (t *CustomT) Log(args ...interface{}) {
	if t.Writer == nil {
		t.t.Log(args...)
	} else {
		fmt.Fprint(t.Writer, args...)
	}
}

func (t *CustomT) Name() string {
	if gt, ok := t.t.(*testing.T); ok {
		return gt.Name()
	}
	return ""
}

func (t *CustomT) Logf(format string, args ...interface{}) {
	if t.Writer == nil {
		t.t.Log(args...)
	} else {
		fmt.Fprintf(t.Writer, format, args...)
	}
}

func TestAssert(t *testing.T) {
	ct, _ := newTestT(t)

	tt := cyu.NewAsserter(false, 4, cyu.VerbosityInsane)
	x := 5
	y := x + 1

	tt.Assert(ct, cyu.Equal, x, x)
	tt.Assert(ct, cyu.Unequal, x, y)
	assertTrue(t, !(ct.Failed || ct.FailedFast), "these tests shouldn't have failed")

	tt.Assert(ct, cyu.Equal, x, y)
	assertTrue(t, ct.Failed, "this test should have failed")
	assertTrue(t, !ct.FailedFast, "this test should not have failed fast")

	ct.Failed = false
	tt.FailFast = true
	tt.Assert(ct, cyu.Equal, x, y)
	assertTrue(t, ct.Failed, "this test should have failed")
	assertTrue(t, ct.FailedFast, "this test should have failed fast")
}

// func TestTweakStackDepth(t *testing.T) {
// 	a := cyu.NewAsserter(t, false, 25, cyu.VerbosityInsane)
// 	a.Assert(cyu.Equal, 2, 5)
// }
