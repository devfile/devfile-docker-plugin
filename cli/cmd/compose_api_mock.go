package cmd

import (
	"context"

	"github.com/compose-spec/compose-go/types"
	"github.com/docker/compose/v2/pkg/api"
)

type apiMock struct {
	upProject       *types.Project
	upOptions       api.UpOptions
	upError         error
	downProjectName string
	downOptions     api.DownOptions
	downError       error

	psContainerSummary []api.ContainerSummary
	psProjectName      string
	psOptions          api.PsOptions
	psError            error

	execProject  *types.Project
	execOptions  api.RunOptions
	execExitCode int
	execError    error
}

// Build executes the equivalent to a `compose build`
func (m *apiMock) Build(ctx context.Context, project *types.Project, options api.BuildOptions) error {
	panic("not implemented")
}

// Push executes the equivalent ot a `compose push`
func (m *apiMock) Push(ctx context.Context, project *types.Project, options api.PushOptions) error {
	panic("not implemented")
}

// Pull executes the equivalent of a `compose pull`
func (m *apiMock) Pull(ctx context.Context, project *types.Project, opts api.PullOptions) error {
	panic("not implemented")
}

// Create executes the equivalent to a `compose create`
func (m *apiMock) Create(ctx context.Context, project *types.Project, opts api.CreateOptions) error {
	panic("not implemented")
}

// Start executes the equivalent to a `compose start`
func (m *apiMock) Start(ctx context.Context, project *types.Project, options api.StartOptions) error {
	panic("not implemented")
}

// Restart restarts containers
func (m *apiMock) Restart(ctx context.Context, project *types.Project, options api.RestartOptions) error {
	panic("not implemented")
}

// Stop executes the equivalent to a `compose stop`
func (m *apiMock) Stop(ctx context.Context, project *types.Project, options api.StopOptions) error {
	panic("not implemented")
}

// Up executes the equivalent to a `compose up`
func (m *apiMock) Up(ctx context.Context, project *types.Project, options api.UpOptions) error {
	m.upProject = project
	m.upOptions = options
	return m.upError
}

// Down executes the equivalent to a `compose down`
func (m *apiMock) Down(ctx context.Context, projectName string, options api.DownOptions) error {
	m.downProjectName = projectName
	m.downOptions = options
	return m.downError
}

// Logs executes the equivalent to a `compose logs`
func (m *apiMock) Logs(ctx context.Context, projectName string, consumer api.LogConsumer, options api.LogOptions) error {
	panic("not implemented")
}

// Ps executes the equivalent to a `compose ps`
func (m *apiMock) Ps(ctx context.Context, projectName string, options api.PsOptions) ([]api.ContainerSummary, error) {
	//panic("not implemented")
	m.psProjectName = projectName
	m.psOptions = options

	return m.psContainerSummary, m.psError
}

// List executes the equivalent to a `docker stack ls`
func (m *apiMock) List(ctx context.Context, options api.ListOptions) ([]api.Stack, error) {
	panic("not implemented")
}

// Convert translate compose model into backend's native format
func (m *apiMock) Convert(ctx context.Context, project *types.Project, options api.ConvertOptions) ([]byte, error) {
	panic("not implemented")
}

// Kill executes the equivalent to a `compose kill`
func (m *apiMock) Kill(ctx context.Context, project *types.Project, options api.KillOptions) error {
	panic("not implemented")
}

// RunOneOffContainer creates a service oneoff container and starts its dependencies
func (m *apiMock) RunOneOffContainer(ctx context.Context, project *types.Project, opts api.RunOptions) (int, error) {
	panic("not implemented")
}

// Remove executes the equivalent to a `compose rm`
func (m *apiMock) Remove(ctx context.Context, project *types.Project, options api.RemoveOptions) error {
	panic("not implemented")
}

// Exec executes a command in a running service container
func (m *apiMock) Exec(ctx context.Context, project *types.Project, opts api.RunOptions) (int, error) {
	m.execProject = project
	m.execOptions = opts
	return m.execExitCode, m.execError
}

// Copy copies a file/folder between a service container and the local filesystem
func (m *apiMock) Copy(ctx context.Context, project *types.Project, opts api.CopyOptions) error {
	panic("not implemented")
}

// Pause executes the equivalent to a `compose pause`
func (m *apiMock) Pause(ctx context.Context, project string, options api.PauseOptions) error {
	panic("not implemented")
}

// UnPause executes the equivalent to a `compose unpause`
func (m *apiMock) UnPause(ctx context.Context, project string, options api.PauseOptions) error {
	panic("not implemented")
}

// Top executes the equivalent to a `compose top`
func (m *apiMock) Top(ctx context.Context, projectName string, services []string) ([]api.ContainerProcSummary, error) {
	panic("not implemented")
}

// Events executes the equivalent to a `compose events`
func (m *apiMock) Events(ctx context.Context, project string, options api.EventsOptions) error {
	panic("not implemented")
}

// Port executes the equivalent to a `compose port`
func (m *apiMock) Port(ctx context.Context, project string, service string, port int, options api.PortOptions) (string, int, error) {
	panic("not implemented")
}

// Images executes the equivalent of a `compose images`
func (m *apiMock) Images(ctx context.Context, projectName string, options api.ImagesOptions) ([]api.ImageSummary, error) {
	panic("not implemented")
}
