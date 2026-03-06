package api

import (
	"context"
	"market/products-api/pkg/database"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	LoginHandler(c *gin.Context)
	RegisterHandler(c *gin.Context)
}

// bookRepository holds shared resources like database and Redis client
type userRepository struct {
	DB  database.Database
	Ctx *context.Context
}

func NewUserRepository(db database.Database, ctx *context.Context) *userRepository {
	return &userRepository{
		DB:  db,
		Ctx: ctx,
	}
}
