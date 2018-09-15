package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// setupRouter is is a function that can be used to set up routes
// in a reusable and deterministic way, good for incorporating into a
// server, *and* for unit tests.
func setupRouter() *gin.Engine {
	r := gin.Default()

	// The canonical route/handler.
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// A route with a *compulsory* variable segment.
	r.GET("/animal/:animalType", func(c *gin.Context) {
		// Extract the variable value and return it as a string.
		animalType := c.Param("animalType")
		c.String(http.StatusOK, animalType)
	})

	// A route with both a *compulsory* and optional variable segments.
	r.GET("/animal/:animalType/*action", func(c *gin.Context) {
		// Extract the compulsory and optional variable value (if is present)
		// and return them as a string.
		animalType := c.Param("animalType")
		action := c.Param("action")
		c.String(http.StatusOK, "Animal type: <%s>, action: <%s>.",
			animalType, action)
	})

	// How to extract query strings from GET URLs.
	r.GET("/shouldhaveqryparams", func(c *gin.Context) {
		price := c.Query("price")
		weight := c.DefaultQuery("weight", "1Kg")
		c.String(http.StatusOK, "Price: <%s>, Weight: <%s>", price, weight)
	})

	return r

}

// main is the function that brings up the server. It delegates the setting up
// of its routes to an external function. (to facilitate testing)
func main() {
	r := setupRouter()
	r.Run(":8080")
}
