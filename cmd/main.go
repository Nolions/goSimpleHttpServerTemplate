package main

import (
	"fmt"
	"simpleHttpServer/config"
	"simpleHttpServer/server"
	"simpleHttpServer/store"
	"simpleHttpServer/web"
	"time"
)

func main() {
	config.Load()

	storeConf := store.New(store.Conf{
		MaxOpen:     300,
		MaxIdle:     200,
		MaxLifetime: time.Minute * 5,
	})

	web.Setup(storeConf)

	e := server.Engine()

	server.Handler(e)

	s := server.New(e, fmt.Sprintf(":%s", config.Conf.App.Port))
	go server.SignalProcess(s)
	server.Run(s)
}
