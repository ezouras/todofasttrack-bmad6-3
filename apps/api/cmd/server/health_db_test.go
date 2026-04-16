package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type stubPinger struct {
	err error
}

func (s stubPinger) PingContext(_ context.Context) error { return s.err }

func TestHealthDB_NotConfigured(t *testing.T) {
	prevMode := gin.Mode()
	gin.SetMode(gin.TestMode)
	t.Cleanup(func() { gin.SetMode(prevMode) })

	router := newRouter(nil) // simulates DATABASE_URL unset

	req := httptest.NewRequest(http.MethodGet, "/health/db", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 when no DB configured, got %d", rec.Code)
	}
	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("body not JSON: %v", err)
	}
	if body["db"] != "not_configured" {
		t.Errorf(`expected db="not_configured", got %v`, body)
	}
}

func TestHealthDB_Reachable(t *testing.T) {
	prevMode := gin.Mode()
	gin.SetMode(gin.TestMode)
	t.Cleanup(func() { gin.SetMode(prevMode) })

	router := newRouter(stubPinger{err: nil})

	req := httptest.NewRequest(http.MethodGet, "/health/db", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 when DB ping succeeds, got %d", rec.Code)
	}
	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("body not JSON: %v", err)
	}
	if body["status"] != "ok" || body["db"] != "reachable" {
		t.Errorf(`expected {status:"ok", db:"reachable"}, got %v`, body)
	}
}

func TestHealthDB_Unreachable(t *testing.T) {
	prevMode := gin.Mode()
	gin.SetMode(gin.TestMode)
	t.Cleanup(func() { gin.SetMode(prevMode) })

	router := newRouter(stubPinger{err: errors.New("connection refused")})

	req := httptest.NewRequest(http.MethodGet, "/health/db", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 when DB ping fails, got %d", rec.Code)
	}
	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("body not JSON: %v", err)
	}
	if body["db"] != "unreachable" {
		t.Errorf(`expected db="unreachable", got %v`, body)
	}
	if body["error"] == "" {
		t.Errorf("expected non-empty error field, got %v", body)
	}
}
