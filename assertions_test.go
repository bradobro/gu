package cyu_test

import (
	"testing"

	"github.com/bradobro/cyu"
)

func TestAssertions(t *testing.T) {
	tstr := cyu.NewAsserter(t, false, 10, cyu.VerbosityInsane)
	tstr.Assert(cyu.Equal, 1, 1, 1, 1, 1, 1, 1)
	tstr.Assert(cyu.Unequal, 1, 1, 0, 1, 1, 1, 1)
}
