// Package main is the entry point for the Tend API server.
//
//	@title			Tend API
//	@version		0.1.0
//	@description	REST API for the Tend planning app.
//	@BasePath		/
package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", healthHandler)

	return r
}

// healthHandler returns the API liveness status.
//
//	@Summary		Liveness check
//	@Description	Returns 200 OK when the API process is running.
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/health [get]
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	addr := ":8080"
	if v := os.Getenv("PORT"); v != "" {
		addr = ":" + v
	}

	slog.Info("starting tend-api", "addr", addr)
	if err := newRouter().Run(addr); err != nil {
		slog.Error("server exited", "err", err)
		os.Exit(1)
	}
}
