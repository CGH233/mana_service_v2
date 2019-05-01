package router

import (
	"net/http"

	"github.com/asynccnu/mana_service_v2/handler/apartment"
	"github.com/asynccnu/mana_service_v2/handler/feedback"
	"github.com/asynccnu/mana_service_v2/handler/ios/banner"
	"github.com/asynccnu/mana_service_v2/handler/ios/config"
	"github.com/asynccnu/mana_service_v2/handler/msg"
	"github.com/asynccnu/mana_service_v2/handler/sd"
	"github.com/asynccnu/mana_service_v2/handler/site"
	"github.com/asynccnu/mana_service_v2/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	public := g.Group("/api")
	{
		public.GET("/apartment", apartment.Get)
		public.GET("/site", site.Get)
		public.GET("/ios/config", config.Get)
		public.GET("/ios/banner", banner.Get)
		public.GET("/msg", msg.Get)
	}

	management := g.Group("/api")
	management.Use(middleware.AuthMiddleware())
	{
		management.PUT("/ios/config", config.Update)
		management.PUT("/ios/banner", banner.Update)
		management.POST("/feedback", feedback.Create)
		management.GET("/feedback", feedback.List)
		management.POST("/msg", msg.Create)
		management.PUT("/msg", msg.Update)
		management.DELETE("/msg", msg.Delete)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
