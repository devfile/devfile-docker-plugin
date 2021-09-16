package cmd

import (
	"errors"
	"fmt"
	"testing"

	"github.com/compose-spec/compose-go/types"
	. "github.com/devfile/devrunner/tests"
)

func TestUpCommand(t *testing.T) {
	t.Run("Works", func(t *testing.T) {
		mock := &apiMock{}
		cmd := createTestDevEnvCMD(UpCommand(mock), "devfile.yaml")
		cmd.SetArgs([]string{"up"})
		err := cmd.Execute()

		if !ExpectNil(t, err, "no errors") {
			return
		}

		if !ExpectNotNil(t, mock.upProject, "docker compose up was called") {
			return
		}

		ExpectEqual(t, mock.upProject.Name, "golang", "project name")
		ExpectEqual(t, mock.upProject.WorkingDir, ".", "default workdir")
	})

	t.Run("Handles errors", func(t *testing.T) {
		expectedErr := fmt.Errorf("some error")
		mock := &apiMock{upError: expectedErr}
		cmd := createTestDevEnvCMD(UpCommand(mock), "devfile.yaml")
		cmd.SetArgs([]string{"up"})

		err := cmd.Execute()
		if err != expectedErr {
			t.Errorf("UpCommand() expected error = %v, got %v", expectedErr, err)
			return
		}
	})

	t.Run("it uses the workdir passed as parameter", func(t *testing.T) {
		mock := &apiMock{}
		cmd := createTestDevEnvCMD(UpCommand(mock), "devfile.yaml")
		cmd.SetArgs([]string{"up", "-w", "/store/projects/repo1"})
		err := cmd.Execute()

		if !ExpectNil(t, err, "no errors") {
			return
		}

		ExpectEqual(t, mock.upProject.WorkingDir, "/store/projects/repo1", "workdir passed as parameter")
	})

	t.Run("when called with '-env' arguments, it adds the environment variables to all services", func(t *testing.T) {
		mock := &apiMock{}
		cmd := createTestDevEnvCMD(UpCommand(mock), "devfile_two_components.yaml")
		cmd.SetArgs([]string{"up", "-e", "AA=1", "-e", "BB=2"})

		err := cmd.Execute()
		if !ExpectNil(t, err, "no errors") {
			return
		}

		service1 := mock.upProject.Services[0]
		if ExpectNotNil(t, service1.Environment["AA"], "AA env variable") {
			ExpectEqual(t, *service1.Environment["AA"], "1", "AA env variable")
		}

		if ExpectNotNil(t, service1.Environment["BB"], "BB env variable") {
			ExpectEqual(t, *service1.Environment["BB"], "2", "BB env variable")
		}

		service2 := mock.upProject.Services[1]
		if ExpectNotNil(t, service2.Environment["AA"], "AA env variable") {
			ExpectEqual(t, *service2.Environment["AA"], "1", "AA env variable")
		}

		if ExpectNotNil(t, service2.Environment["BB"], "BB env variable") {
			ExpectEqual(t, *service2.Environment["BB"], "2", "BB env variable")
		}
	})

	t.Run("when called with '-volume' arguments, it adds the volumes to all services", func(t *testing.T) {
		mock := &apiMock{}
		cmd := createTestDevEnvCMD(UpCommand(mock), "devfile_two_components.yaml")
		cmd.SetArgs([]string{"up", "-v", "/store:/data", "-v", "/store2:/data2"})

		err := cmd.Execute()
		if !ExpectNil(t, err, "no errors") {
			return
		}

		service1 := mock.upProject.Services[0]
		if !ExpectEqual(t, len(service1.Volumes), 3, "wrong number of volumes") {
			return
		}

		ExpectEqual(t, types.ServiceVolumeConfig{
			Type:   "bind",
			Source: "/store",
			Target: "/data",
		}, service1.Volumes[1], "first volume")
		ExpectEqual(t, types.ServiceVolumeConfig{
			Type:   "bind",
			Source: "/store2",
			Target: "/data2",
		}, service1.Volumes[2], "second volume")
	})

	t.Run("when called with invalid '-volume' arguments, it return an error", func(t *testing.T) {
		mock := &apiMock{}
		cmd := createTestDevEnvCMD(UpCommand(mock), "devfile_two_components.yaml")
		cmd.SetArgs([]string{"up", "-v", "/store:/data", "-v", "AAAA"})
		err := cmd.Execute()
		expectedErr := errors.New("invalid format for 'volume' option 'AAAA'. the format should be host_path:container_path")
		ExpectEqual(t, err, expectedErr, "expected error")
	})
}
