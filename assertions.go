package gu

import (
	"errors"
	"fmt"
	"reflect"
)

/* Assertion Helpers */

// isEqual returns "" if a equals b, otherwise it may explain the inequality
func isEqual(a, b interface{}) error {
	// for now we just use reflect.DeepEqual, but we may want to enhance this.
	if reflect.DeepEqual(a, b) {
		return nil
	}
	return errors.New("do not equal")
}

const (
	needExactValues = "This assertion requires exactly %d params (you provided %d)."
	needsMoreValues = "This assertion requires at least %d params (you provided %d)."
	needFewerValues = "This assertion allows %d or fewer comparison values (you provided %d)."
)

// Needs checks that an exact number of parameters is used
func Needs(needed int, params []interface{}) error {
	if len(params) != needed {
		return fmt.Errorf(needExactValues, needed, len(params))
	}
	return nil
}

// NeedsAtLeast ensures that at least minimum number of parameters is used
func NeedsAtLeast(minimum int, params []interface{}) error {
	if len(params) < minimum {
		return fmt.Errorf(needsMoreValues, minimum, len(params))
	}
	return nil
}

// NeedsAtMost checks that no more than a maximum number of parameters is used
func NeedsAtMost(max int, params []interface{}) error {
	if len(params) > max {
		return fmt.Errorf(needFewerValues, max, len(params))
	}
	return nil
}

/* Basic Assertions */

// Fail always fails
func Fail(params ...interface{}) error {
	return errors.New("forced failure")
}

// Skip skips a test
func Skip(params ...interface{}) error {
	return nil
}

// Nil fails if any of its params are not nil
func Nil(params ...interface{}) (err error) {
	for _, x := range params {
		if x != nil {
			return errors.New("should be nil")
		}
	}
	return
}

// NotNil fails if any of its params are nil
func NotNil(params ...interface{}) (err error) {
	for _, x := range params {
		if x == nil {
			return errors.New("should not be nil")
		}
	}
	return
}

// Passes might be clearer when checking for nil error values
var Passes = Nil

// Fails might be clearer when checking for non-nil error values
var Fails = NotNil

// True fails if any of its params are false
func True(params ...interface{}) (err error) {
	for _, x := range params {
		if b, ok := x.(bool); !ok {
			return fmt.Errorf("expecting a bool, got %#v.(%T)", x, x)
		} else if !b {
			return errors.New("expecting true, got false")
		}
	}
	return
}

// False fails if any of its params are true
func False(params ...interface{}) (err error) {
	for _, x := range params {
		if b, ok := x.(bool); !ok {
			return fmt.Errorf("expecting a bool, got %#v.(%T)", x, x)
		} else if b {
			return errors.New("expecting false, got true")
		}
	}
	return
}

// Equal fails if the first param does not equal all the later params.
func Equal(params ...interface{}) (err error) {
	if err = NeedsAtLeast(2, params); err != nil {
		return
	}
	first := params[0]
	for _, x := range params[1:] {
		if isEqual(first, x) != nil {
			return fmt.Errorf("expected %#v == %#v", first, x)
		}
	}
	return
}

// Unequal fails if the first param equals any of the later params
func Unequal(params ...interface{}) (err error) {
	if err = NeedsAtLeast(2, params); err != nil {
		return
	}
	first := params[0]
	for _, x := range params[1:] {
		if isEqual(first, x) == nil {
			return fmt.Errorf("expected %#v != %#v", first, x)
		}
	}
	return
}
