package handlers

import (
	"fmt"

	"github.com/C-m3-Codin/buffalo_gin_bench/gin/models"
	"github.com/C-m3-Codin/buffalo_gin_bench/gin/services"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func UpdateUser(c *gin.Context) {
	uid := uuid.FromStringOrNil(c.Param("userId"))

	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println("Panic at disco")
		panic("binding failed")
		return
	}

	user.ID = uid

	fmt.Println("\n\n\n\n\n", user, "\n\n\n")

	err = services.DB.UpdateColumns(user).Error
	if err != nil {
		fmt.Println("Panic at disco")
		panic("Error on update failed")
		return
	}

	return
}
