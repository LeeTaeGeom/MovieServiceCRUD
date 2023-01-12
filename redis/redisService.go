package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
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

// ------------------

func (db redisDB) CreateMovieSV(ctx *gin.Context, movie *Movie) (*Movie, error) {
	c := db.getClient()
	movie.ID = uuid.New().String()
	fmt.Println("movie.ID(UUID) : ", movie.ID)
	json, err := json.Marshal(movie)
	if err != nil {
		return nil, err
	}
	c.HSet(ctx, "movies", movie.ID, json)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (db redisDB) GetMovieSV(ctx *gin.Context, id string) (*Movie, error) {
	c := db.getClient()
	val, err := c.HGet(ctx, "movies", id).Result()
	fmt.Println(val, err)
	if err == redis.Nil {
		fmt.Printf("key does not exist")
	}
	if err != nil {
		return nil, err
	}
	movie := &Movie{}
	err = json.Unmarshal([]byte(val), movie)

	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (db redisDB) GetMoviesSV(ctx *gin.Context) ([]*Movie, error) {
	c := db.getClient()
	movies := []*Movie{}
	val, err := c.HGetAll(ctx, "movies").Result()
	if err != nil {
		return nil, err
	}
	for _, item := range val {
		movie := &Movie{}
		err := json.Unmarshal([]byte(item), movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (db redisDB) UpdateMovieSV(ctx *gin.Context, movie *Movie) (*Movie, error) {
	c := db.getClient()
	json, err := json.Marshal(&movie)
	if err != nil {
		return nil, err
	}
	c.HSet(ctx, "movies", movie.ID, json)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
func (db redisDB) DeleteMovieSV(ctx *gin.Context, id string) error {
	c := db.getClient()
	numDeleted, err := c.HDel(ctx, "movies", id).Result()
	if numDeleted == 0 {
		return errors.New("movie to delete not found")
	}
	if err != nil {
		return err
	}
	return nil
}
