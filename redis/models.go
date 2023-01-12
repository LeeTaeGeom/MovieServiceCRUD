package redis

import (
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Json        map[string]string `json:"json"`
}

type MovieService interface {
	GetMovieSV(ctx *gin.Context, id string) (*Movie, error)
	GetMoviesSV(ctx *gin.Context) ([]*Movie, error)
	CreateMovieSV(ctx *gin.Context, movie *Movie) (*Movie, error)
	UpdateMovieSV(ctx *gin.Context, movie *Movie) (*Movie, error)
	DeleteMovieSV(ctx *gin.Context, id string) error
}

// var ctx = context.Background()

type redisDB struct {
	host string
	db   int
	exp  time.Duration
}

func getEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}

func NewRedisDB(db int, exp time.Duration) MovieService {

	host := getEnvVar("DB_HOST")
	return &redisDB{
		host: host,
		db:   db,
		exp:  exp,
	}
}

func (db redisDB) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     db.host,
		Password: "",
		DB:       db.db,
	})
}
