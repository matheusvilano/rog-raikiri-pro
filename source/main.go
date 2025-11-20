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
const version = "v0.2"

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

	// User MUST have sudo/admin privileges, or the command will [likely] fail.

	switch args[0] {

	case "add":
		if err := add(); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

	case "remove":
		if err := remove(); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

	case "status":
		if err := status(); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

	default: // Show help if an unknown command is entered
		help()
	}
}
