package server_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/config"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/logger"
	"github.com/alechekz/online-car-auction-clean-archv2/services/vehicle/internal/platform/server"
)

// TestMain sets up the testing environment
func TestMain(m *testing.M) {
	logger.Init()
	os.Exit(m.Run())
}

// TestNewServer_HandlerResponds checks that the demo server's handler responds to requests
func TestNewServer_HandlerResponds(t *testing.T) {
	srv, _ := server.New(context.Background(), config.New())

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	srv.Handler().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
