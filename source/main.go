// Copyright 2025 Matheus Vilano
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"os"
)

// The name of the executable.
const executableName = "rog-raikiri-pro"

// Version of the program.
const version = "v0.1"

// Entry point.
func main() {

	fmt.Println("xpad-raikiri-pro -", version)
	fmt.Println("")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" || args[0] == "help" {
		help()
		return
	}

	switch args[0] {

	case "add": // Ensure the user has sudo/admin privileges
		if err := add(); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

	case "remove": // Ensure the user has sudo/admin privileges
		if err := remove(); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

	default: // Show help if an unknown command is entered
		help()
	}
}
