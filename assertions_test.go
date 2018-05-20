package cyu_test

import (
	"testing"

	"github.com/bradobro/cyu"
)

func TestAssertions(t *testing.T) {
	testStringEqual(t, cyu.Fail(1, 4, 3, 1), "forced failure")
	testStringEqual(t, cyu.Skip(cyu.Equal, 4, 3, 1), "")
	testStringEqual(t, cyu.Equal(1, 1, 1, 1, 1, 1, 1), "")
	testStringEqual(t, cyu.Equal(1, 1, 1, 1, 1, 0), "expected 1 == 0")
	testStringEqual(t, cyu.Unequal(1, 0, 0, -1, 2, 5, 9), "")
	testStringEqual(t, cyu.Unequal(1, 0, 0, 1, 2, 5, 9), "expected 1 != 1")

	testStringEqual(t, cyu.True(true, true, true), "")
	testStringEqual(t, cyu.True(true, true, false), "expecting true, got false")
	testStringEqual(t, cyu.False(false, false, false), "")
	testStringEqual(t, cyu.False(false, false, true), "expecting false, got true")
	testStringEqual(t, cyu.AllNil(nil, nil, nil), "")
	testStringEqual(t, cyu.AllNil(nil, "nil", nil), "should be nil")
	testStringEqual(t, cyu.NoneNil("nil", "nil", "nil"), "")
	testStringEqual(t, cyu.NoneNil("nil", "nil", nil), "should not be nil")
}
