package cli

import (
	"errors"
	"fmt"

	"github.com/suzuki-shunsuke/matchfile/pkg/controller"
	"github.com/urfave/cli/v2"
)

func (runner Runner) setCLIArg(c *cli.Context, params controller.Params) (controller.Params, error) {
	args := c.Args()
	if args.Len() != 2 { //nolint:gomnd
		return controller.Params{}, errors.New(`two arguments are required.
Usage: matchfile <file path to the list of file paths> <file path to the list of conditions>`)
	}
	params.CheckedFilePath = args.First()
	params.ConditionFilePath = args.Get(1)
	if logLevel := c.String("log-level"); logLevel != "" {
		params.LogLevel = logLevel
	}
	return params, nil
}

func (runner Runner) action(c *cli.Context) error {
	params, err := runner.setCLIArg(c, controller.Params{})
	if err != nil {
		return fmt.Errorf("parse the command line arguments: %w", err)
	}

	ctrl, params, err := controller.New(c.Context, params)
	if err != nil {
		return fmt.Errorf("initialize a controller: %w", err)
	}

	return ctrl.Run(c.Context, params)
}
