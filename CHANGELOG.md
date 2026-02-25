# Changelog

All notable changes to this project will be documented in this file.

## [v1.2.0] - 2026-02-25

### Added
- Unit and integration tests with Docker-based test runner
- `EventPath()` to `VirtualKeyboard`, `VirtualMouse`, and `VirtualTouchpad` interfaces
- `EventPath()` to `VirtualDevice` and `VirtualGamepad` interfaces
- `Send()` method to all virtual device classes
- `.gitignore` entries for `.DS_Store` and compiled binaries
- Doc comments to all exported symbols

### Fixed
- Makefile `shell` target now uses `docker-entrypoint.sh` to mount devtmpfs
- `WaitForEventFile` timeout increased from 500ms to 2s for reliability
- Race condition: panic when `Send()` is called on an unregistered device
- Touchpad `fingerCount` going negative causing index out of bounds panic
- Gamepad `Press`/`Release` not returning early on unassigned button
- Repeat events incorrectly using `EV_MSC` instead of `EV_REP`
- `FFEffect` struct to use byte array matching C union size
- AZERTY keymap conflicts for `]` and `ù` characters
- Mangled `SwitchEvent` constant names missing space before type
- `writeEvent` error output to use stderr with newline
- Various documentation fixes (typos, incorrect types, missing calls)

### Changed
- Align Dockerfile `GO_VERSION` with go.mod (1.21)
- Rename `GenricMouse.go` to `GenericMouse.go`
- Remove unused `IsDenormalized` field from `AbsAxis`
- Translate French comment to English in `atomicBool.go`
- Simplify range loop by removing unused blank identifier

## [v1.1.0] - 2025-01-08

### Added
- MSC generic support
- JoyCon IMU virtual device
- Nintendo Switch JoyCon L & R virtual devices
- Xbox One Elite 2 controller
- `TapKey` and missing `SyncReport` to `VirtualKeyboard`
- Configurable tap duration for keyboard typing
- `Send()` and `EventPath()` to `VirtualDevice` interface
- Generic `SendMiscEvent` replacing specialized `SendScanCode`

### Changed
- **Breaking:** Replace `SendScanCode` with generic `SendMiscEvent`

## [v1.0.0] - 2024-12-24

### Added
- Initial release
- Virtual device abstraction over Linux uinput
- Virtual keyboard with typing and keymap support
- Virtual mouse with click, scroll, and move support
- Virtual touchpad with multitouch protocol A and B support
- Virtual gamepad with digital buttons and analog sticks
- Predefined controller profiles (Xbox 360, Xbox One S, Stadia, and more)
- Linux input event codes and uinput bindings