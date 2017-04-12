# GoTest

[![Build Status](https://semaphoreci.com/api/v1/kindrid/gotest/branches/master/shields_badge.svg)](https://semaphoreci.com/kindrid/gotest) [![Issue Count](https://codeclimate.com/github/kindrid/gotest/badges/issue_count.svg)](https://codeclimate.com/github/kindrid/gotest)

Package gotest provides rich assertions for use within and beyond Go's `testing` package.

GoTest plays well with "vernacular" Go `testing` tests, providing a rich set of assertions to test HTTP Responses, JSON and JSON:API data, general equality, numeric comparison, collections, strings, panics, types, and time.

Look at these [GoDocs for the latest documentation](https://godoc.org/github.com/kindrid/gotest).

## Quickstart

Grab the package and import it:

  go get github.com/kindrid/gotest
  import "github.com/kindrid/gotest"
  import "github.com/kindrid/gotest/should"

Code normal `testing`-style tests, but use `gotest.Assert` and assertions found
in `should` like this:

  gotest.Assert(t, actualJson, should.HaveFields,
    "name", reflect.String,
    "children", reflect.Map,
    "hobbies", reflect.Slice)

Assertions are just funcions that accept interfaces and return a non-empty string if there's an error. Look at `gotest.should.Assertion` for the details. Look at `should/doc.go` and `should/assertion.go` ofr more details.

## Overview

GoTest plays well with "vernacular" Go `testing` tests, providing a rich set of
assertions to test HTTP Responses, JSON and JSON:API data, general equality,
numeric comparison, collections, strings, panics, types, and time.

Most of these rich assertions are provided by SmartyStreet's excellent assertion
library (https://github.com/smartystreets/assertions) which builds off Aaron
Jacobs' Oglematchers (https://github.com/jacobsa/oglematchers).

In addition, any SmartyTreets-style assertion can be used as is (see https://github.com/smartystreets/goconvey/wiki/Custom-Assertions).

## Why

We like Go's stdlib `testing` because it's simple, fast, familiar to most Go
coders, has good tooling support, benchmark support, and coverage support.

In earlier versions of Go, we missed subtests for test organization and a more
BDD approach. We looked at GinkGo (`github.com/onsi/ginkgo`) and GoConvey
(github.com/smartystreets/goconvey/convey)--both with benefits--and chose
GoConvey because  of the simple and consistent approach it took to writing
custom assertions. See the documenataion for "gotest/should" for more details on
that. It also had a pretty test runner.

But around the release of Go 1.7 we ran into some problems: the growth of
parameterized (table-driven) tests  in our code didn't play well with GoConvey.
GoConvey's approach to building test suites made it hard to focus specific
subtests.

We also ran into an opportunity: `testing` now supported subtests. See
https://godoc.org/testing#hdr-Subtests_and_Sub_benchmarks. We were able to drop
a lot of fancy suite construction code  and gain easier focussing on particular
tests. But we had come to appreciate GoConvey's excellent  assertion pattern. We
considered moving to `github.com/stretchr/testify` and used it's pattern for the
custom assertion  wrapper (see `Assert` below). It also had a simple
custom-assertion pattern, but GoConvey's seemed simpler and more useful. We took
a sideways glance at GUnit (`github.com/smartystreets/gunit`) and Labix's
GoCheck (`gopkg.in/check.v1 `). Very cool packages, but we wanted to stay closer
to `testing` with its new versatility.
