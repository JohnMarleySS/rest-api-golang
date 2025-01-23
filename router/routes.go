package router

import (
	"net/http"

	"github.com/JohnMarleySS/rest-api-golang/prisma/db"
	"github.com/gin-gonic/gin"
)

type Task struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func RunRoutes(client *db.PrismaClient) *db.PrismaClient {
	r := gin.Default()

	r.GET("/tasks", func(c *gin.Context) {

		tasks, err := client.Tasks.FindMany().Exec(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		}

		c.JSON(http.StatusOK, tasks)
	})

	r.POST("/tasks", func(c *gin.Context) {

		var task Task

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"title":     task.Title,
			"completed": task.Completed,
		})

		postTask, err := client.Tasks.CreateOne(
			db.Tasks.Title.Set(task.Title),
			db.Tasks.Completed.Set(task.Completed),
		).Exec(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
			return
		}
		c.JSON(http.StatusOK, postTask)
	})

	r.PUT("tasks/:id", func(c *gin.Context) {

		id := c.Param("id")

		var task Task

		if err := c.ShouldBindBodyWithJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updateTask, err := client.Tasks.FindUnique(
			db.Tasks.ID.Equals(id),
		).Update(
			db.Tasks.Title.Set(task.Title),
			db.Tasks.Completed.Set(task.Completed),
		).Exec(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}
		c.JSON(http.StatusOK, updateTask)

	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {

		id := c.Param("id")

		deleteTask, err := client.Tasks.FindUnique(
			db.Tasks.ID.Equals(id),
		).Delete().Exec(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, deleteTask)

	})

	r.Run(":8080")

	return client
}
