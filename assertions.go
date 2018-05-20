package cyu

import (
	"fmt"
	"reflect"
)

/* Assertion Helpers */

// isEqual returns "" if a equals b, otherwise it may explain the inequality
func isEqual(a, b interface{}) string {
	// for now we just use reflect.DeepEqual, but we may want to enhance this.
	if reflect.DeepEqual(a, b) {
		return ""
	}
	return fmt.Sprintf("do not equal")
}

const (
	success         = ""
	needExactValues = "This assertion requires exactly %d params (you provided %d)."
	needsMoreValues = "This assertion requires at least %d params (you provided %d)."
	needFewerValues = "This assertion allows %d or fewer comparison values (you provided %d)."
)

// Needs checks that an exact number of parameters is used
func Needs(needed int, params []interface{}) string {
	if len(params) != needed {
		return fmt.Sprintf(needExactValues, needed, len(params))
	}
	return success
}

// NeedsAtLeast ensures that at least minimum number of parameters is used
func NeedsAtLeast(minimum int, params []interface{}) string {
	if len(params) < minimum {
		return fmt.Sprintf(needsMoreValues, minimum, len(params))
	}
	return success
}

// NeedsAtMost checks that no more than a maximum number of parameters is used
func NeedsAtMost(max int, params []interface{}) string {
	if len(params) > max {
		return fmt.Sprintf(needFewerValues, max, len(params))
	}
	return success
}

/* Basic Assertions */

// Fail always fails
func Fail(params ...interface{}) string {
	return "forced failure"
}

// Skip skips a test
func Skip(params ...interface{}) string {
	return ""
}

// Nil fails if any of its params are not nil
func Nil(params ...interface{}) (fail string) {
	for _, x := range params {
		if x != nil {
			return "should be nil"
		}
	}
	return
}

// NotNil fails if any of its params are nil
func NotNil(params ...interface{}) (fail string) {
	for _, x := range params {
		if x == nil {
			return "should not be nil"
		}
	}
	return
}

// Passes might be clearer when checking for nil error values
var Passes = Nil

// Fails might be clearer when checking for non-nil error values
var Fails = NotNil

// True fails if any of its params are false
func True(params ...interface{}) (fail string) {
	for _, x := range params {
		if b, ok := x.(bool); !ok {
			return fmt.Sprintf("expecting a bool, got %#v.(%T)", x, x)
		} else if !b {
			return "expecting true, got false"
		}
	}
	return
}

// False fails if any of its params are true
func False(params ...interface{}) (fail string) {
	for _, x := range params {
		if b, ok := x.(bool); !ok {
			return fmt.Sprintf("expecting a bool, got %#v.(%T)", x, x)
		} else if b {
			return "expecting false, got true"
		}
	}
	return
}

// Equal fails if the first param does not equal all the later params.
func Equal(params ...interface{}) (fail string) {
	if fail = NeedsAtLeast(2, params); fail != success {
		return
	}
	first := params[0]
	for _, x := range params[1:] {
		if isEqual(first, x) != "" {
			return fmt.Sprintf("expected %#v == %#v", first, x)
		}
	}
	return
}

// Unequal fails if the first param equals any of the later params
func Unequal(params ...interface{}) (fail string) {
	if fail = NeedsAtLeast(2, params); fail != success {
		return
	}
	first := params[0]
	for _, x := range params[1:] {
		if isEqual(first, x) == "" {
			return fmt.Sprintf("expected %#v != %#v", first, x)
		}
	}
	return
}
