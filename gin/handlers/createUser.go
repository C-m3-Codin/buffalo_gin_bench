package handlers

import (
	"fmt"
	"net/http"

	"github.com/C-m3-Codin/buffalo_gin_bench/gin/models"
	"github.com/C-m3-Codin/buffalo_gin_bench/gin/services"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func CreateUser(c *gin.Context) {
	var err error
	user := models.User{}
	user.ID, err = uuid.NewV1()
	if err != nil {
		fmt.Println("Error while generating uuid")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fmt.Println("got id ", user.ID)
	err = c.BindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = services.DB.Create(&user).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, user)

}
