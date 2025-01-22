package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() {
	r := gin.Default()

	r.GET("/tasks", func(c *gin.Context) {
		c.String(http.StatusOK, "GET")
	})

	r.POST("/tasks", func(c *gin.Context) {
		c.String(http.StatusOK, "POST")
	})

	r.PUT("tasks/{id}", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT")
	})

	r.DELETE("/tasks/{id}", func(c *gin.Context) {
		c.String(http.StatusOK, "DELETE")
	})

	r.Run(":8080")
}
