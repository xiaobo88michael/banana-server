package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/golang-jwt/jwt/v4"
)

type GetRiskUsersListParams struct {
	request.PageInfo
}

type LoginReq struct {
	QQID  string `json:"qqID"`
	Phone string `json:"phone"`
}

type GetUserDetailReq struct {
	UserID string `json:"user_id" form:"user_id"`
}

type GetAppealInfoReq struct {
	AppealID int `json:"appeal_id"`
}

type GetComplaintInfoReq struct {
	ComplaintID int `json:"complaint_id"`
}

type GetConversationListReq struct {
	UserID string `json:"user_id" form:"user_id"`
	request.PageInfo
}

type DealAppealReq struct {
	AppealID     int    `json:"appeal_id"`
	DealType     int    `json:"deal_type"`
	RefuseReason string `json:"refuse_reason"`
}

type DealComplaintReq struct {
	ComplaintID int    `json:"complaint_id"`
	DealType    int    `json:"deal_type"`
	TagReason   string `json:"tagReason"`
}

type DealUserRiskTag struct {
	UserID      string `json:"user_id"`
	DealType    int    `json:"deal_type"`
	LimitedType int    `json:"limited_type"`
}

type GetConversionDetailReq struct {
	FromUserID string `json:"from_user_id" form:"from_user_id"`
	ToUserID   string `json:"to_user_id" form:"to_user_id"`
	request.PageInfo
}

type GetUserInfosListReq struct {
	request.PageInfo
	AuthStatus int `json:"authStatus"`
}

type GetDeviceInfoByUserIDReq struct {
	UserID string `json:"user_id"`
}

type GetUserAssetByUserIDReq struct {
	UserID  string `json:"user_id"`
	KeyWord string `json:"keyword"`
}

type GetUserAppealInfoListReq struct {
	request.PageInfo
	QueryType     int `json:"query_type"`
	ComplaintType int `json:"complaint_type"`
}

type RemoveRiskControlReq struct {
	UserID string `json:"user_id"`
}

type TagRiskControlReq struct {
	UserID string `json:"user_id"`
}

type SwitchInfosReq struct {
	RedPacket     int `json:"red_packet"`
	GroupApp      int `json:"group_app"`
	Wallet        int `json:"wallet"`
	Transfer      int `json:"transfer"`
	GroupWithDraw int `json:"group_with_draw"`
}

type UserClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID    uint
	QQID  string
	Phone string
}
