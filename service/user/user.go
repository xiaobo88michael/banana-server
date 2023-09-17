package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type UserService struct{}

func (s *UserService) UserLogin(qqID string, phone string) (*user.User, error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var users user.User
	rows := global.GVA_DB.Where("qq_id = ?", qqID).First(&users).RowsAffected
	if rows == 0 {
		//用户不存在直接走注册流程
		tempUser := user.User{
			QQID:  qqID,
			Phone: phone,
		}
		return &tempUser, global.GVA_DB.Create(&tempUser).Error
	} else {
		//用户存在走账号校验
		if users.Phone == phone {
			return &users, nil
		} else {
			return nil, errors.New("isValid")
		}
	}
}

func (s *UserService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

func (s *UserService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	fmt.Println("time:", timer)
	err = global.GVA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	fmt.Printf("err:", err)
	return err
}

func (s *UserService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

func (s *UserService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}
