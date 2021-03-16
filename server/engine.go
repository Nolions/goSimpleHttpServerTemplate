package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simpleHttpServer/config"
)

func Engine() *gin.Engine {
	e := &gin.Engine{}

	if !config.Conf.Debug {
		gin.SetMode(gin.DebugMode)
		e = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		e = gin.New()
	}

	c := cors.DefaultConfig()
	c.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	c.AllowAllOrigins = true
	e.Use(cors.New(c))

	return e
}

func Run(s *http.Server) {
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
