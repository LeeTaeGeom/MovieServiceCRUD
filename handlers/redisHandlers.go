package handlers

import (
	"net/http"
	"redisDBService/redis"

	"github.com/gin-gonic/gin"
)

var (
	redisDB = redis.NewRedisDB(0, 1)
)

func CreateMovie(ctx *gin.Context) {
	var movie redis.Movie
	if err := ctx.ShouldBind(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := redisDB.CreateMovieSV(ctx, &movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"movie": res,
	})
}

func GetAllMovies(ctx *gin.Context) {
	movies, err := redisDB.GetMoviesSV(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"movies": movies,
	})
}

func GetMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	movie, err := redisDB.GetMovieSV(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "movie not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})
}

func UpdateMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := redisDB.GetMovieSV(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var movie redis.Movie

	if err := ctx.ShouldBind(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res.Title = movie.Title
	res.Description = movie.Description
	res, err = redisDB.UpdateMovieSV(ctx, res)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"movie": res,
	})
}

func DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	err := redisDB.DeleteMovieSV(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "movie deleted successfully",
	})
}
