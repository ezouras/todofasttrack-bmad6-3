// Package main is the entry point for the Tend API server.
//
//	@title			Tend API
//	@version		0.1.0
//	@description	REST API for the Tend planning app.
//	@BasePath		/
package main

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// DBPinger is the minimal contract /health/db needs. *sql.DB satisfies it.
// Defined as an interface so tests can stub without spinning up Postgres.
type DBPinger interface {
	PingContext(ctx context.Context) error
}

func newRouter(db DBPinger) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", healthHandler)
	r.GET("/health/db", healthDBHandler(db))

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

// healthDBHandler returns DB connectivity status. Returns 503 if DATABASE_URL
// is not configured (db is nil) or if the ping fails.
//
//	@Summary		DB connectivity check
//	@Description	Returns 200 if the API can reach Postgres, 503 otherwise.
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		503	{object}	map[string]string
//	@Router			/health/db [get]
func healthDBHandler(db DBPinger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if db == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "degraded",
				"db":     "not_configured",
				"error":  "DATABASE_URL is not set",
			})
			return
		}
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		if err := db.PingContext(ctx); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "degraded",
				"db":     "unreachable",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"db":     "reachable",
		})
	}
}

// openDB returns a configured *sql.DB if DATABASE_URL is set, or nil if unset.
// Returns an error only when DATABASE_URL is set but malformed.
func openDB(dsn string) (*sql.DB, error) {
	if dsn == "" {
		return nil, nil
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// Conservative pool sizing for a single Railway instance behind /health/db.
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxIdleTime(5 * time.Minute)
	return db, nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	addr := ":8080"
	if v := os.Getenv("PORT"); v != "" {
		addr = ":" + v
	}

	db, err := openDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("failed to configure database", "err", err)
		os.Exit(1)
	}

	// Avoid the typed-nil-pointer-in-interface trap: only assign db to the
	// pinger interface when it's non-nil. Otherwise pinger stays a true nil
	// interface and the handler's `db == nil` check works as expected.
	var pinger DBPinger
	if db != nil {
		pinger = db
		slog.Info("database configured")
	} else {
		slog.Warn("DATABASE_URL not set — /health/db will report 503")
	}

	slog.Info("starting tend-api", "addr", addr)
	if err := newRouter(pinger).Run(addr); err != nil {
		slog.Error("server exited", "err", err)
		os.Exit(1)
	}
}
