package list_test

import (
	"bytes"
	"connorcli/cmd/connorcli/list"
	"testing"
)


func TestListCmd_Execute(t *testing.T) {
	// Create Command
	cmd := list.ListCommand()

	// Redirect stdout
	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	// Catch Errors
	err := cmd.Execute()

	if err != nil {
		t.Errorf("Unexpected Error: %v", err)
	}
}