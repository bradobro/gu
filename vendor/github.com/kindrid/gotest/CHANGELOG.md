# GoTest Change Log

## Conventions
http://keepachangelog.com/en/0.3.0/

Each version should group changes to describe their impact on the project, as follows:
- **Added**: for new features.
- **Changed**: for changes in existing functionality.
- **Deprecated**: for once-stable features removed in upcoming releases.
- **Removed**: for deprecated features removed in this release.
- **Fixed**: for any bug fixes.
- **Security**: to invite users to upgrade in case of vulnerabilities.

## [0.9.?] unreleased
### Added
- Verbosity control: `--gotest-verbosity`
- Stack trace control: `--gotest-stack`
- Added JSON assertions:
  - `should.BeSortedByField`
  - `should.CountAtLease`
- Added JSON:API assertion: `BeJsonapiError`
### Changed
- `should.MatchHTTPStatusCode` is less verbose
- Dependency updates
- Removed `vendor/`
- Removed all traces of `goconvey`
- Documentation improvements
- Verbosity of Debug (4) or Insane (5) shows information even for successes.


## [0.9.4] 2017-02-21
### Added
- StructureExplorer.GetPathCheck() is like GetPath() but returns a second value, ok bool, to verify whether the value was found.
- StructureExplorer.PathExists() returns true if a path points to a structure element with a non-nil value
- GetVersion() and GetCommit() give info about the code version.
### Changed
- Migrating some json checks to use the `StructureExplorer` interface instead of `*gabs.Container`. (Should probably make `StructureExplorer` its own subpackage.)

## [0.9.3] 2017-02-16
### Added
- Interface `should.StructureExplorer`, a minimal JSON destructuring interface to decouple libraries using `gotest` from `github.com/Jeffail/gabs`.
- Method `should.ParseJSON()` that returns a `should.StructureExplorer` so outside libraries can write their own JSON assertions.
- method `gotest.Later()` to sketch out unimplemented tests

## [0.9.2] 2017-02-09
### Changed
- Updated dependencies with `glide`
- Added `HaveOnlyCamelcaseKeys` to 'BeJSONAPIRecord' (Limitation: BeJSONAPIRecord doesn't have the `ignore` option yet to explicitly allow some snake_case fields)

### Fixed
- camelCase detection regexp now allows numbers but requires initial lowercase letter

## [0.9.1] 2017-02-08
### Added
- `HaveOnlyCamelcaseKeys` assertions makes sure JSON object attributes aren't snake_case.

## [0.9.0] 2017-02-07
### Added
- Vendored dependencies with `glide`
