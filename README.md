# Faye Client

[![Go Test](https://github.com/beeper/faye-client/actions/workflows/go.yml/badge.svg)](https://github.com/beeper/faye-client/actions/workflows/go.yml)

A [Faye](https://faye.jcoglan.com/) client for Go that supports long-polling
and websockets.

The primary purpose of this library is to eventually be used for the real-time
component of a GroupMe library for a Matrix bridge.

## Credit

This library has taken inspiration from
[autogrow/wray](https://github.com/autogrow/wray) and
[karmanyaahm/wray](https://github.com/karmanyaahm/wray). Both of those projects
are licensed under the MIT License, and this project is licensed under the
Apache 2.0 License.

This project adds support for websocekts and exposes a more idiomatic interface
to consumers.
