package gocu

import "testing"

func TestSuite(t *testing.T) {
	suite := &Suite{
		FeaturePath: "testdata",
	}
	suite.Export()
}
