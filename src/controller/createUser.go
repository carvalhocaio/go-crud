package controller

import (
	"fmt"

	"github.com/carvalhocaio/go-crud/src/configuration/rest_err"
	"github.com/carvalhocaio/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError("Some fields are incorrect")
		fmt.Println(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
