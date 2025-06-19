package controller

import (
	"net/http"
	"strconv"

	"github.com/AfomiaTadesse/Afomia_M/backend/domain"
	"github.com/AfomiaTadesse/Afomia_M/backend/usecase"
	"github.com/gin-gonic/gin"
)

type MovieController struct {
	movieUsecase usecase.MovieUsecase
}

func NewMovieController(movieUsecase usecase.MovieUsecase) *MovieController {
	return &MovieController{movieUsecase: movieUsecase}
}

func (ctrl *MovieController) CreateMovie(c *gin.Context) {
	var req domain.CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.BaseResponse{
			Success: false,
			Message: "Invalid request body",
			Errors:  []string{err.Error()},
		})
		return
	}

	userID, _ := c.Get("userID")
	req.UserID = userID.(string)

	response, err := ctrl.movieUsecase.CreateMovie(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.BaseResponse{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (ctrl *MovieController) GetMovies(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	response, err := ctrl.movieUsecase.GetMovies(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.PaginatedResponse{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *MovieController) SearchMovies(c *gin.Context) {
	title := c.Query("title")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	response, err := ctrl.movieUsecase.SearchMovies(title, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.PaginatedResponse{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *MovieController) GetMovieByID(c *gin.Context) {
	id := c.Param("id")

	response, err := ctrl.movieUsecase.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.BaseResponse{
			Success: false,
			Message: "Movie not found",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *MovieController) UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	
	var req domain.UpdateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.BaseResponse{
			Success: false,
			Message: "Invalid request body",
			Errors:  []string{err.Error()},
		})
		return
	}

	userID, _ := c.Get("userID")

	response, err := ctrl.movieUsecase.UpdateMovie(id, userID.(string), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.BaseResponse{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	if !response.Success {
		c.JSON(http.StatusForbidden, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *MovieController) DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	
	userID, _ := c.Get("userID")

	response, err := ctrl.movieUsecase.DeleteMovie(id, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.BaseResponse{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	if !response.Success {
		c.JSON(http.StatusForbidden, response)
		return
	}

	c.JSON(http.StatusOK, response)
}