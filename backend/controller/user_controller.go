package controller

import (
	"net/http"

	"github.com/AfomiaTadesse/Afomia_M/backend/domain"
	"github.com/AfomiaTadesse/Afomia_M/backend/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (ctrl *UserController) Signup(c *gin.Context) {
	var req domain.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.AuthResponse{
			Success: false,
			Message: "Invalid request body",
			Errors:  []string{err.Error()},
		})
		return
	}

	response, err := ctrl.userUsecase.Signup(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.AuthResponse{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	status := http.StatusOK
	if !response.Success {
		status = http.StatusBadRequest
	}

	c.JSON(status, response)
}

func (ctrl *UserController) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.AuthResponse{
			Success: false,
			Message: "Invalid request body",
			Errors:  []string{err.Error()},
		})
		return
	}

	response, err := ctrl.userUsecase.Login(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.AuthResponse{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	status := http.StatusOK
	if !response.Success {
		status = http.StatusUnauthorized
	}

	c.JSON(status, response)
}