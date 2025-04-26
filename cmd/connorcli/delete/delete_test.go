// Implement test for this later

package delete_test

import (
	"bytes"
	"connorcli/cmd/connorcli/create"
	del "connorcli/cmd/connorcli/delete"
	"testing"
)

func TestRenameCommand(t *testing.T) {
	// Create a file to be deleted
	setupCmd := create.CreateCommand()


	// Set Args for new file
	args := []string{"test"}
	setupCmd.SetArgs(args)

	// Create the new file
	err := setupCmd.Execute()

	if err != nil {
		t.Error("An error occured during setup ", err)
	}

	// Create Cmd for testing
	cmd := del.DeleteCommand()

	// Redirect stdout
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	// Setup args
	cmd.SetArgs(args)

	// Catch errs and execute

	err = cmd.Execute()

	if err != nil {
		t.Error(err)
	}

	// Validate output

	expectedOutput := "test is successfully deleted!\n"

	if expectedOutput != stdout.String() {
		t.Errorf("Expected output: %q, but received: %q", expectedOutput, stdout.String())
	}
}