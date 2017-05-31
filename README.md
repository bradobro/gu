# cuketest-go

Cucumber-style test generation for go (Golang).

https://github.com/DATA-DOG/godog is official and much more developed, but it takes a very non-go-spirited approach to testing. It had to--`go test` was very limited.

Howeveer, as of Go 1.7, the`testing` library has become a lot richer. Staying closer to the standard libs makes it possible to leverage BDD code with tools for profiling, debugging, and benchmarking.

cyu is an experiment in translating Gherkin to Go in a maintainable fashion. It's inspired by (and uses the same Gherkin parser as) [Godog](https://github.com/DATA-DOG/godog) and [Hiptest.net](https://hiptest.net/).
