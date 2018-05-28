package gu_test

import (
	"bytes"
	"errors"
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
	// invalid assertion functions
	assertEquals(t, gu.Apply(1, 2, 3, 4).Error(), "this test library expects an assertion function returning only an error, got 1 (type int)")
	assertStringContains(t, gu.Apply(t).Error(), "*testing.T")
	assertStringContains(t, gu.Apply(func() {}).Error(), "func()")
	assertStringContains(t, gu.Apply(func() int { return 0 }).Error(), "func()")

	// Valid assertion functions
	assertNil(t, gu.Apply(func() error { return nil }))
	assertEquals(t, gu.Apply(func() error {
		return errors.New("TestApply")
	}).Error(), "TestApply")
	assertEquals(t,
		gu.Apply(func(a, b, c string) error {
			return errors.New(a + b + c)
		}, "ab", "cd", "ef").Error(),
		"abcdef")
}

func TestApplyArgs(t *testing.T) {
	// a function with fixed arguments
	fixed := func(a, b int, c string) error {
		if a == 0 {
			return nil
		}
		return fmt.Errorf("%d.%d.%s", a, b, c)
	}
	assertNil(t, gu.Apply(fixed, 0, 4, "a"))                                                               // no error
	assertEquals(t, gu.Apply(fixed, 1, 2, "b").Error(), "1.2.b")                                           // actual error
	assertEquals(t, gu.Apply(fixed, 1, 2).Error(), "test function expecting 3 args, got 2")                // too few args
	assertEquals(t, gu.Apply(fixed, 1, 2, "b", "c").Error(), "test function expecting 3 args, got 4")      // too many args
	assertEquals(t, gu.Apply(fixed, 1, "a", "b").Error(), "arg 1 (\"a\") not assignable to param 1 (int)") // wrong types args

	// a function with untyped variadics
	untyped := func(a, b int, c ...interface{}) error {
		if a+b-2 == len(c) {
			return nil
		}
		return fmt.Errorf("len(args) != %d + %d", a, b)
	}
	// no error
	assertNil(t, gu.Apply(untyped, 0, 3, "a"))
	// no error, mixed types
	assertNil(t, gu.Apply(untyped, 0, 4, "a", 64.2))
	// failure
	assertEquals(t, gu.Apply(untyped, 0, 4, "a").Error(), "len(args) != 0 + 4")
	// bad types
	assertEquals(t, gu.Apply(untyped, 0, "a", 60.4).Error(), "arg 1 (\"a\") not assignable to param 1 (int)")

	// typed variadic
	typed := func(a, b int, c ...string) error {
		if a+b-2 == len(c) {
			return nil
		}
		return fmt.Errorf("len(args) != %d + %d", a, b)
	}
	// no error
	assertNil(t, gu.Apply(typed, 0, 3, "a"))
	// failure
	assertEquals(t, gu.Apply(typed, 0, 4, "a").Error(), "len(args) != 0 + 4")
	// error in variadic types
	assertEquals(t, gu.Apply(typed, 0, 4, "a", 64.2).Error(), "arg 3 (64.2) not assignable to variadic param 3 (string)")

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
