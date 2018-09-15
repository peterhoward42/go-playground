package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPingRoute is the canonical test example for gin.
// It appears to let you exercise the router and the corresponding handler
// behaviour without launching a server.
func TestPingRoute(t *testing.T) {
	// Build the routes.
	router := setupRouter()

	// Exploit a standard library Writer that records what you write to it.
	w := httptest.NewRecorder()
	// Construct an http request that will exercise a handled route.
	req, _ := http.NewRequest("GET", "/ping", nil)
	// Instruct the router to handle this request.
	router.ServeHTTP(w, req)

	// Scrutinise what the writer has recorded.
	assert.Equal(t, 200, w.Code)             // Response code OK
	assert.Equal(t, "pong", w.Body.String()) // Body content.
}
