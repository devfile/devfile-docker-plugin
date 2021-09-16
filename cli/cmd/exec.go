package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/compose-spec/compose-go/types"
	"github.com/containerd/console"
	"github.com/docker/cli/cli"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/spf13/cobra"
)

type execOpts struct {
	service     string
	command     []string
	environment []string
	workingDir  string

	noTty      bool
	user       string
	detach     bool
	index      int
	privileged bool
}

func ExecCommand(backend api.Service) *cobra.Command {
	opts := execOpts{
		index: 1,
	}
	runCmd := &cobra.Command{
		Use:   "exec [options] [-e KEY=VAL...] [--] SERVICE COMMAND [ARGS...]",
		Short: "Execute a command in a running container.",
		Args:  cobra.MinimumNArgs(2),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.service = args[0]
			opts.command = args[1:]
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			project, err := convertToProject(cmd.Flag("devfile").Value.String())
			if err != nil {
				return err
			}

			return runExec(cmd.Context(), backend, project, opts)
		},
	}

	runCmd.Flags().BoolVarP(&opts.detach, "detach", "d", false, "Detached mode: Run command in the background.")
	runCmd.Flags().StringArrayVarP(&opts.environment, "env", "e", []string{}, "Set environment variables")
	runCmd.Flags().BoolVarP(&opts.privileged, "privileged", "", false, "Give extended privileges to the process.")
	runCmd.Flags().StringVarP(&opts.user, "user", "u", "", "Run the command as this user.")
	runCmd.Flags().BoolVarP(&opts.noTty, "no-TTY", "T", false, "Disable pseudo-TTY allocation. By default `exec` allocates a TTY.")
	runCmd.Flags().StringVarP(&opts.workingDir, "workdir", "w", "", "Path to workdir directory for this command.")

	runCmd.Flags().SetInterspersed(false)
	return runCmd
}

func runExec(ctx context.Context, backend api.Service, project *types.Project, opts execOpts) error {
	execOpts := api.RunOptions{
		Service:     opts.service,
		Command:     opts.command,
		Environment: opts.environment,
		Tty:         !opts.noTty,
		User:        opts.user,
		Privileged:  opts.privileged,
		Index:       opts.index,
		Detach:      opts.detach,
		WorkingDir:  opts.workingDir,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if execOpts.Tty {
		con := console.Current()
		if err := con.SetRaw(); err != nil {
			return err
		}
		defer func() {
			if err := con.Reset(); err != nil {
				fmt.Println("Unable to close the console")
			}
		}()

		execOpts.Stdin = con
		execOpts.Stdout = con
		execOpts.Stderr = con
	}
	exitCode, err := backend.Exec(ctx, project, execOpts)
	if exitCode != 0 {
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		}
		return cli.StatusError{StatusCode: exitCode, Status: errMsg}
	}
	return err
}
