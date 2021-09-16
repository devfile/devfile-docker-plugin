package cmd

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestExecCommand(t *testing.T) {
	t.Run("Works", func(t *testing.T) {
		mock := &apiMock{}
		cmd := ExecCommand(mock)
		cmd.Flags().String("devfile", "test-data/devfile.yaml", "")
		cmd.SetArgs([]string{"golang", "echo", "date"})
		cmd.Flags().Set("no-TTY", "true")
		err := cmd.Execute()
		if err != nil {
			t.Errorf("ExecCommand() error = %v", err)
			return
		}
		if !(mock.execProject != nil && mock.execProject.Name == "golang") {
			t.Errorf("ExecCommand() has wrong project name, wanted %q, got %q", "golang", mock.execProject.Name)
		}

		if mock.execOptions.Service != "golang" {
			t.Errorf("ExecCommand() has wrong service name, wanted %q, got %q", "golang", mock.execOptions.Service)
		}

		if !reflect.DeepEqual(mock.execOptions.Command, []string{"echo", "date"}) {
			t.Errorf("ExecCommand() has wrong command, wanted %q, got %q", []string{"echo", "date"}, mock.execOptions.Command)
		}
	})

	t.Run("Propages exit code", func(t *testing.T) {
		mock := &apiMock{
			execExitCode: 42,
		}
		cmd := ExecCommand(mock)
		cmd.Flags().String("devfile", "test-data/devfile.yaml", "")
		cmd.SetArgs([]string{"golang", "echo", "date"})
		cmd.Flags().Set("no-TTY", "true")
		err := cmd.Execute()
		if err == nil || !strings.Contains(err.Error(), "42") {
			t.Errorf("ExecCommand() should have the expected exit code = %v", err)
			return
		}
	})

	t.Run("Propages errors", func(t *testing.T) {
		expectedErr := fmt.Errorf("some error")
		mock := &apiMock{
			execError: expectedErr,
		}
		cmd := ExecCommand(mock)
		cmd.Flags().String("devfile", "test-data/devfile.yaml", "")
		cmd.SetArgs([]string{"golang", "echo", "date"})
		cmd.Flags().Set("no-TTY", "true")
		err := cmd.Execute()
		if err != expectedErr {
			t.Errorf("ExecCommand() expected error = %v, got %v", expectedErr, err)
			return
		}
	})
}
