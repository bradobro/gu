package should

import (
	"reflect"
	"testing"
)

const jsonSimpleObject = `{"x": 1, "camelCase": true, "type1Condition": true}`
const jsonObject = `{"a": 1, "b": [1,2,3], "c": true, "d": "yes", "nested": ` + jsonSimpleObject + `}`
const jsonArray = "[" + jsonObject + "," + jsonObject + "," + jsonObject + "]"
const snakeyObject = `{"a": 1, "b": [1,2,3], "snake_case": true, "d": "yes", "nested": ` + jsonSimpleObject + `}`
const snakeyArray = "[" + jsonObject + "," + snakeyObject + "," + jsonObject + "]"

//TestJSON exercise the JSON tests
func TestJSON(t *testing.T) {
	t.Parallel()
	t.Run("should.BeJSON", testBeJSON)
	t.Run("should.HaveFields", testHaveFields)
	t.Run("should.HaveOnlyCamelcaseKeys", testCamelcaseKeys)
}

func testHaveFields(t *testing.T) {
	var simpleFields = []interface{}{"a", reflect.Float64, "b", reflect.Slice, "c",
		reflect.Bool, "d", reflect.String, "nested", reflect.Map}
	var wrongFields = []interface{}{"a", reflect.String, "c", reflect.Float64, "d",
		reflect.Bool, "g", reflect.Interface}
	Passes(t, "Should have fields pass",
		HaveFields, jsonObject, simpleFields...)
	Fails(t, "Should have fields fail",
		HaveFields, jsonObject, wrongFields...)
	Passes(t, "Only have fields pass", HaveOnlyFields, jsonObject, simpleFields...)
	Fails(t, "Only have fields pass", HaveOnlyFields, jsonObject, "a", reflect.String)
	Fails(t, "Only have fields pass", HaveOnlyFields, jsonObject, "z", reflect.String)

	// Test documentation
	Passes(t, "Self documents if passed nil", StartWith, HaveFields(nil), "HaveFields expects")
}

func testBeJSON(t *testing.T) {
	// Test the happy paths.
	Passes(t, "A simple object should parse", BeJSON, jsonSimpleObject)
	Passes(t, "A compound object should parse", BeJSON, jsonObject)
	Passes(t, "An array should parse", BeJSON, jsonArray)
	Passes(t, "An []byte of the same array should parse", BeJSON, []byte(jsonArray))

	// Test failing paths.
	Fails(t, "Malformed objects should fail", BeJSON, jsonSimpleObject[1:])
	Fails(t, "Malformed ARRAYS should fail", BeJSON, jsonArray[:len(jsonArray)-1])
	Fails(t, "Nil fails", BeJSON, nil)
	Fails(t, "Non-strings fail", BeJSON, 5)

	// Test re-use of container.
	parsedJSONContainer, err := parseJSON(jsonArray)
	Passes(t, "...checking parse", BeNil, err)
	Passes(t, "A pre-parsed object should pass through.", BeJSON, parsedJSONContainer)
}

func testCamelcaseKeys(t *testing.T) {
	Passes(t, "A compound object with camelCase keys should pass", HaveOnlyCamelcaseKeys, jsonObject)
	Passes(t, "An array holding objects with camelCase keys should pass", HaveOnlyCamelcaseKeys, jsonArray)
	Fails(t, "A compound object with snake_case keys should fail", HaveOnlyCamelcaseKeys, snakeyObject)
	Fails(t, "An array holding objects with snake_case keys should fail", HaveOnlyCamelcaseKeys, snakeyArray)
	Passes(t, "A compound object ignoring it's snake_case keys should pass", HaveOnlyCamelcaseKeys, snakeyObject, "snake_case")
	Passes(t, "An array holding objects ignoring it's snake_case keys should pass", HaveOnlyCamelcaseKeys, snakeyArray, "snake_case")
	Fails(t, "Asking the camelcaser to ignore a non-string fails.", HaveOnlyCamelcaseKeys, jsonObject, make(map[int]int))
}
