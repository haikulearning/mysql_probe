# Change Log
All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

Because this project has not reached 1.0.0 yet anything may change at any time. The public API should not be considered stable. Still, all breaking changes to the public API will be documented in this file.

## [Unreleased][unreleased]
### Breaking Changes
* `connection_count_lte_*` report file will no longer be created. Use `threads_connected_count_lte_*` report files instead (they're functionally equivalent).

### Added
* This Change Log!
* Initial support for cross-compiling via make

## [0.1.0] - 2015-08-27
### Breaking Changes
* Use `mysql_probe start` instead of just `mysql_probe`

### Added
* Use `mysql_probe test` or `mysql_probe serve` to run only the test or status server functionality, respectively.
* Status server supports `--reports` and `--server_port` flags.
* A very basic Makefile

## [0.0.3] - 2015-08-20
### Added
* A basic status HTTP server
* Tests time out if MySQL is slow to respond

### Changed
* Moved the testing logic of `mysql_probe` into a background Goroutine.

## [0.0.2] - 2014-12-18
### Added
* JSON logs
* Reports directory

## [0.0.1] - 2014-12-18
### Added
* "Everything!" This was the initial release.


[unreleased]: https://github.com/haikulearning/mysql_probe/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/haikulearning/mysql_probe/compare/v0.0.3...v0.1.0
[0.0.3]: https://github.com/haikulearning/mysql_probe/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/haikulearning/mysql_probe/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/haikulearning/mysql_probe/compare/0c36901b85f8e...v0.0.1