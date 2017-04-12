/*Package should gathers helpful tests (a.k.a. matchers, checks, assertions,
/*predicates) for kindrid test libraries.

Design Philosophy

1. The matchers should play well with go's testing library and the `go test`
command, but...

2. The matchers should also be usable without a `*testing.T` for use outside
of Go's test framework.

3. Performance isn't top priority with the matchers, so reflection is okay.

4. We may import, facade, or duplicate code from other libraries within their
licenses as needed to avoid re-implementation. Note code sources in the next
section.

5. All exported functions will imply Should, Must, Expect, etc. We won't
prefix names, but for readability it's best if function names conjugate well
with such a prefix. For instance, `Equal` instead of `ShouldEqual` ().

Patterns

The base pattern we'll use is this:

  satisfiesSomething(actual, expected ...interface{}) (passMessage, failMessage string)

And we'll call that type `Assertion`

1. Accept interface arguments.

2. Return one empty string and one non-empty string to signal success or failure.

3. That string should do terse reporting on one line, providing
longer explanation after a singl `\n`. TODO: consider stack trace and other
debug info after `\n\n\n` or something.

5. Wrappers for other test systems go in the root directory.

Thanks to https://github.com/smartystreets/assertions for their inspiration and
excellent assertion library. Be sure to look at the Variables section (below)
for those assertions aliased into this package.

Future

- [ ] Do we want to pass in options via flags, magic args ("option:save-golden"),
or environment variables (SHOULD_SAVE_GOLDEN=1)?

*/
package should
