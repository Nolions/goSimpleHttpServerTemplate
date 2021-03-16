package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func New(r *gin.Engine, addr string) *http.Server {
	log.Printf("Listening on %s", addr)
	return &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func SignalProcess(srv *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	s:= <-c
	switch s {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		log.Printf("signal is %s", s)
		srv.Close()
		return
	case syscall.SIGHUP:
	default:
		return
	}
}

func Handler(router *gin.Engine) {
	router.GET("/", indexHandler)
}

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}
