# Gu: Go Foundational Testing

[ ![Codeship Status for bradobro/gu](https://app.codeship.com/projects/9544f7b0-3e90-0136-3d5d-32fde32a4b52/status?branch=master)](https://app.codeship.com/projects/290760) [![Maintainability](https://api.codeclimate.com/v1/badges/ff20026525c9ef26b98f/maintainability)](https://codeclimate.com/github/bradobro/gu/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/ff20026525c9ef26b98f/test_coverage)](https://codeclimate.com/github/bradobro/gu/test_coverage)

[GoDoc](https://godoc.org/github.com/bradobro/gu)

**Gu** strives to be:

- Compatible: it relies on Go's testing framework and plays well with parallel testing.
- Legible: it should make for short, easy-to-comprehend assertions.
- Minimal:
    - The main API is about 50 SLOC, with one main function and a helper.
    - The reporting is about 150 SLOC, with options to dial output up or down.
    - The basic 8 assertions and a few helpers to write your own are just over 100 SLOC.
- Extensible: it's easy to write your own assertions *with controllable verbosity*.
- Covered: it's major code paths should be tested in ways that it doesn't hide its own breakage.

**gu** is focussed on unit testing--testing Go modules near their implementation in small chunks.

(If you love BDD (and I do) look for **gu**'s upcoming sister project that does just that.)

# Installation

`go get github.com/bradobro/gu`

# Quick Start

```Go
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
```

# Deeper Start

See the documentation locally at http://localhost:6060:

`godoc -http=:6060`

# Assertion API

## Basic Assertions

## Writing Your Own Assertions

# Asserter API
