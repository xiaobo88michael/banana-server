package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/certificate"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/user"
)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	User        user.RouterGroup
	Certificate certificate.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
