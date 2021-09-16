package cmd

import (
	"fmt"
	"testing"
)

func TestDownCommand(t *testing.T) {
	t.Run("Works", func(t *testing.T) {
		mock := &apiMock{}
		cmd := DownCommand(mock)
		cmd.Flags().String("devfile", "test-data/devfile.yaml", "")
		err := cmd.Execute()
		if err != nil {
			t.Errorf("DownCommand() error = %v", err)
			return
		}
		if mock.downProjectName != "golang" {
			t.Errorf("DownCommand() has wrong project name, wanted %q, got %q", "golang", mock.downProjectName)
		}
	})

	t.Run("Handles errors", func(t *testing.T) {
		expectedErr := fmt.Errorf("some error")
		mock := &apiMock{downError: expectedErr}
		cmd := DownCommand(mock)
		cmd.Flags().String("devfile", "test-data/devfile.yaml", "")
		err := cmd.Execute()
		if err != expectedErr {
			t.Errorf("DownCommand() expected error = %v, got %v", expectedErr, err)
			return
		}
	})
}
