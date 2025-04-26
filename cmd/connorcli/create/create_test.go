package create_test

import (
	"bytes"
	"connorcli/cmd/connorcli/create"
	del "connorcli/cmd/connorcli/delete"
	"testing"
)

func TestRenameCommand(t *testing.T) {
	// Create create command
	cmd := create.CreateCommand()

	// Redirect stdout
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	// Set Args
	args := []string{"test"}
	cmd.SetArgs(args)

	// Catch errs and execute

	err := cmd.Execute()

	if err != nil {
		t.Error(err)
	}

	// Validate output

	expectedOutput := "test is successfully created!\n"

	if expectedOutput != stdout.String() {
		t.Errorf("Expected output: %q, but received: %q", expectedOutput, stdout.String())
	}

	// Cleanup files

	cleanCmd := del.DeleteCommand()

	cleanCmd.SetArgs(args)

	err = cleanCmd.Execute()

	if err != nil {
		t.Error("An error occured during cleanup ", err)
	}

}