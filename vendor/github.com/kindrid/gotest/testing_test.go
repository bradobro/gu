package gotest

import (
	"testing"

	"github.com/kindrid/gotest/should"
)

//TestTesting exercise the basic testing wrappers
func TestTesting(t *testing.T) {
	// Happy Path
	Assert(t, "forced", should.AlwaysPass)
	Deny(t, "forced", should.AlwaysFail)

	// Failure Paths
	tAssert := &testing.T{}
	Assert(t, tAssert.Failed(), should.BeFalse)
	Assert(tAssert, "forced", should.AlwaysFail)
	Assert(t, tAssert.Failed(), should.BeTrue)
	tDeny := &testing.T{}
	Assert(t, tDeny.Failed(), should.BeFalse)
	Deny(tDeny, "forced", should.AlwaysPass)
	Assert(t, tDeny.Failed(), should.BeTrue)
}
