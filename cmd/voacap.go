package main

import (
	"os"
	"voacap/internal/voacap"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "go.uber.org/automaxprocs"
)

func main() {
	hlog.SetLevel(hlog.LevelDebug)

	cmd := voacap.NewCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
