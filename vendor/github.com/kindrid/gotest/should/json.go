package should

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/Jeffail/gabs"
)

/* About the JSON parser: https://github.com/tidwall/gjson and
/* https://github.com/tidwall/gjson both fit most of our needs. Gjson is faster
/* most of the time, but uses unsafe and doesn't give distinct parse errors. */

// ParseJSON accepts JSON in various formats (including its output format) and returns traversable output.
func ParseJSON(actual interface{}) (StructureExplorer, error) {
	var result *GabsExplorer
	// gabs := &gabs.Container{}
	gabs, err := parseJSON(actual)
	if err != nil {
		return nil, err
	}
	result = (*GabsExplorer)(gabs)
	return result, nil
}

func parseJSON(actual interface{}) (*gabs.Container, error) {
	switch v := actual.(type) {
	case string:
		container, err := gabs.ParseJSON([]byte(v))
		if err != nil {
			return nil, fmt.Errorf(FormatFailure("Error parsing JSON.", err.Error(), "", ""))
		}
		return container, err
	case *string:
		container, err := gabs.ParseJSON([]byte(*v))
		if err != nil {
			return nil, fmt.Errorf(FormatFailure("Error parsing JSON.", err.Error(), "", ""))
		}
		return container, err
	case []byte:
		return gabs.ParseJSON(v)
	case *gabs.Container:
		return v, nil
	case *GabsExplorer: // until we convert the other tests over to use StructureExplorers
		return (*gabs.Container)(v), nil
	default:
		return nil, fmt.Errorf(
			FormatFailure(
				"JSON parser given a value it can't parse.",
				fmt.Sprintf("Expecting a JSON string or a structure representing one, not a %T.", actual),
				"", "",
			),
		)
	}
}

// HaveFields passes if the JSON container or string has fields with certain values or types of values:
//
//   HaveFields(json, "id", reflect.String)  // assert that there is a field `id` with a  string value.
//   HaveFields(json, "count", reflect.Float64)  // assert that there is a field `count` with a numeric value.
//   HaveFields(json, "count", 100)  // assert that there is a field `count` with a numeric value equal to 100.
//   HaveFields(json, "default", reflect.Interface)  // assert that there is a field `default` with any type of value.
//
func HaveFields(actual interface{}, expected ...interface{}) (fail string) {
	usage := "HaveFields expects parseable JSON to be compared to fieldPath string, fieldKind reflect.Kind pairs."
	if actual == nil {
		return usage
	}
	json, err := parseJSON(actual)
	if err != nil {
		return err.Error()
	}
	return haveFields(json, true, expected...)
}

// AllowFields passes if fields in the JSON container or string either don't exist or match expected types.
//
//   AllowFields(json, "id", reflect.String)  // assert that there is a field `id` with a  string value.
//   AllowFields(json, "count", reflect.Float64)  // assert that there is a field `count` with an numeric value.
//   AllowFields(json, "default", reflect.Interface)  // assert that there is a field `default` with any type of value.
//
func AllowFields(actual interface{}, expected ...interface{}) (fail string) {
	usage := "AllowFields expects parseable JSON to be compared to fieldPath string, fieldKind reflect.Kind pairs."
	if actual == nil {
		return usage
	}
	json, err := parseJSON(actual)
	if err != nil {
		return err.Error()
	}
	fail = haveFields(json, false, expected...)
	if fail != "" {
		fail = FormatFailure("JSON field(s) do not match expected types.", fail, "", "")
	}
	return
}

// haveFields checks to see if json contains fields and types matching expected.
// if required=false, it tolerates fields not appearing in the object.
// expected is [fieldPath string, fieldKind reflect.Kind, ...]  pairs
func haveFields(json *gabs.Container, required bool, expected ...interface{}) (fail string) {
	for i := 0; i < len(expected); i += 2 {
		// check existence of key
		fieldPath := expected[i].(string)
		container := json.Path(fieldPath)
		if container == nil || container.Data() == nil { // field not found
			if required {
				fail += fmt.Sprintf("Field '%s' is missing.\n%s", fieldPath, json)
			}
			continue
		}

		expectedInterface := expected[i+1]
		expectedKind, ok := expectedInterface.(reflect.Kind)
		if !ok { // We have a value, not a Kind
			return Equal(container.Data(), expectedInterface)
		}
		// check type of value
		if expectedKind == reflect.Interface { // allow any type
			continue
		}
		actualKind := reflect.ValueOf(container.Data()).Kind()
		if actualKind != expectedKind {
			fail += fmt.Sprintf("Expecting a '%s' value of type %s, got %s.\nJSON: %s", fieldPath, expectedKind, actualKind, json)
		}
	}

	return
}

