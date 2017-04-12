package should

// Exports JSON (and potentially other data structures or mocks) as an interface

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

// StructureExplorer considers generalizing *gabs.Container with the methods
// needed to test a complex data structure's content and schema.
type StructureExplorer interface {
	// String prettyprints the structure.
	String() string

	// Data gets the datum stored within a StructureExplorer
	Data() interface{}

	// IsArray returns true if the Structure has ordered values
	IsArray() bool
	// Len returns the number of items in an array structure
	Len() int
	// GetElement returns the the i-th element of an array structure (structure[i])
	GetElement(i int) StructureExplorer

	// IsObject returns true if the Structure explorer holds (possibly) unordered named  values
	IsObject() bool
	// Has returns true if the key names an attribute in a structure
	Keys() []string
	// PathExists returns true if the structure has an element at path with a non-null value
	PathExists(path string) bool
	// GetPath returns the the element from an object structure by name if path exists and sets ok to true
	GetPath(path string) StructureExplorer
	// GetPathCheck returns (element, true) if path exists it's element has a
	// non-null value. Otherwise it returns (undefined, false.)
	GetPathCheck(path string) (result StructureExplorer, ok bool)
}

// func StructureHasKey(key string) bool{
// 	for _, k
// }

// GabsExplorer wraps a StructureExplorer over a gabs.Container
type GabsExplorer gabs.Container

// Data gets the datum stored within a StructureExplorer
func (ge *GabsExplorer) Data() interface{} {
	g := (*gabs.Container)(ge)
	return g.Data()
}

// IsArray returns true if the Structure has ordered values
func (ge *GabsExplorer) IsArray() bool {
	g := (*gabs.Container)(ge)
	_, ok := g.Data().([]interface{})
	return ok
}

// Len returns the number of items in an array structure
func (ge *GabsExplorer) Len() int {
	g := (*gabs.Container)(ge)
	result, ok := g.Data().([]interface{})
	if !ok {
		return -1
	}
	return len(result)
}

// GetElement returns the the i-th element of an array structure (structure[i])
func (ge *GabsExplorer) GetElement(i int) StructureExplorer {
	g := (*gabs.Container)(ge)
	result := g.Index(i)
	return (*GabsExplorer)(result)
}

// IsObject returns true if the Structure explorer holds (possibly) unordered named  values
func (ge *GabsExplorer) IsObject() bool {
	g := (*gabs.Container)(ge)
	_, ok := g.Data().(map[string]interface{})
	return ok
}

// Keys returns the names of the values within a structure
func (ge *GabsExplorer) Keys() (result []string) {
	g := (*gabs.Container)(ge)
	m, err := g.ChildrenMap()
	if err != nil {
		result = append(result, fmt.Sprintf("Error getting structure map: %s\n%#v", err, g))
	} else {
		for k := range m {
			result = append(result, k)
		}
	}
	return
}

// GetPathCheck returns the the element from an object structure by name
func (ge *GabsExplorer) GetPathCheck(path string) (se StructureExplorer, ok bool) {
	se = ge.GetPath(path)
	if se.Data() != nil {
		ok = true
	}
	return
}

// PathExists returns true if the path exists and its element is non-null
// element isn't found, return's .Data() == nil
func (ge *GabsExplorer) PathExists(path string) bool {
	se := ge.GetPath(path)
	return se.Data() != nil
}

// GetPath returns the the element from an object structure by name. If the
// element isn't found, return's .Data() == nil
func (ge *GabsExplorer) GetPath(path string) StructureExplorer {
	g := (*gabs.Container)(ge)
	result := g.Path(path)
	return (*GabsExplorer)(result)
}

// String converts the explorer to a pretty string
func (ge *GabsExplorer) String() string {
	g := (*gabs.Container)(ge)
	return g.StringIndent("", "  ")
}
