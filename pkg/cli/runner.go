package cli

import (
	"context"
	"io"

	"github.com/suzuki-shunsuke/matchfile/pkg/constant"
	"github.com/urfave/cli/v2"
)

type Runner struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (runner Runner) Run(ctx context.Context, args ...string) error {
	app := cli.App{
		Name:    "matchfile",
		Usage:   "Check file paths are matched to the condition. https://github.com/suzuki-shunsuke/matchfile",
		Version: constant.Version,
		Commands: []*cli.Command{
			{
				Name:      "run",
				Usage:     "Check file paths are matched to the condition",
				ArgsUsage: "<checked file> <condition file>",
				Action:    runner.action,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "log-level",
						Usage: "log level",
					},
				},
			},
			{
				Name:      "list",
				Usage:     "List file paths which matches to the condition",
				ArgsUsage: "<checked file> <condition file>",
				Action:    runner.listAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "log-level",
						Usage: "log level",
					},
				},
			},
		},
	}

	return app.RunContext(ctx, args)
}
