package main

import (
	"github.com/ryo-arima/visualizer/pkg/client"
	"github.com/ryo-arima/visualizer/pkg/config"
)

func main() {
	conf := config.NewBaseConfig()
	client.ClientForAdminUser(conf)
}