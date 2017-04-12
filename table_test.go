package gocu

import "testing"

func TestStringTable(t *testing.T) {

	var tbl Table
	tbl = &SimpleTable{
		{"a", "b", "c"},
		{"a", "b", "c"},
	}
	_ = tbl.Get(0, "a")

}
