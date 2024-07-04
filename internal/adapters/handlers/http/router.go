package http

import (
	"github.com/arasan1289/shell-test/internal/adapters/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new Router instance
func NewRouter(config *config.Container, visitorHandler VisitorHandler) (*Router, error) {
	// Disable debug mode in production
	if config.App.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowOrigins = config.HTTP.AllowedOrigins

	router := gin.New()

	// Middlewares
	router.Use(gin.Recovery(), cors.New(ginConfig))

	v1 := router.Group("/v1")
	{
		v1.GET("/visitors-count", visitorHandler.GetVisitorCount)
		v1.POST("/visitor", visitorHandler.NewVisitor)
	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
