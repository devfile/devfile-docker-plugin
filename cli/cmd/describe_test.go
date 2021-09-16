package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/docker/compose/v2/pkg/api"
)

func TestDescribeCommand(t *testing.T) {
	t.Run("Works with text output", func(t *testing.T) {
		marker := time.Now().Format(time.RFC3339)
		mock := &apiMock{
			psContainerSummary: []api.ContainerSummary{
				{Service: "golang", State: marker},
			},
		}
		cmd := DescribeCommand(mock)
		cmd.Flags().String("devfile", "test-data/devfile.yaml", "")
		cmd.Flags().String("output", "text", "")
		out := new(bytes.Buffer)
		cmd.SetOut(out)
		err := cmd.Execute()
		if err != nil {
			t.Errorf("DescribeCommand() error = %v", err)
			return
		}
		if !strings.Contains(out.String(), marker) {
			t.Errorf("DescribeCommand() has output, wanted %q, got %q", marker, out.String())
		}
	})

	t.Run("Works with json output", func(t *testing.T) {
		marker := time.Now().Format(time.RFC3339)
		mock := &apiMock{
			psContainerSummary: []api.ContainerSummary{
				{Service: "golang", State: marker},
			},
		}
		cmd := DescribeCommand(mock)
		cmd.Flags().String("devfile", "test-data/devfile.yaml", "")
		cmd.Flags().String("output", "json", "")
		out := new(bytes.Buffer)
		cmd.SetOut(out)
		err := cmd.Execute()
		if err != nil {
			t.Errorf("DescribeCommand() error = %v", err)
			return
		}
		var data interface{}
		output := out.Bytes()
		err = json.Unmarshal(output, &data)
		if err != nil {
			t.Errorf("DescribeCommand() invalid json (%v): %q", string(output), err)
			return
		}
		if !strings.Contains(fmt.Sprint(data), marker) {
			t.Errorf("DescribeCommand() has output, wanted %q, got %q", marker, out.String())
		}
	})

	t.Run("Handles errors", func(t *testing.T) {
		expectedErr := fmt.Errorf("some error")
		mock := &apiMock{psError: expectedErr}
		cmd := DescribeCommand(mock)
		cmd.Flags().String("devfile", "test-data/devfile.yaml", "")
		cmd.Flags().String("output", "json", "")
		err := cmd.Execute()
		if err != expectedErr {
			t.Errorf("DescribeCommand() expected error = %v, got %v", expectedErr, err)
			return
		}
	})
}
