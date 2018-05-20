package cyu

import (
	"fmt"
	"reflect"
)

/* Assertion Helpers */

// IsEqual returns "" if a equals b, otherwise it may explain the inequality
func IsEqual(a, b interface{}) string {
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

func Needs(needed int, params []interface{}) string {
	if len(params) != needed {
		return fmt.Sprintf(needExactValues, needed, len(params))
	}
	return success
}

func NeedsAtLeast(minimum int, params []interface{}) string {
	if len(params) < minimum {
		return fmt.Sprintf(needsMoreValues, minimum, len(params))
	}
	return success
}

func NeedsAtMost(max int, params []interface{}) string {
	if len(params) > max {
		return fmt.Sprintf(needFewerValues, max, len(params))
	}
	return success
}

/* Basic Assertions */

func Fail(params ...interface{}) string {
	return "forced failure"
}

func Skip(params ...interface{}) string {
	return ""
}

func Nil(params ...interface{}) (fail string) {
	for _, x := range params {
		if x != nil {
			return "should be nil"
		}
	}
	return
}

func NotNil(params ...interface{}) (fail string) {
	for _, x := range params {
		if x == nil {
			return "should not be nil"
		}
	}
	return
}

var Passes = Nil

var Fails = NotNil

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
		if IsEqual(first, x) != "" {
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
		if IsEqual(first, x) == "" {
			return fmt.Sprintf("expected %#v != %#v", first, x)
		}
	}
	return
}
