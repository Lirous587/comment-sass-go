//go:build wireinject
// +build wireinject

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"comment/internal/domain/admin/controller"
	"comment/internal/domain/admin/infrastructure/cache"
	"comment/internal/domain/admin/infrastructure/db"
	"comment/internal/domain/admin/router"
	"comment/internal/domain/admin/service"
	"comment/internal/domain/admin/worker"
	"comment/pkg/repository"
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

func InitWorker() worker.Worker {
	gormDB := repository.GormDB()
	dbDB := db.NewDB(gormDB)
	client := repository.RedisClient()
	cacheCache := cache.NewCache(client)
	workerWorker := worker.NewWorker(dbDB, cacheCache)
	return workerWorker
}
