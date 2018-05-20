# Au: Golden Testing in Go

[ ![Codeship Status for bradobro/au](https://app.codeship.com/projects/9544f7b0-3e90-0136-3d5d-32fde32a4b52/status?branch=master)](https://app.codeship.com/projects/290760) [![Maintainability](https://api.codeclimate.com/v1/badges/ff20026525c9ef26b98f/maintainability)](https://codeclimate.com/github/bradobro/au/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/ff20026525c9ef26b98f/test_coverage)](https://codeclimate.com/github/bradobro/au/test_coverage)

**Au** == *aurum* == gold.

**Au** strives to be:

- Compatible: it relies on Go's testing framework and plays well with parallel testing.
- Legible: it should make for short, easy-to-comprehend assertions.
- Minimal:
    - The main API is about 50 SLOC, with one main function and a helper.
    - The reporting is about 150 SLOC, with options to dial output up or down.
    - The basic 8 assertions and a few helpers to write your own are just over 100 SLOC.
- Extensible: it's easy to write your own assertions *with controllable verbosity*.
- Covered: it's major code paths should be tested in ways that it doesn't hide its own breakage.

**Au** is focussed on unit testing--testing Go modules near their implementation in small chunks.

(If you love BDD (and I do) look for **Au**'s upcoming sister project that does just that.)

# Installation

`go get github.com/bradobro/au`

# Quick Start

```Go
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
    // create some structs
	a := Message{1, "Hello, World!"}
	b := Message{1, "Hello, World!"}

	// test their equality
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
```

# Assertion API

## Basic Assertions

## Writing Your Own Assertions

# Asserter API
