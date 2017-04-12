VERSION = $(shell sed -n 's/^ *Version *= *"\(.*\)"/\1/p' < version.go)
COMMIT = ${shell git log --pretty='format:%h' -n 1}
BRANCH = ${shell git rev-parse --abbrev-ref HEAD}
# THIS WON'T WORK WITH A LIBRARY BUILD!
BUILD_VARS = -X ${INJECT_VARS_SITE}.Version=${VERSION} -X ${INJECT_VARS_SITE}.Commit=${COMMIT}

glide: ${GOPATH}/bin/glide
${GOPATH}/bin/glide:
	curl https://glide.sh/get | sh

test:
	go clean
	go test ./ ./should ./debug

cover:
	go clean
	go test -cover ./ ./should

test-watch:
	modd

# Make sure there's no debug code etc.
code-quality:
	@echo "Checking for debugging figments"
	@! grep --exclude Makefile --exclude glide.yaml --exclude-dir vendor -nIR 'y0ssar1an/q' *
	@! egrep --exclude Makefile --exclude-dir vendor -nIR '// *DEBUG' *

# package location for compiled-in values
INJECT_VARS_SITE = gotest

# create libs
build: glide
	glide install
	go build -v -ldflags "${BUILD_VARS}"  $(glide novendor)

# create any distribution files
dist: build

release:
	github-release kindrid/gotest ${VERSION} ${BRANCH} copyTheChangeLogManually CHANGELOG.md

# Convention for our vendored builds on Semaphore
ci-build: build

# Semaphore preliminaries
ci-before: code-quality ci-build

# First semaphore job
ci-job1: code-quality test
