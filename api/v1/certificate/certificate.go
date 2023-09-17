package certificate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	cerReq "github.com/flipped-aurora/gin-vue-admin/server/model/certificate/request"
	cerRes "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type CertApi struct{}

// UserAduitMaterialsRelease
// @Tags      UserAduitMaterialsRelease
// @Summary   个人资料上传
// @accept    application/json
// @Produce   application/json
// @Param     data  body      cerReq.UserAduitMaterialsReleaseReq      true  "个人资料上传"
// @Success   200   {object}  cerRes.Response{msg=string}  "个人资料上传成功"
// @Router    /certificate/user_audit_material_release [post]
func (s *CertApi) UserAduitMaterialsRelease(c *gin.Context) {
	var req cerReq.UserAduitMaterialsReleaseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cerRes.FailWithMessage(err.Error(), c)
		return
	}
	err := cerService.UserAuditMaterialRelease(&req)
	if err != nil {
		if err.Error() == "UserNoExist" {
			global.GVA_LOG.Error("用户不存在")
			cerRes.FailWithMessage("用户不存在", c)
			return
		}
		global.GVA_LOG.Error(err.Error())
		cerRes.FailWithMessage(err.Error(), c)
		return
	}
	cerRes.Result(cerRes.SUCCESS, nil, "个人资料上传成功", c)
}

// CompanyAduitMaterialsRelease
// @Tags      CompanyAduitMaterialsRelease
// @Summary   企业资料上传
// @accept    application/json
// @Produce   application/json
// @Param     data  body      cerReq.CompanyAduitMaterialsReleaseReq     true  "企业资料上传"
// @Success   200   {object}  cerRes.Response{msg=string}  "企业资料上传成功"
// @Router    /certificate/company_audit_material_release [post]
func (s *CertApi) CompanyAduitMaterialsRelease(c *gin.Context) {
	var req cerReq.CompanyAduitMaterialsReleaseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cerRes.FailWithMessage(err.Error(), c)
		return
	}
	err := cerService.CompanyAuditMaterialRelease(&req)
	if err != nil {
		if err.Error() == "MerchantsNoExist" {
			global.GVA_LOG.Error("商户不存在")
			cerRes.FailWithMessage("商户不存在", c)
			return
		}
		global.GVA_LOG.Error(err.Error())
		cerRes.FailWithMessage(err.Error(), c)
		return
	}
	cerRes.Result(cerRes.SUCCESS, nil, "企业资料上传成功", c)
}

// ReleaseGoods
// @Tags      ReleaseGoods
// @Summary   拍品发布
// @accept    application/json
// @Produce   application/json
// @Param     data  body      cerReq.ReleaseGoodReq     true  "拍品发布"
// @Success   200   {object}  cerRes.Response{msg=string}  "拍品发布成功"
// @Router    /certificate/good_release [post]
func (s *CertApi) ReleaseGoods(c *gin.Context) {
	var req cerReq.ReleaseGoodReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cerRes.FailWithMessage(err.Error(), c)
		return
	}
	err := cerService.GoodsRelease(&req)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		cerRes.FailWithMessage(err.Error(), c)
		return
	}
	cerRes.Result(cerRes.SUCCESS, nil, "拍品发布成功", c)
}