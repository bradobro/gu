package cyu_test

import (
	"testing"

	"github.com/bradobro/cyu"
)

func TestAssert(t *testing.T) {
	tt := cyu.NewAsserter(t, false, 4, cyu.VerbosityInsane)
	x := 5
	y := x + 1

	tt.Assert(cyu.Equal, x, x)
	tt.Assert(cyu.Unequal, x, y)
}
