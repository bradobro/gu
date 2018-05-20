# Cyu: composable unit and bdd testing in Go.

https://github.com/DATA-DOG/godog is official and much more developed, but it takes a very non-go-spirited approach to testing. It had to--`go test` was very limited.

Howeveer, as of Go 1.7, the`testing` library has become a lot richer. Staying closer to the standard libs makes it possible to leverage BDD code with tools for profiling, debugging, and benchmarking.
