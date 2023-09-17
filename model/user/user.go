package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type User struct {
	global.GVA_MODEL
	// 用户UUID
	Username string `json:"userName" gorm:"index;comment:用户昵称"`                                                // 用户登录名
	Password string `json:"-"  gorm:"comment:用户登录密码"`                                                          // 用户登录密码
	QQID     string `json:"qqID" gorm:"comment:用户QQ号"`                                                         // 用户昵称
	Avatar   string `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Phone    string `json:"phone"  gorm:"comment:用户手机号"`                                                       // 用户手机号
	Email    string `json:"email"  gorm:"comment:用户邮箱"`                                                        // 用户邮箱
	Enable   int    `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                   //用户是否被冻结 1正常 2冻结
}

func (User) TableName() string {
	return "user"
}
