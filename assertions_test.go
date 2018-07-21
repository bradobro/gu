package gu_test

import (
	"testing"

	"github.com/bradobro/gu"
)

func TestAssertionHelpers(t *testing.T) {
	params := []interface{}{"a", 1, struct{}{}}
	assertEquals(t, gu.ParamCount(2, params).Error(), "This assertion requires exactly 2 params (you provided 3).")
	assertNil(t, gu.ParamCount(3, params))
	assertEquals(t, gu.ParamCount(4, params).Error(), "This assertion requires exactly 4 params (you provided 3).")
	assertNil(t, gu.ParamMin(2, params))
	assertNil(t, gu.ParamMin(3, params))
	assertEquals(t, gu.ParamMin(4, params).Error(), "This assertion requires at least 4 params (you provided 3).")
	assertEquals(t, gu.ParamMax(2, params).Error(), "This assertion allows 2 or fewer comparison values (you provided 3).")
	assertNil(t, gu.ParamMax(3, params))
	assertNil(t, gu.ParamMax(4, params))
}

func TestFailAndSkip(t *testing.T) {
	assertEquals(t, gu.Fail(1, 4, 3, 1).Error(), "forced failure")
	assertNil(t, gu.Skip(gu.Equal, 4, 3, 1))
}
func TestEqualAndUnequal(t *testing.T) {
	assertNil(t, gu.Equal(1, 1, 1, 1, 1, 1, 1))
	assertEquals(t, gu.Equal(1).Error(), "This assertion requires at least 2 params (you provided 1).")
	assertEquals(t, gu.Equal(1, 1, 1, 1, 1, 0).Error(), "expected 1 == 0")
	assertNil(t, gu.Unequal(1, 0, 0, -1, 2, 5, 9))
	assertEquals(t, gu.Unequal(1, 0, 0, 1, 2, 5, 9).Error(), "expected 1 != 1")
	assertEquals(t, gu.Unequal(1).Error(), "This assertion requires at least 2 params (you provided 1).")
}

type eqStruct struct {
	I int
	F float32
	L float64
	S string
	A []int
}

func TestStructEquality(t0 *testing.T) {
	// primitive fixtures
	A := eqStruct{5, 5.0, 5.0, "five", []int{0, 1, 2, 3, 5}}
	dup := eqStruct{5, 5.0, 5.0, "five", []int{0, 1, 2, 3, 5}}
	badF := eqStruct{5, 5.2, 5.0, "five", []int{0, 1, 2, 3, 5}}
	badA := eqStruct{5, 5.0, 5.0, "five", []int{1, 0, 2, 3, 5}}
	badS := eqStruct{5, 5.0, 5.0, "four", []int{0, 1, 2, 3, 5}}

	// build test table
	table := []struct {
		desc      string
		assertion func(params ...interface{}) (err error)
		actual    interface{}
		expected  interface{}
	}{
		// identities
		{"same struct", gu.Equal, A, A},
		{"equivalent struct", gu.Equal, A, dup},
		{"nil == nil", gu.Equal, nil, nil},
		{"same struct pointer", gu.Equal, &A, &A},
		{"equivalent struct pointer", gu.Equal, &A, &dup},

		// pointer details
		{"struct to nil", gu.Unequal, A, nil},
		{"struct to nil", gu.Unequal, nil, A},
		{"struct pointers not dereferenced", gu.Unequal, A, &A},

		// element differences
		{"arrays differ", gu.Unequal, A, badA},
		{"floats differ", gu.Unequal, A, badF},
		{"strings differ", gu.Unequal, A, badS},
		{"ptrs to arrays differ", gu.Unequal, &A, &badA},
		{"ptrs to floats differ", gu.Unequal, &A, &badF},
		{"ptrs to strings differ", gu.Unequal, &A, &badS},
		{"sanity check on arrays differ", gu.Equal, badA, badA},
		{"sanity check on floats differ", gu.Equal, badF, badF},
		{"sanity check on strings differ", gu.Equal, badS, badS},
	}

	// test against the table
	t0.Parallel() // run subtests in parallel
	for _, tst := range table {
		t0.Run(tst.desc, func(t *testing.T) {
			if err := tst.assertion(tst.actual, tst.expected); err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func TestTrueFalse(t *testing.T) {
	assertNil(t, gu.True(true, true, true))
	assertEquals(t, gu.True(true, true, false).Error(), "expecting true, got false")
	assertEquals(t, gu.True(true, "a string", true).Error(), "expecting a bool, got \"a string\".(string)")
	assertNil(t, gu.False(false, false, false))
	assertEquals(t, gu.False(false, false, true).Error(), "expecting false, got true")
	assertEquals(t, gu.False(false, "a string", false).Error(), "expecting a bool, got \"a string\".(string)")
}
func TestTrueNilNotNil(t *testing.T) {
	assertNil(t, gu.Nil(nil, nil, nil))
	assertEquals(t, gu.Nil(nil, "nil", nil).Error(), "should be nil")
	assertNil(t, gu.NotNil("nil", "nil", "nil"))
	assertEquals(t, gu.NotNil("nil", "nil", nil).Error(), "should not be nil")

}
