package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.GET("/index", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Docker")
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
