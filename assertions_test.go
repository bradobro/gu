package gu_test

import (
	"testing"

	"github.com/bradobro/gu"
)

func TestAssertionHelpers(t *testing.T) {
	params := []interface{}{"a", 1, struct{}{}}
	assertEquals(t, gu.Needs(2, params), "This assertion requires exactly 2 params (you provided 3).")
	assertEquals(t, gu.Needs(3, params), "")
	assertEquals(t, gu.Needs(4, params), "This assertion requires exactly 4 params (you provided 3).")
	assertEquals(t, gu.NeedsAtLeast(2, params), "")
	assertEquals(t, gu.NeedsAtLeast(3, params), "")
	assertEquals(t, gu.NeedsAtLeast(4, params), "This assertion requires at least 4 params (you provided 3).")
	assertEquals(t, gu.NeedsAtMost(2, params), "This assertion allows 2 or fewer comparison values (you provided 3).")
	assertEquals(t, gu.NeedsAtMost(3, params), "")
	assertEquals(t, gu.NeedsAtMost(4, params), "")
}

func TestFailAndSkip(t *testing.T) {
	assertEquals(t, gu.Fail(1, 4, 3, 1), "forced failure")
	assertEquals(t, gu.Skip(gu.Equal, 4, 3, 1), "")
}
func TestEqualAndUnequal(t *testing.T) {
	assertEquals(t, gu.Equal(1, 1, 1, 1, 1, 1, 1), "")
	assertEquals(t, gu.Equal(1), "This assertion requires at least 2 params (you provided 1).")
	assertEquals(t, gu.Equal(1, 1, 1, 1, 1, 0), "expected 1 == 0")
	assertEquals(t, gu.Unequal(1, 0, 0, -1, 2, 5, 9), "")
	assertEquals(t, gu.Unequal(1, 0, 0, 1, 2, 5, 9), "expected 1 != 1")
	assertEquals(t, gu.Unequal(1), "This assertion requires at least 2 params (you provided 1).")
}
func TestTrueFalse(t *testing.T) {
	assertEquals(t, gu.True(true, true, true), "")
	assertEquals(t, gu.True(true, true, false), "expecting true, got false")
	assertEquals(t, gu.True(true, "a string", true), "expecting a bool, got \"a string\".(string)")
	assertEquals(t, gu.False(false, false, false), "")
	assertEquals(t, gu.False(false, false, true), "expecting false, got true")
	assertEquals(t, gu.False(false, "a string", false), "expecting a bool, got \"a string\".(string)")
}
func TestTrueNilNotNil(t *testing.T) {
	assertEquals(t, gu.Nil(nil, nil, nil), "")
	assertEquals(t, gu.Nil(nil, "nil", nil), "should be nil")
	assertEquals(t, gu.NotNil("nil", "nil", "nil"), "")
	assertEquals(t, gu.NotNil("nil", "nil", nil), "should not be nil")

}
