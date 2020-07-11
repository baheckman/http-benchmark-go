package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FibonacciLoop(n uint64) uint64 {
	f := make([]uint64, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := uint64(2); i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Computation test
	r.GET("/fib", func(c *gin.Context) {
		c.String(http.StatusOK, strconv.FormatUint(FibonacciLoop(93), 10))
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
