package main

import (
	"context"
	"log"
	"market/products-api/pkg/api"
	"market/products-api/pkg/cache"
	"market/products-api/pkg/database"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	redisClient := cache.NewRedisClient()
	db := database.NewDatabase()
	dbWrapper := &database.GormDatabase{DB: db}
	ctx := context.Background()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	gin.SetMode(gin.DebugMode)

	r := api.NewRouter(logger, dbWrapper, redisClient, &ctx)

	if err := r.Run(":8001"); err != nil {
		log.Fatal(err)
	}
}
