package controllers

import (
	"net/http"

	"example.com/goflix/models"
	"example.com/goflix/services"
	"github.com/gin-gonic/gin"
)

type MovieController struct {
	MovieServices services.MovieServices
}

func New(movieservice services.MovieServices) MovieController {
	return MovieController{
		MovieServices: movieservice,
	}
} 

func (uc *MovieController) CreateMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.MovieServices.CreateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *MovieController) GetMovie(ctx *gin.Context) {
	moviename := ctx.Param("name")
	movie, err := uc.MovieServices.GetMovie(&moviename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movie)

}

func (uc *MovieController) GetAll(ctx *gin.Context) {
	movies, err := uc.MovieServices.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (uc *MovieController) UpdateMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.MovieServices.UpdateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *MovieController) DeleteMovie(ctx *gin.Context) {
	var moviename string = ctx.Param("name")
	err := uc.MovieServices.DeleteMovie(&moviename)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *MovieController) RegisterMovieRoutes(rg *gin.RouterGroup) {
	movieroute := rg.Group("/movie")
	movieroute.POST("/create", uc.CreateMovie)
	movieroute.GET("/get/:name", uc.GetMovie)
	movieroute.GET("/getall", uc.GetAll)
	movieroute.PATCH("/update", uc.UpdateMovie)
	movieroute.DELETE("/delete", uc.DeleteMovie)
}
