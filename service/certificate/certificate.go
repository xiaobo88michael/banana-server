package certificate

import (
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/certificate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/certificate/request"
	"time"
)

type CerService struct{}

func (s *CerService) UserAuditMaterialRelease(req *request.UserAduitMaterialsReleaseReq) error {
	var thinkMember certificate.ThinkMember

	rows := global.GVA_DB.Table("think_member").Where("id=?", req.UserID).Find(&thinkMember).RowsAffected
	if rows == 0 {
		return errors.New("UserNoExist")
	}
	if thinkMember.AuditStatus != 0 {
		return errors.New("audited")
	}
	images, _ := json.Marshal(req.Images)
	thinkMember.IdCardType = req.IdCardType
	thinkMember.IdCard = req.IdCard
	thinkMember.Address = req.Address
	thinkMember.IdCardImg = string(images)
	thinkMember.Name = req.Name
	thinkMember.DueDate = req.DueDate
	thinkMember.AuditStatus = 1
	return global.GVA_DB.Table("think_member").Save(&thinkMember).Error
}

func (s *CerService) CompanyAuditMaterialRelease(req *request.CompanyAduitMaterialsReleaseReq) error {
	var thinkMec certificate.ThinkMerchantsInfo

	rows := global.GVA_DB.Table("think_merchants_info").Where("id=?", req.CompanyID).Find(&thinkMec).RowsAffected
	if rows == 0 {
		return errors.New("MerchantsNoExist")
	}
	if thinkMec.AuditStatus != 0 {
		return errors.New("audited")
	}
	mainBody, _ := json.Marshal(&req.MainBodyInfo)
	corporation, _ := json.Marshal(&req.CorporationInfo)
	operator, _ := json.Marshal(&req.OperatorInfo)
	thinkMec.CorporationInfo = string(corporation)
	thinkMec.OperatorInfo = string(operator)
	thinkMec.MainBodyInfo = string(mainBody)
	thinkMec.CompanyType = uint(req.CompanyType)
	thinkMec.AuditStatus = 1
	return global.GVA_DB.Table("think_merchants_info").Save(&thinkMec).Error
}

func (s *CerService) GoodsRelease(req *request.ReleaseGoodReq) error {
	imagesStr, _ := json.Marshal(&req.Images)
	var thinkgood = certificate.ThinkGoods{
		Title:       req.Titile,
		Description: req.Description,
		Images:      string(imagesStr),
		CreateTime:  time.Now(),
	}

	return global.GVA_DB.Table("think_goods").Create(&thinkgood).Error
}
