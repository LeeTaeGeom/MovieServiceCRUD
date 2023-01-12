package routers

import (
	"redisDBService/handlers"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/movies", handlers.CreateMovie)

	r.GET("/movies", handlers.GetAllMovies)

	r.GET("/movies/:id", handlers.GetMovie)

	// r.GET("/movies/Json/:id", handlers.GetMovieJson)

	r.PUT("/movies/:id", handlers.UpdateMovie)

	r.DELETE("/movies/:id", handlers.DeleteMovie)

	return r
}
