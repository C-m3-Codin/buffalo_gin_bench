package handlers

import (
	"net/http"

	"github.com/C-m3-Codin/buffalo_gin_bench/gin/models"
	"github.com/C-m3-Codin/buffalo_gin_bench/gin/services"
	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {

	type Users []models.User
	var users Users
	services.DB.Where("").Find(&users)

	c.JSON(http.StatusOK, users)
	return

}
