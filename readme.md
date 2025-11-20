# ROG Raikiri Pro on Linux

This is a simple Go tool to force the [xpad](https://github.com/paroj/xpad) Linux kernel driver to claim the ROG Raikiri
Pro and recognize it as a gamepad. This is by no means an official project—it has no affiliation with ASUS or ROG.

## Requirements

**All you need is a Linux machine with the [xpad](https://github.com/paroj/xpad) Linux kernel driver installed**. If you
intend to build the tool yourself or run it with `go run`, you'll also need [Go](https://golang.org/) installed. This
project was developed and tested with Go 1.25.

## Usage

You can add support for the ROG Raikiri Pro on Linux in three ways:

1) Run it in command-line: `./xpad-raikiri-pro` (it will give you instructions). \*
2) Run the source code directly: `go run xpad-raikiri-pro` (it will give you instructions).
3) Manually copy/remove `./resources/99-rog-raikiri-pro.rules` to/from `/etc/udev/rules.d/`.

\* _Binaries are available in the [releases](https://github.com/paroj/xpad-raikiri-pro/releases) section, but you can
also build them yourself with `go build <path>`._

## Compatibility

**This tool was made specifically for [xpad](https://github.com/paroj/xpad)**, and it has **NOT** been tested with any
other drivers.

You are welcome to adapt it to other drivers, though—and perhaps share it? 👀 There _has_to be someone out there using
the ROG Raikiri Pro
on [xpadneo](https://github.com/atar-axis/xpadneo), [xone](https://github.com/medusalix/xone?tab=readme-ov-file),
or similar. Modifying the `./resources/99-rog-raikiri-pro.rules` file to match your driver _should_ be enough.
