package user

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitRouter(pr *gin.RouterGroup) {
	g := pr.Group("user")
	api := v1.ApiGroupApp.UserApiGroup.UserGroup
	{
		g.POST("login", api.Login)
		g.POST("get_qq_by_token", middleware.JWTAuthUser(), api.GetUserInfoByToken)
		g.POST("logout", middleware.JWTAuthUser(), api.JsonInBlacklist)
	}
}
