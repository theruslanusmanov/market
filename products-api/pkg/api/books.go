package api

import (
	"context"
	"encoding/json"
	"market/products-api/pkg/cache"
	"market/products-api/pkg/database"
	"market/products-api/pkg/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BookRepository interface {
	Healthcheck(c *gin.Context)
	FindBooks(c *gin.Context)
	//CreateBook(c *gin.Context)
	//FindBook(c *gin.Context)
	//UpdateBook(c *gin.Context)
	//DeleteBook(c *gin.Context)
}

type bookRepository struct {
	DB          database.Database
	RedisClient cache.Cache
	Ctx         *context.Context
}

func NewBookRepository(db database.Database, redisClient cache.Cache, ctx *context.Context) *bookRepository {
	return &bookRepository{
		DB:          db,
		RedisClient: redisClient,
		Ctx:         ctx,
	}
}

func (r *bookRepository) Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func (r *bookRepository) FindBooks(c *gin.Context) {
	var books []models.Book

	offsetQuery := c.DefaultQuery("offset", "0")
	limitQuery := c.DefaultQuery("limit", "10")

	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset format"})
		return
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit format"})
		return
	}

	cacheKey := "books_offset_" + offsetQuery + "_limit_" + limitQuery

	cachedBooks, err := r.RedisClient.Get(*r.Ctx, cacheKey).Result()
	if err != nil {
		err := json.Unmarshal([]byte(cachedBooks), &books)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal cached data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": books})
		return
	}

	r.DB.Offset(offset).Limit(limit).Find(&books)

	serializedBooks, err := json.Marshal(books)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to marshal data"})
		return
	}
	err = r.RedisClient.Set(*r.Ctx, cacheKey, serializedBooks, time.Minute).Err()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to set cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}
