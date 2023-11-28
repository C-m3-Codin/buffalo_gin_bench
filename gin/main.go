package main

import (
	"log"

	"github.com/C-m3-Codin/buffalo_gin_bench/gin/handlers"
	"github.com/C-m3-Codin/buffalo_gin_bench/gin/services"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	_, err = services.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	users := r.Group("/users")
	{
		users.GET(":userId", handlers.GetUser)
		users.GET("", handlers.ListUsers)
		users.POST("", handlers.CreateUser)
		users.PUT(":userId", handlers.UpdateUser)
		// users.PUT()

	}

	r.Run("0.0.0.0:3000") // listen and serve on 0.0.0.0:8080
}
