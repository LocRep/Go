package server

import (
	"github.com/gin-gonic/gin"
	"github.com/locrep/locrep-go/config"
	"github.com/locrep/locrep-go/maven"
)

func mavenMiddleware(config config.Conf) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		mvnHandler := maven.NewHandler(config)
		mvnHandler.Handle(ctx)
		ctx.Next()
	}
}

func NewServer(config config.Conf) *gin.Engine {
	gin.SetMode(config.Environment.BuildMode())

	server := gin.New()
	server.Use(mavenMiddleware(config))

	return server
}