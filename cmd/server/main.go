package main

import (
	"github.com/ryo-arima/visualizer/pkg/config"
	"github.com/ryo-arima/visualizer/pkg/server"
)

func main() {
	conf := config.NewBaseConfig()
	server.Main(conf)
}