package au_test

import (
	"testing"

	"github.com/bradobro/au"
)

type Message struct {
	Priority int
	Text     string
}

func TestExample(t *testing.T) {
	a := Message{1, "Hello, World!"}
	b := Message{1, "Hello, World!"}

	// Test structure equality
	au.Assert(t, au.Equal, a, b)

	// easily test multiple equality
	c := b
	d := a
	au.Assert(t, au.Equal, a, b, c, d)

	// easily test multiple inequality
	b.Text = "Hello, Venus!"
	c.Text = "Hello, Mars!"
	d.Text = "Hello, Universe!"
	au.Assert(t, au.Unequal, a, b, c, d)
}
