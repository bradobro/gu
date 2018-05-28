package gu_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/bradobro/gu"
)

func assertTrue(t gu.T, actual bool, msg string) {
	if !actual {
		t.Errorf("Expected true: %s", msg)
	}
}

func assertNil(t gu.T, actual interface{}) {
	if actual != nil {
		t.Errorf("Expected nil: %#v", actual)
	}
}

func assertEquals(t gu.T, actual, expected string) {
	if actual != expected {
		t.Errorf("Actual %#v\nExpected %#v", actual, expected)
	}
}

func assertStringContains(t gu.T, actual, contains string) {
	if !strings.Contains(actual, contains) {
		t.Errorf("%#v\ndoes not contain %#v", actual, contains)
	}
}

/*CustomT wraps a *testing.T in ways that test packages can be tested.
 */
type CustomT struct {
	t          gu.T
	NoFail     bool // if true, t.Fail* and t.Error* only mark a flag
	Failed     bool
	FailedFast bool
	Writer     io.Writer
}

func NewT(t gu.T, noFail bool, writer io.Writer) *CustomT {
	ct := &CustomT{t: t, NoFail: noFail, Writer: writer}
	return ct
}

func newTestT(t gu.T) (*CustomT, *bytes.Buffer) {
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

func TestApply(t *testing.T) {
	assertEquals(t, gu.Apply(1, 2, 3, 4).Error(), "expecting a function returning an error, got 1 (type int)")
	assertStringContains(t, gu.Apply(t).Error(), "*testing.T")
	assertStringContains(t, gu.Apply(func() {}).Error(), "func()")
	// assertStringContains(t, gu.Apply(func() error { return nil }).Error(), "*testing.T")

}

func TestAssert(t *testing.T) {
	ct, _ := newTestT(t)

	tt := gu.NewAsserter(false, 4, gu.VerbosityInsane)
	x := 5
	y := x + 1

	tt.Assert(ct, func(a, b, c int) error {
		return nil
	}, 1, 2, 3)
	tt.Assert(ct, gu.Equal, x, x)
	tt.Assert(ct, gu.Unequal, x, y)
	assertTrue(t, !(ct.Failed || ct.FailedFast), "these tests shouldn't have failed")

	tt.Assert(ct, gu.Equal, x, y)
	assertTrue(t, ct.Failed, "this test should have failed")
	assertTrue(t, !ct.FailedFast, "this test should not have failed fast")

	ct.Failed = false
	tt.FailFast = true
	tt.Assert(ct, gu.Equal, x, y)
	assertTrue(t, ct.Failed, "this test should have failed")
	assertTrue(t, ct.FailedFast, "this test should have failed fast")
}

// func TestTweakStackDepth(t *testing.T) {
// 	a := gu.NewAsserter(t, false, 25, gu.VerbosityInsane)
// 	a.Assert(gu.Equal, 2, 5)
// }
