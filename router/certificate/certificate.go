package certificate

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CertificateRouter struct{}

func (r *CertificateRouter) InitRouter(pr *gin.RouterGroup) {
	g := pr.Group("certificate")
	api := v1.ApiGroupApp.CertApiGroup.CertApi
	{
		g.POST("user_audit_material_release", api.UserAduitMaterialsRelease)
		g.POST("company_audit_material_release", api.CompanyAduitMaterialsRelease)
		g.POST("good_release", api.ReleaseGoods)
		g.POST("update_audit_status", api.UpdateUserAuditStatus)
	}
}
