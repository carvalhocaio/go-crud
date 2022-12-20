package controller

import (
	"fmt"
	"net/http"

	"github.com/carvalhocaio/go-crud/src/configuration/logger"
	"github.com/carvalhocaio/go-crud/src/configuration/validation"
	"github.com/carvalhocaio/go-crud/src/controller/model/request"
	"github.com/carvalhocaio/go-crud/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   int8(userRequest.Age),
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, response)
}
