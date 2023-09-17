package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/certificate"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/user"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	UserServiceGroup    user.ServiceGroup
	CertiServiceGroup   certificate.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
