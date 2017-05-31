package cyu

import (
	"fmt"
	"testing"
)

type StepFunc func(tbl Table, args ...string) (fail string)

// World defines a test context for using cyu tests in Golang. It should probably just be templated in...maybe?
type World interface {
	Step(step StepFunc, tbl Table, args ...string)
	Assert(t *testing.T)
	// BeforeSuite(t *testing.T) error
	// AfterSuite(t *testing.T) error
	// BeforeFeature(t *testing.T) error
	// AfterFeature(t *testing.T) error
	// BeforeScenario(t *testing.T) error
	// AfterScenario(t *testing.T) error

	// // Internal Key Value Stores
	// Clear()
	// SetString(name, value string)
	// GetString(name string) string
	// SetInt(name string, value int)
	// GetInt(name string) int
	// SetValue(name string, value interface{})
	// GetValue(name string) interface{}
	// SetTable(name string, value Table)
	// GetTable(name string) Table
}

// BaseWorld implements the World interface and is meant to be embedded in a particular suite's custom context struct
type BaseWorld struct {
	Fail string
}

// Step needs documenting
func (bw *BaseWorld) Step(step StepFunc, tbl Table, args ...string) {
	// could log step if verbose
	if bw.Fail == "" {
		bw.Fail = step(tbl, args...)
	}
}

// Assert needs documenting
func (bw *BaseWorld) Assert(t *testing.T) {
	if bw.Fail != "" {
		t.Error(bw.Fail)
	}
}

// CheckArgLength needs documenting
func (bw *BaseWorld) CheckArgLength(args []string, length int) (fail string) {
	if len(args) != length {
		fail = fmt.Sprintf("Expected %d args but got %v#.", length, args)
	}
	return
}

// Args2 needs documenting
func (bw *BaseWorld) Args2(args []string) (a1, a2, fail string) {
	if fail = bw.CheckArgLength(args, 2); fail == "" {
		a1, a2 = args[0], args[1]
	}
	return
}
