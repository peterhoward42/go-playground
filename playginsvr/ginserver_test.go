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

// Show how the handlers can easily extract part of the URL (in this case
// the *cat* part when the route is declared in terms of a compulsory variable.
func TestUsingCompulsoryVariablesInRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/animal/cat", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "cat", w.Body.String())
}

// Show how the handlers can fish out optional and compulsory URL segments.
func TestUsingOptionalVariablesInRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	// First furnish both bits.
	req, _ := http.NewRequest("GET", "/animal/cat/stroke", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Animal type: <cat>, action: </stroke>.", w.Body.String())
}

// A comment.
func TestExtractingQueryParams(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/shouldhaveqryparams?price=£3.14&", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Price: <£3.14>, Weight: <1Kg>", w.Body.String())
}
