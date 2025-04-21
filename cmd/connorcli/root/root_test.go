package connorcli_test

import (
	"bytes"
	connorcli "connorcli/cmd/connorcli/root"
	"testing"
)

func TestRootCmd_Execute(t *testing.T) {
	// Create root
	cmd := connorcli.RootCommand()

	// Redirect stdout to a buffer to capture
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	// Catch errors and execute Root
	err := cmd.Execute()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if output is correct

	expectedOutput := "Welcome to Connor CLI!\n"

	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %q, but received: %q", expectedOutput, stdout.String())
	}
}