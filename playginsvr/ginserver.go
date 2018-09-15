package main

import "github.com/gin-gonic/gin"

// setupRouter is is a function that can be used to set up routes
// in a reusable and deterministic way, good for incorporating into a
// server, *and* for unit tests.
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

// main is the function that brings up the server. It delegates the setting up
// of its routes to an external function. (to facilitate testing)
func main() {
	r := setupRouter()
	r.Run(":8080")
}
