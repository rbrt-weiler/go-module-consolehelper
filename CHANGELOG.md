# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

(Nothing as of now.)

## [1.0.2] - 2021-02-28

### Changed

1. Update dependencies.

### Fixed

1. Update go.mod to reflect the need for go1.16. This is due to the move away from ioutil.

## [1.0.1] - 2021-02-22

### Changed

1. Added gopls to the default tools used by the VS Code Dev Container.
1. No longer using ioutil-based code due to [deprecation of ioutil](https://golang.org/doc/go1.16#ioutil).
1. The included Dev Container has been upgraded to use [go1.16](https://golang.org/doc/go1.16).
1. The CI configuration has been upgraded to use go1.16 as well.

## [1.0.0] - 2021-01-14

### Added

1. Contributing guide.
1. VS Code Dev Container config.

### Changed

1. New() returns 0 instead of -1 for the dimensions in case the terminal dimensions cannot be determined.
1. Updated documentation.

## [0.4.0] - 2021-01-08

### Added

1. Printf().
1. Print().
1. Println().

## [0.3.0] - 2021-01-07

### Added

1. Sprintln().
1. Fprintf().
1. Fprint().
1. Fprintln().

## [0.2.0] - 2021-01-03

### Added

1. This changelog.
1. Go tests.

### Changed

1. Updated package doc that is more complete than before.
1. Signature of Sprint() matches fmt.Sprint() now.

## [0.1.0] - 2021-01-02

Initial public release.

### Added

1. ConsoleHelper in general.
1. ConsoleHelper.UpdateDimensions().
1. ConsoleHelper.Sprintf().
1. ConsoleHelper.Sprint().

[Unreleased]: https://gitlab.com/rbrt-weiler/go-module-consolehelper/-/compare/v1.0.2...master
[1.0.2]: https://gitlab.com/rbrt-weiler/go-module-consolehelper/-/tree/v1.0.2
[1.0.1]: https://gitlab.com/rbrt-weiler/go-module-consolehelper/-/tree/v1.0.1
[1.0.0]: https://gitlab.com/rbrt-weiler/go-module-consolehelper/-/tree/v1.0.0
[0.4.0]: https://gitlab.com/rbrt-weiler/go-module-consolehelper/-/tree/v0.4.0
[0.3.0]: https://gitlab.com/rbrt-weiler/go-module-consolehelper/-/tree/v0.3.0
[0.2.0]: https://gitlab.com/rbrt-weiler/go-module-consolehelper/-/tree/v0.2.0
[0.1.0]: https://gitlab.com/rbrt-weiler/go-module-consolehelper/-/tree/v0.1.0
