package rename_test

import (
	"bytes"
	"connorcli/cmd/connorcli/rename"
	"testing"
)

func TestRenameCommand(t *testing.T) {
	// Create rename command
	cmd := rename.RenameCommand()

	// Redirect stdout
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	// Set Args
	args := []string{"test", "test2"}
	cmd.SetArgs(args)

	// Catch errs and execute

	err := cmd.Execute()

	if err != nil {
		t.Error(err)
	}

	// Validate output

	expectedOutput := "Successfully renamed: test\n"

	if expectedOutput != stdout.String() {
		t.Errorf("Expected output: %q, but received: %q", expectedOutput, stdout.String())
	}

	// Set Args
	args = []string{"test2", "test"}
	cmd.SetArgs(args)
	
	// Catch errs and execute
	
	err = cmd.Execute()
	
	if err != nil {
		t.Error(err)
	}

}