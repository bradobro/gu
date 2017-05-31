package cyu

import (
	"reflect"
	"regexp"
)

type Step struct {
	Prose  string          // the porose representation of the step
	Re     regexp.Regexp   // RE2 patten to match the step with matching groups for params
	Params []*string       // parameter names
	Types  []*reflect.Kind // if available, hints at parameter types
}
