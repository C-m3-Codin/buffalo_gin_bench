package handlers

import (
	"fmt"
	"net/http"

	"github.com/C-m3-Codin/buffalo_gin_bench/gin/models"
	"github.com/C-m3-Codin/buffalo_gin_bench/gin/services"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user models.User
	userID := c.Param("userId")
	fmt.Println("\n\n\n\n got userid \n ", userID)

	err := services.DB.Where("id = ?", userID).First(&user).Error

	if err != nil {
		fmt.Println("\n\n\n\\n", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fmt.Println("got from db ", user)

	c.JSON(200, user)

}
