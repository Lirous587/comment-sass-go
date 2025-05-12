package router

import (
	"comment/internal/domain/user/controller"

	"github.com/gin-gonic/gin"
)

func RegisterV1(r *gin.RouterGroup, ctrl controller.Controller) error {
	g := r.Group("/v1/user")
	{
		g.Any("/login", ctrl.Login)
		g.POST("/refresh_token", ctrl.RefreshToken)
	}
	return nil
}
