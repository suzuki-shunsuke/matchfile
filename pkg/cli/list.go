package cli

import (
	"fmt"

	"github.com/suzuki-shunsuke/matchfile/pkg/controller"
	"github.com/urfave/cli/v2"
)

func (runner Runner) listAction(c *cli.Context) error {
	params, err := runner.setCLIArg(c, controller.Params{})
	if err != nil {
		return fmt.Errorf("parse the command line arguments: %w", err)
	}

	ctrl, params, err := controller.New(c.Context, params)
	if err != nil {
		return fmt.Errorf("initialize a controller: %w", err)
	}

	return ctrl.List(c.Context, params)
}
