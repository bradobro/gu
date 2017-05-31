
run:
	go run cmd/cyu.go

test:
	go test . ./cmd

test-watch:
	watcher -depth 2 make test