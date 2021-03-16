package main

import (
	"fmt"
	"simpleHttpServer/config"
	"simpleHttpServer/server"
)

func main() {
	config.Load()

	e := server.Engine()

	server.Handler(e)
	s := server.New(e, fmt.Sprintf(":%s", config.Conf.Port))
	go server.SignalProcess(s)
	server.Run(s)
}

