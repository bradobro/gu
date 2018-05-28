# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [2.0.0] - 2018-05-28
### Changed
- BREAKING: Assertion functions now must return errors, not strings.
- Assertion functions can have any number of typed params.
- Assertion functions may have  or omit typed or untyped variadic params.

## [1.0.3] - 2018-05-20
### Added
- DefaultAsserter exposing a package-level Assert() and AssertSkip()
### Changed
- Changed project name from `au` to `gu`.

## [1.0.0] - 2018-05-20
### Added
- simpler approach to GoConvey-style assertions
### Removed
- generator approach
