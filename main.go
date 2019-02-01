package main

import (
	"github.com/EDDYCJY/edge-pprof/pkg/setting"
	"github.com/EDDYCJY/edge-pprof/router"
	"github.com/EDDYCJY/edge-pprof/server"
)

func init() {
	setting.NewSetting()
	server.NewPProf()
}

func main() {
	r := router.InitRouter()
	r.Run(setting.ServerSetting.HttpAddr + ":" + setting.ServerSetting.HttpPort)
}
