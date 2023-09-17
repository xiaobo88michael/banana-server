package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/shopspring/decimal"
	"time"
)

type RiskUserInfo struct {
	Category string `json:"category"`
	Addr     string `json:"addr" gorm:"column:user_id"`
	FCID     string `json:"fcid" gorm:"column:freechat_id"`
}

type GetRiskUsersListRsp struct {
	Total int64           `json:"total"`
	Data  []*RiskUserInfo `json:"data"`
}

type LoginUserRsp struct {
	User      user.User `json:"user"`
	Token     string    `json:"token"`
	ExpiresAt int64     `json:"expiresAt"`
}

type GetUserInfosListRsp struct {
	Total int64            `json:"total"`
	Data  []*UserInfosList `json:"userInfos"`
}

type UserInfosList struct {
	UserID   string `json:"userId" gorm:"column:user_id"`
	Name     string `json:"name" gorm:"column:name"`
	FCID     string `json:"fcid" gorm:"column:freechat_id"`
	IsLimit  int64  `json:"islimit" gorm:"column:islimit"`
	IsMalice int64  `json:"isMalice" gorm:"column:isMalice"`
}

type UserDetail struct {
	Nick   string                     `json:"nick" gorm:"column:name"`
	Avatar string                     `json:"avatar" gorm:"column:face_url"`
	Addr   string                     `json:"addr" gorm:"column:user_id"`
	FCID   string                     `json:"fcid" gorm:"column:freechat_id"`
	Gender int                        `json:"gender" gorm:"column:gender"`
	Assets map[string]decimal.Decimal `json:"assets"`
}

type GetDeviceInfosByUserIDRsp struct {
	Total int64          `json:"total"`
	Data  []*DeviceInfos `json:"deviceInfos"`
}

type GetAppealInfoRsp struct {
	Content string   `json:"content" gorm:"content"`
	Images  []string `json:"images" gorm:"images"`
	Video   []string `json:"video" gorm:"video"`
}

type GetUserAssetByUserIDRsp struct {
	Total int64         `json:"total"`
	Data  []*AssetInfos `json:"assets"`
}

type GetUserActiveRsp struct {
	User UserActiveInfo `json:"user"`
}

type GetUserActiveData struct {
	Data GetUserActiveRsp `json:"data"`
}

type UserActiveInfo struct {
	Total       int `json:"total"`
	DayActive   int `json:"dayActive"`
	WeekActive  int `json:"weekActive"`
	MonthActive int `json:"monthActive"`
}

type GetUserAppealInfoListRsp struct {
	Total int64          `json:"total"`
	Data  []*AppealInfos `json:"appealInfos"`
}

type GetUserComplaintInfoListRsp struct {
	Total int64            `json:"total"`
	Data  []*ComplaintInfo `json:"complaintInfos"`
}

type ComplaintInfo struct {
	ID               int    `json:"id" gorm:"column:id"`
	UserID           string `json:"user_id" gorm:"column:user_id"`
	GroupID          string `json:"group_id"`
	ComplaintID      string `json:"complaint_id" gorm:"column:complaint_id"`
	ComplaintType    int    `json:"complaint_type" gorm:"column:complaint_type"`
	ComplaintContent string `json:"complaint_content"`
	PostID           int    `json:"post_id"`
	CommentID        int    `json:"comment_id"`
	Status           int    `json:"status"`
}
type AppealInfos struct {
	ID     int64  `json:"id" gorm:"column:id"`
	UserID string `json:"user_id" gorm:"column:user_id"`
	Name   string `json:"name" gorm:"column:name"`
	FCID   string `json:"freechat_id" gorm:"column:freechat_id"`
	Status int    `json:"status" gorm:"status"`
}

type AssetInfos struct {
	CoinID   int             `json:"coin_id" gorm:"column:coin_id"`
	CoinName string          `json:"coin_name" gorm:"column:symbol"`
	CoinLogo string          `json:"logo_url" gorm:"column:logo_url"`
	Amount   decimal.Decimal `json:"amount" gorm:"column:amount"`
}

type DeviceInfos struct {
	DeviceFinger string   `json:"device_finger"`
	DeviceId     string   `json:"device_id"`
	Ip           []string `json:"ip"`
	Account      []string `json:"account"`
	DeviceOs     string   `json:"device_os"`
}

type UserConversationInfo struct {
	Object string `json:"object"`
	Scene  string `json:"scene"`
}

type SenderAndReceiver struct {
	SendID    string `gorm:"column:send_id"`
	ReceiveID string `gorm:"column:recv_id"`
}

type GetUserConversationListRsp struct {
	Total int64                   `json:"total"`
	Data  []*UserConversationInfo `json:"data"`
}

type ConversationContent struct {
	Time        time.Time `json:"time" gorm:"column:send_time"`
	Content     string    `json:"content" gorm:"column:content"`
	ContentType int       `json:"content_type" gorm:"column:content_type"`
}

type GetConversationContentListRsp struct {
	Total int64                  `json:"total"`
	Data  []*ConversationContent `json:"data"`
}

type UserAsset struct {
	Amount   decimal.Decimal `gorm:"column:amount"`
	CoinName string          `gorm:"column:coin_name"`
}

type GetAppSwitchRsp struct {
	RedPacket     SwitchMeta `json:"red_packet"`
	GroupApp      SwitchMeta `json:"group_app"`
	Wallet        SwitchMeta `json:"wallet"`
	Transfer      SwitchMeta `json:"transfer"`
	GroupWithDraw SwitchMeta `json:"group_with_draw"`
}

type SwitchInfos struct {
	RedPacket     int `json:"red_packet"`
	GroupApp      int `json:"group_app"`
	Wallet        int `json:"wallet"`
	Transfer      int `json:"transfer"`
	GroupWithDraw int `json:"group_with_draw"`
}

type SwitchMeta struct {
	Status int    `json:"status"`
	Tag    string `json:"tag"`
}
