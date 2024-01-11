package main

import (
	

	"github.com/Andy4433/Repositorio_go.git/config"
	"github.com/Andy4433/Repositorio_go.git/router"
)

var(
	logger config.Logger
)
func main() {
	logger = *config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}