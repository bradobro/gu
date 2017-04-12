package gocu

import (
	"log"
	"strings"
)

type Table interface {
	Rows() int
	Fieldnames() []string
	FieldNumber(name string) int
	Cols() int
	AddRow(values ...string)
	Row(i int) []string
	Get(row int, fieldname string) string
	Apply(row int, step string) string
}



// SimpleTable is a slow, stupid, table implementation, but it's easy to make literals of
type SimpleTable [][]string

// Rows returns the number of rows in the table
func (tbl *SimpleTable) Rows() int {
	return len(*tbl) - 1
}

// Fieldnames returns the names of fields
func (tbl *SimpleTable) Fieldnames() []string {
	return (*tbl)[0]
}

// FieldNumber returns the column number of a field name
func (tbl *SimpleTable) FieldNumber(name string) int {
	for i, n := range tbl.Fieldnames() {
		if n == name {
			return i
		}
	}
	log.Fatalf("Field name '%s' not found in table with columns '%s", name, strings.Join(tbl.Fieldnames(), ", "))
	return 0 // won't be reached, but syntax requires
}

// Cols returns the number of columns (fields) in the table
func (tbl *SimpleTable) Cols() int {
	return len(tbl.Fieldnames())
}

// AddRow adds a row of field values to the table
func (tbl *SimpleTable) AddRow(values ...string) {
	if len(values) != tbl.Cols() {
		log.Fatalf("Attempting to add %d values as a row in a %d column table: %s",
			len(values), tbl.Cols(),
			strings.Join(values, ", "))
	}
	*tbl = append(*tbl, values)
}

// Row gets a row of values
func (tbl *SimpleTable) Row(i int) []string {
	return (*tbl)[i+1]
}

// Get retrieves a value from the row
func (tbl *SimpleTable) Get(row int, fieldname string) string {
	return tbl.Row(row)[tbl.FieldNumber(fieldname)]
}

// Apply applies a scenario outline row to a step
func (tbl *SimpleTable) Apply(row int, step string) string {
	// TODO: replace scenario outline params
	return step
}
