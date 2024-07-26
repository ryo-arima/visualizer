package main

import (
	"github.com/ryo-arima/visualizer/pkg/agent"
	"github.com/ryo-arima/visualizer/pkg/config"
)

func main() {
	conf := config.NewBaseConfig()
	agent.AgentForAdminUser(conf)
}