// HaveOnlyFields passes if the JSON container or string has fields with certain types of values:
//
//   HaveOnlyFields(json, "id", reflect.String)  // assert that there may a field `id` with a string value.
//   HaveOnlyFields(json, "count", reflect.Float64)  // assert that there may a field `count` with an numeric value.
//   HaveOnlyFields(json, "default", reflect.Interface)  // assert that there may a field `default` with any type of value.
//
func HaveOnlyFields(actual interface{}, allowed ...interface{}) (fail string) {
	usage := "HaveOnlyFields expects parseable JSON to be compared to an fieldPath string, fieldKind reflect.Kind pairs."
	if actual == nil {
		return usage
	}
	json, err := parseJSON(actual)
	if err != nil {
		return err.Error()
	}

	fail = haveOnlyKeys(json, allowed...)
	if fail != "" {
		fail = FormatFailure("JSON has unexpected fields.", fail, "", "")
		return
	}
	fail = haveFields(json, false, allowed...)
	if fail != "" {
		fail = FormatFailure("JSON fields do not match expected type.", fail, "", "")
	}
	return
}

func haveOnlyKeys(json *gabs.Container, allowed ...interface{}) (fail string) {
	children, err := json.ChildrenMap()
	if err != nil {
		return err.Error()
	}
	for key, child := range children {
		found := false
		for _, v := range allowed {
			name, ok := v.(string)
			if !ok { // onlhy pay attention to strings
				continue
			}
			if key == name {
				found = true
				break
			}
		}
		if !found {
			fail += fmt.Sprintf("fields['%s'] (%s) is not allowed. ", key, child)
		}
	}
	return
}

// BeJSON asserts that the first argument can be parsed as JSON.
func BeJSON(actual interface{}, expected ...interface{}) (fail string) {
	usage := "BeJson expects a single string argument and passes if that argument parses as JSON."
	if actual == nil {
		return usage
	}
	_, err := parseJSON(actual)
	if err != nil {
		return err.Error()
	}
	return ""
}

// HaveOnlyCamelcaseKeys passes if all the attributes within a JSON container or
// string contain only upper and lower case ASCII letters.
//
func HaveOnlyCamelcaseKeys(actual interface{}, ignored ...interface{}) (fail string) {
	usage := "HaveOnlyCamelcaseKeys expects parseable JSON in actual. It's keys will recursively checked to make sure they have no snake_case keys. Any right-side arguments should be snake_case key names to be ignored."
	if actual == nil {
		return usage
	}
	json, err := parseJSON(actual)
	if err != nil {
		return err.Error()
	}

	ignoreMap := make(map[string]bool)
	for _, igI := range ignored {
		igS, ok := igI.(string)
		if !ok {
			return fmt.Sprintf("%s. One of the ignored values (%#v) is a %T, not a string.\n%#v", usage, igI, igI, ignored)
		}
		ignoreMap[igS] = true
	}

	fail = checkCamelcaseKeys(json, ignoreMap)
	if fail != "" {
		fail = FormatFailure("JSON field names is not camelCase.", fail, "", "")
	}
	return
}

var camelCaseRegexp = regexp.MustCompile(`^[a-z][a-zA-Z0-9]*$`)

func checkCamelcaseKeys(j *gabs.Container, ignores map[string]bool) (fail string) {
	// if j is an Object with keys, check each of keys and children
	children, notObjectErr := j.ChildrenMap()
	if notObjectErr == nil {
		for k, v := range children {
			if ignores[k] {
				return
			}
			if camelCaseRegexp.MatchString(k) {
				fail = checkCamelcaseKeys(v, ignores)
			} else {
				fail = fmt.Sprintf("Expecting only camelCase keys: found '%s'.\nWithin: %s ",
					k, j)
			}
			if fail != "" {
				return // fail fast to limit recursion
			}
		}
		return
	}

	// If j is an array, check each of it's elements
	// j.Children() will succeed for objects but lose the keys, so order is
	// it's important to check this AFTER the j.ChildMap() check
	elements, notArrayErr := j.Children()
	if notArrayErr == nil {
		for _, element := range elements {
			fail = checkCamelcaseKeys(element, ignores)
			if len(fail) > 0 {
				return // fail fast to limit recursion
			}
		}
		return
	}

	// otherwise this is an atomic object. No check necessary.
	return ""
}

func argsForBesortedByField(
	actual interface{}, args ...interface{}) (
	parsedActual StructureExplorer,
	elementPath, fieldName string, isDescending bool,
	fail string,
) {
	var (
		err error
		ok  bool
	)
	usage := "BeSortedByField expects parseable JSON and at least two right-side args. \nThe JSON is in the left-side arg. It extracts an item arg[0] from the json and checks that it is sorted on field arg[1], ascending by default, descending if arg[2] == true."
	if actual == nil {
		fail = FormatFailure(usage, "", "", "")
		return
	}
	// grab the JSON in actual
	if parsedActual, err = ParseJSON(actual); err != nil {
		fail = FormatFailure("Error Parsing JSON", fmt.Sprintf("%s\n%s", usage, err.Error()), "", "")
		return
	}
	// grab the elementPath and field Name
	if len(args) < 2 {
		fail = FormatFailure(fmt.Sprintf("Expecting at least 2 right-side arguments, got %d", len(args)), usage, "", "")
	}
	if elementPath, ok = args[0].(string); !ok {
		msg := fmt.Sprintf("Expected elementPath to be a string but got %#v instead", args[0])
		fail = FormatFailure(msg, usage, "", "")
		return
	}
	if fieldName, ok = args[1].(string); !ok {
		msg := fmt.Sprintf("Expected fieldName to be a string but got %#v instead", args[1])
		fail = FormatFailure(msg, usage, "", "")
		return
	}
	// grab optional direction param
	if len(args) >= 3 {
		if isDescending, ok = args[2].(bool); !ok {
			msg := fmt.Sprintf("Expected isDecending, if present, to be a bool but got %#v instead", args[2])
			fail = FormatFailure(msg, usage, "", "")
			return
		}
	}
	return
}

