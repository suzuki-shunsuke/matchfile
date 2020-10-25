package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/matchfile/pkg/cli"
	"github.com/suzuki-shunsuke/matchfile/pkg/signal"
)

func main() {
	if err := core(); err != nil {
		logrus.Fatal(err)
	}
}

func core() error {
	runner := cli.Runner{}
	ctx, cancel := context.WithCancel(context.Background())
	go signal.Handle(os.Stderr, cancel)
	return runner.Run(ctx, os.Args...)
}
