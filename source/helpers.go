// Copyright 2025 Matheus Vilano
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// The name of the udev rules file containing instructions for adding support for the ROG Raikiri Pro gamepad.
const rulesFileName = "99-xpad-raikiri-pro.rules"

// The directory where the udev rules file is located.
const rulesSourceDir = "./resources/"

// The directory where the udev rules file will be copied to.
const rulesTargetDir = "/etc/udev/rules.d/"

// Show help and exit.
func help() {
	fmt.Println("Usage:")
	fmt.Println("  ", executableName, "         // Display instructions")
	fmt.Println("  ", executableName, "add      // Add support for the ROG Raikiri Pro gamepad")
	fmt.Println("  ", executableName, "remove   // Remove support for the ROG Raikiri Pro gamepad")
	fmt.Println("  ", executableName, "status   // Check if support for the ROG Raikiri Pro gamepad is set up")
	fmt.Println("")
	fmt.Println("Notes:")
	fmt.Println("   Support is added/removed by creating/deleting a", rulesFileName, "file in", rulesTargetDir)
	fmt.Println("   You might need to run this command with sudo/admin privileges")
	fmt.Println("   You will likely need to restart your computer for the changes to take effect")
}

// Add support for the Raikiri Pro gamepad by copying the "99-xpad-raikiri-pro.rules" file to "/etc/udev/rules.d/".
func add() error {

	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	source := filepath.Join(workingDir, rulesSourceDir, rulesFileName)
	destination := filepath.Join(rulesTargetDir, rulesFileName)

	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("could not open source file: %w", err)
	}
	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {
			panic(err)
		}
	}(sourceFile)

	destinationFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("could not create destination file: %w", err)
	}
	defer func(destinationFile *os.File) {
		err := destinationFile.Close()
		if err != nil {

		}
	}(destinationFile)

	_, err = exec.Command("cp", source, destination).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	fmt.Println("Successfully [re]added support for the ROG Raikiri Pro gamepad.")
	fmt.Println("You will likely need to restart your computer for the changes to take effect.")
	return nil
}

// Remove support for the Raikiri Pro gamepad by deleting the "99-xpad-raikiri-pro.rules" file at "/etc/udev/rules.d/".
func remove() error {

	filePath := filepath.Join(rulesTargetDir, rulesFileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Nothing to do: 'xpad-raikiri-pro add' was never run before.")
		return nil
	}

	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to remove the file: %w", err)
	}

	fmt.Println("Successfully removed support for the ROG Raikiri Pro gamepad.")
	fmt.Println("You will likely need to restart your computer for the changes to take effect.")
	return nil
}
