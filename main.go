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

	err:=config.Init()
	if err!=nil{
		logger.Errorf("Config initialization error: %v",err)
		return
		//panic(err)
	}
	router.Initialize()
}