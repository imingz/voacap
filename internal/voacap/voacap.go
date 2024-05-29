package voacap

import (
	"time"
	"voacap/internal/voacap/biz/dal"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func init() {
	dal.Init()
}

func Run() {
	h := server.Default(server.WithHostPorts(":3000"), server.WithExitWaitTime(time.Second))

	register(h)
	h.Spin()
}
