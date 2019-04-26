package router

import (
	"net/http"

	"github.com/asynccnu/mana_service_v2/handler/apartment"
	"github.com/asynccnu/mana_service_v2/handler/ios/banner"
	"github.com/asynccnu/mana_service_v2/handler/ios/config"
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

	// The user handlers, requiring authentication
	u := g.Group("/api")
	// u.Use(middleware.AuthMiddleware())
	{
		u.GET("/apartment", apartment.Get)
		u.GET("/site", site.Get)
		u.GET("/ios/config", config.Get)
		u.PUT("ios/config", config.Update)
		u.GET("/ios/banner", banner.Get)
		u.PUT("ios/banner", banner.Update)
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
