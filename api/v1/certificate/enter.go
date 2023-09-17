package certificate

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CertApi
}

var (
	cerService = service.ServiceGroupApp.CertiServiceGroup.CerService
)
