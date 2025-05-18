package size_test

import (
	"bytes"
	"connorcli/cmd/connorcli/size"
	"testing"
)

func TestSizeCmd_Execute(t *testing.T) {
	// Create Command
	cmd := size.SizeCommand()

	// Redirect stdout
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	// Catch Errors
	err := cmd.Execute()

	if err != nil {
		t.Errorf("Unexpected Error: %v", err)
	}

	// Check if output is correct

	expectedOutput := "2086 Bytes\n"
	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %q, but received: %q", expectedOutput, stdout.String())
	}
}
