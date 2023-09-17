package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/certificate"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/user"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	UserApiGroup    user.ApiGroup
	ExampleApiGroup example.ApiGroup
	CertApiGroup    certificate.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
