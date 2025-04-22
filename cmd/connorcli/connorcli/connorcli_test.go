package connorcli_test

import (
	"bytes"
	"connorcli/cmd/connorcli/connorcli"
	"testing"
)

func TestGreetCmd_Execute(t *testing.T) {
	// Create root
	cmd := connorcli.GreetCommand()

	// Redirect stdout to a buffer to capture
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)
	// Set CLI arguments
	cmd.SetArgs([]string{"Connor"})

	// Catch errors and execute Root
	err := cmd.Execute()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if output is correct

	expectedOutput := "Hello, Connor!\n"

	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %q, but received: %q", expectedOutput, stdout.String())
	}
}