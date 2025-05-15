package cmd

import (
	"os"
	"os/exec"
)

// runCommand executes a shell command in the current directory
func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// buildVite builds the Vite assets
func buildVite() error {
	if err := runCommand("pnpm", "vite", "build"); err != nil {
		return err
	}
	return runCommand("pnpm", "vite", "build", "--ssr")
}

// buildGo builds the Go code
func buildGo() error {
	return runCommand("go", "build", "-o", "./build/main")
}

// buildGoTemp builds the Go code temporarily
func buildGoTemp() error {
	// Ensure tmp directory exists
	if err := os.MkdirAll("./tmp", 0755); err != nil {
		return err
	}
	return runCommand("go", "build", "-o", "./tmp/main")
}
