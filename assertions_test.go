package au_test

import (
	"testing"

	"github.com/bradobro/au"
)

func TestAssertionHelpers(t *testing.T) {
	params := []interface{}{"a", 1, struct{}{}}
	assertEquals(t, au.Needs(2, params), "This assertion requires exactly 2 params (you provided 3).")
	assertEquals(t, au.Needs(3, params), "")
	assertEquals(t, au.Needs(4, params), "This assertion requires exactly 4 params (you provided 3).")
	assertEquals(t, au.NeedsAtLeast(2, params), "")
	assertEquals(t, au.NeedsAtLeast(3, params), "")
	assertEquals(t, au.NeedsAtLeast(4, params), "This assertion requires at least 4 params (you provided 3).")
	assertEquals(t, au.NeedsAtMost(2, params), "This assertion allows 2 or fewer comparison values (you provided 3).")
	assertEquals(t, au.NeedsAtMost(3, params), "")
	assertEquals(t, au.NeedsAtMost(4, params), "")
}

func TestFailAndSkip(t *testing.T) {
	assertEquals(t, au.Fail(1, 4, 3, 1), "forced failure")
	assertEquals(t, au.Skip(au.Equal, 4, 3, 1), "")
}
func TestEqualAndUnequal(t *testing.T) {
	assertEquals(t, au.Equal(1, 1, 1, 1, 1, 1, 1), "")
	assertEquals(t, au.Equal(1), "This assertion requires at least 2 params (you provided 1).")
	assertEquals(t, au.Equal(1, 1, 1, 1, 1, 0), "expected 1 == 0")
	assertEquals(t, au.Unequal(1, 0, 0, -1, 2, 5, 9), "")
	assertEquals(t, au.Unequal(1, 0, 0, 1, 2, 5, 9), "expected 1 != 1")
	assertEquals(t, au.Unequal(1), "This assertion requires at least 2 params (you provided 1).")
}
func TestTrueFalse(t *testing.T) {
	assertEquals(t, au.True(true, true, true), "")
	assertEquals(t, au.True(true, true, false), "expecting true, got false")
	assertEquals(t, au.True(true, "a string", true), "expecting a bool, got \"a string\".(string)")
	assertEquals(t, au.False(false, false, false), "")
	assertEquals(t, au.False(false, false, true), "expecting false, got true")
	assertEquals(t, au.False(false, "a string", false), "expecting a bool, got \"a string\".(string)")
}
func TestTrueNilNotNil(t *testing.T) {
	assertEquals(t, au.Nil(nil, nil, nil), "")
	assertEquals(t, au.Nil(nil, "nil", nil), "should be nil")
	assertEquals(t, au.NotNil("nil", "nil", "nil"), "")
	assertEquals(t, au.NotNil("nil", "nil", nil), "should not be nil")

}
