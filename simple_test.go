package gu_test

import (
	"testing"

	"github.com/bradobro/gu"
)

type Message struct {
	Priority int
	Text     string
}

func TestExample(t *testing.T) {
	// create a couple structures
	a := Message{1, "Hello, World!"}
	b := Message{1, "Hello, World!"}

	// test their equality
	gu.Assert(t, gu.Equal, a, b)

	// easily test multiple equality
	c := b
	d := a
	gu.Assert(t, gu.Equal, a, b, c, d)

	// easily test multiple inequality
	b.Text = "Hello, Venus!"
	c.Text = "Hello, Mars!"
	d.Text = "Hello, Universe!"
	gu.Assert(t, gu.Unequal, a, b, c, d)
}