// BeSortedByField passes if actual parses to JSON and has an element named
// args[0] in which every element has a field named arg[1] which is sorted,
// ascending by default, if arg[2] is true, the sort is descending.
func BeSortedByField(actual interface{}, args ...interface{}) (fail string) {
	json, path, field, isDescending, fail := argsForBesortedByField(actual, args...)
	if fail != "" {
		return
	}
	_, _, _, _ = json, path, field, isDescending
	data, ok := json.GetPathCheck(path)
	if !(ok && data.IsArray()) {
		fail = FormatFailure(fmt.Sprintf("Actual.%s should be an array, but it's not.", path), "", "", "")
		return
	}
	var (
		n           = data.Len()
		direction   = ">="
		assertion   = BeGreaterThanOrEqualTo
		cur, prev   string
		initialized bool
	)
	if isDescending {
		direction = "<="
		assertion = BeLessThanOrEqualTo
	}

	getFieldFrom := func(i int) string {
		record := data.GetElement(i)
		dataNode, ok := record.GetPathCheck(field)
		if !ok {
			fail = FormatFailure(fmt.Sprintf("Couldn't find field %s in item %d of sorted array", field, i), record.String(), "", "")
		}
		data, ok := dataNode.Data().(string)
		if !ok {
			short := fmt.Sprintf("Expected items[%d].%s to be a string, but it was a %T: %#v", i, field, dataNode.Data(), dataNode.Data())
			long := fmt.Sprintf("items[%d] = %s\nraw(items[%d].%s)=%s", i, record, i, field, dataNode)
			fail = FormatFailure(short, long, "", "")
		}
		return data
	}

	for i := 0; i < n; i++ {

		cur = getFieldFrom(i)
		if fail != "" {
			return
		}
		if !initialized {
			initialized = true
			prev = cur
			continue
		}
		fail = assertion(cur, prev)
		if fail != "" {
			msg := fmt.Sprintf("Sorted array a[%d].%s=%s is not %s item a[%d],%s=%s", i, field, cur, direction, i-1, field, prev)
			long := fmt.Sprintf("a[%d] = %s\nWHICH SHOULD HAVE .%s %s...\na[%d] = %s", i, data.GetElement(i), field, direction, i-1, data.GetElement(i-1))
			fail = FormatFailure(msg, long, "", "")
		}
		prev = cur
	}

	return
}

func argsForCountTests(
	actual interface{}, args ...interface{}) (
	parsedActual StructureExplorer,
	elementPath string,
	count int,
	fail string,
) {
	var (
		err error
		ok  bool
	)

	usage := `CountAtLeast expects parseable JSON and at  two right-side args. The
	JSON is in the left-side arg. It extracts an item arg[0] from the json and
	checks that it is an array field with at least arg[1] items.`

	if actual == nil {
		fail = FormatFailure(usage, "", "", "")
		return
	}
	// grab the JSON in actual
	if parsedActual, err = ParseJSON(actual); err != nil {
		fail = FormatFailure("Error Parsing JSON", fmt.Sprintf("%s\n%s", usage, err.Error()), "", "")
		return
	}
	// grab the elementPath
	if len(args) < 2 {
		fail = FormatFailure(fmt.Sprintf("Expecting  2 right-side arguments, got %d", len(args)), usage, "", "")
	}
	if elementPath, ok = args[0].(string); !ok {
		msg := fmt.Sprintf("Expected elementPath to be a string but got %#v instead", args[0])
		fail = FormatFailure(msg, usage, "", "")
		return
	}
	if count, ok = args[1].(int); !ok {
		msg := fmt.Sprintf("Expected count to be an integer but got %#v instead", args[1])
		fail = FormatFailure(msg, usage, "", "")
		return
	}
	return
}

// CountAtLeast passes if actual parses to a JSON and has an element named
// args[0] which is an array with >= arg[1] items
func CountAtLeast(actual interface{}, args ...interface{}) (fail string) {
	json, path, count, fail := argsForCountTests(actual, args...)
	if fail != "" {
		return
	}
	data, ok := json.GetPathCheck(path)
	if !(ok && data.IsArray()) {
		fail = FormatFailure(fmt.Sprintf("Actual.%s should be an array, but it's not.", path), "", "", "")
		return
	}
	if data.Len() < count {
		fail = FormatFailure(
			fmt.Sprintf("Expected at least %d items in Actual.%s, but found %d.",
				count, path, data.Len()),
			data.String(), "", "",
		)
	}
	return
}
