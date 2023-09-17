package user

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	UserGroup
}

var (
	userService = service.ServiceGroupApp.UserServiceGroup.UserService
)
