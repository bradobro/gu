package gu

var (
	// DefaultAsserter allows for simple assertions without creating an asserter.
	DefaultAsserter = NewAsserter(false, 10, VerbosityDebug)
	// Assert uses the DefaultAsserter to make assertions.
	Assert = DefaultAsserter.Assert
	// AssertSkip uses the DefaultAsserter to make assertions, optionally skipping frames in stack traces.
	AssertSkip = DefaultAsserter.AssertSkip
)
