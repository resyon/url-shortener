package store

import (
	// "context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storageService = &StorageService{}
	// ctx = context.Background()
)

// TODO: move to configuration file
const CacheDuration = 6 * time.Hour

func InitStorageService() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "resyon.cn:6379",
		Password: "",
		DB: 0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	
	fmt.Printf("\n Redis started successfully: pong message: {%s}", pong)
	storageService.redisClient = redisClient

	return storageService
}

func SaveUrlMapping(shortUrl, originalUrl, userId string) {
	err := storageService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %v", err, originalUrl))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storageService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed Retrieving URL | Error: %v - shortUrl: %v", err, shortUrl))
	}
	return result
}


