# ROG Raikiri on Linux

This is a simple tool made with the Go programming language. It can force the [xpad](https://github.com/paroj/xpad)
Linux kernel driver to claim and drive a ROG Raikiri gamepad. This is by no means an official project—it has no
affiliation with ASUS or ROG.

## Compatibility

The following gamepads are supported:

- ROG Raikiri Pro (`0b05:1a3c`)
- ROG Raikiri II (`0b05:1c92`)

**This tool was made specifically for [xpad](https://github.com/paroj/xpad)**, and it has **NOT** been tested with any
other drivers.

## Requirements

**All you need is a Linux machine with the [xpad](https://github.com/paroj/xpad) Linux kernel driver installed**. If you
intend to build and/or run the tool yourself, you'll also need [Go](https://golang.org/) installed. This
project was developed and tested with Go 1.25, but other versions will likely still work.

## Usage

You can add support for ROG Raikiri gamepads on Linux in **one of three ways**:

- Run the binary in command-line: `./rog-raikiri-pro` (it will give you instructions). \*
- Run the source code directly: `go run rog-raikiri-pro` (it will give you instructions).
- Manually copy/remove `./resources/99-xpad-raikiri-pro.rules` to/from `/etc/udev/rules.d/`.

\* _Binaries are available in the [releases](https://github.com/matheusvilano/rog-raikiri-pro/releases) section, but
you can also build them yourself with `go build <path>`._

## Contributions

You are welcome to add support for other gamepads and/or drivers (e.g. [xpadneo](https://github.com/atar-axis/xpadneo)
or [xone](https://github.com/medusalix/xone))! Creating new
rules similar to those listed in `./resources/99-xpad-raikiri-pro.rules` to match your gamepad/driver _should_ be
enough.
