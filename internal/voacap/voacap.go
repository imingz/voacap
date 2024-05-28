package voacap

import (
	"time"
	"voacap/pkg/version"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			version.PrintAndExitIfRequested()
			run()
			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	cmd.Flags().AddFlagSet(pflag.CommandLine)

	return cmd
}

func run() {
	h := server.Default(server.WithHostPorts(":3000"), server.WithExitWaitTime(time.Second))

	register(h)
	h.Spin()
}
