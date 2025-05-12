//go:build wireinject
// +build wireinject

package user

import (
	"comment/internal/domain/user/controller"
	"comment/internal/domain/user/infrastructure/cache"
	"comment/internal/domain/user/infrastructure/db"
	"comment/internal/domain/user/router"
	"comment/internal/domain/user/service"
	"comment/pkg/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitV1(r *gin.RouterGroup) error {
	wire.Build(
		router.RegisterV1,
		controller.NewController,
		service.NewService,
		repository.GormDB,
		repository.RedisClient,
		db.NewDB,
		cache.NewCache,
	)
	return nil
}
