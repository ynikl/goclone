package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func fetch(command string) string {
	parts := strings.Fields(command)
	head := parts[0]
	parts = parts[1:]

	cmd := exec.Command(head, parts...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("clone failed", err)
	}
	return stdout.String()
}

func execute(command string) error {
	parts := strings.Fields(command)
	head := parts[0]
	parts = parts[1:]

	cmd := exec.Command(head, parts...)
	// print command result to console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
