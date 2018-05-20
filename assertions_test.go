package cyu_test

import (
	"testing"

	"github.com/bradobro/cyu"
)

func TestAssertionHelpers(t *testing.T) {
	params := []interface{}{"a", 1, struct{}{}}
	testStringEqual(t, cyu.Needs(2, params), "This assertion requires exactly 2 params (you provided 3).")
	testStringEqual(t, cyu.Needs(3, params), "")
	testStringEqual(t, cyu.Needs(4, params), "This assertion requires exactly 4 params (you provided 3).")
	testStringEqual(t, cyu.NeedsAtLeast(2, params), "")
	testStringEqual(t, cyu.NeedsAtLeast(3, params), "")
	testStringEqual(t, cyu.NeedsAtLeast(4, params), "This assertion requires at least 4 params (you provided 3).")
	testStringEqual(t, cyu.NeedsAtMost(2, params), "This assertion allows 2 or fewer comparison values (you provided 3).")
	testStringEqual(t, cyu.NeedsAtMost(3, params), "")
	testStringEqual(t, cyu.NeedsAtMost(4, params), "")
}

func TestFailAndSkip(t *testing.T) {
	testStringEqual(t, cyu.Fail(1, 4, 3, 1), "forced failure")
	testStringEqual(t, cyu.Skip(cyu.Equal, 4, 3, 1), "")
}
func TestEqualAndUnequal(t *testing.T) {
	testStringEqual(t, cyu.Equal(1, 1, 1, 1, 1, 1, 1), "")
	testStringEqual(t, cyu.Equal(1), "This assertion requires at least 2 params (you provided 1).")
	testStringEqual(t, cyu.Equal(1, 1, 1, 1, 1, 0), "expected 1 == 0")
	testStringEqual(t, cyu.Unequal(1, 0, 0, -1, 2, 5, 9), "")
	testStringEqual(t, cyu.Unequal(1, 0, 0, 1, 2, 5, 9), "expected 1 != 1")
	testStringEqual(t, cyu.Unequal(1), "This assertion requires at least 2 params (you provided 1).")
}
func TestTrueFalse(t *testing.T) {
	testStringEqual(t, cyu.True(true, true, true), "")
	testStringEqual(t, cyu.True(true, true, false), "expecting true, got false")
	testStringEqual(t, cyu.True(true, "a string", true), "expecting a bool, got \"a string\".(string)")
	testStringEqual(t, cyu.False(false, false, false), "")
	testStringEqual(t, cyu.False(false, false, true), "expecting false, got true")
	testStringEqual(t, cyu.False(false, "a string", false), "expecting a bool, got \"a string\".(string)")
}
func TestTrueNilNotNil(t *testing.T) {
	testStringEqual(t, cyu.Nil(nil, nil, nil), "")
	testStringEqual(t, cyu.Nil(nil, "nil", nil), "should be nil")
	testStringEqual(t, cyu.NotNil("nil", "nil", "nil"), "")
	testStringEqual(t, cyu.NotNil("nil", "nil", nil), "should not be nil")

}
