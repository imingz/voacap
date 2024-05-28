package main

import (
	"voacap/cmd"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "go.uber.org/automaxprocs"
)

func main() {
	hlog.SetLevel(hlog.LevelInfo)

	cmd.Execute()
}
