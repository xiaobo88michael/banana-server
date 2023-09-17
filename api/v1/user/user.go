package user

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	response2 "github.com/flipped-aurora/gin-vue-admin/server/model/user/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type UserGroup struct{}

func (g *UserGroup) Login(c *gin.Context) {
	var req request.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, err := userService.UserLogin(req.QQID, req.Phone)
	if err != nil {
		if err.Error() == "isValid" {
			global.GVA_LOG.Error("登录失败！ 输入的QQ号码与手机号码不匹配!")
			response.FailWithMessage("输入的QQ号码与手机号码不匹配!", c)
			return
		}
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	if user.Enable != 1 {
		global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
		response.FailWithMessage("用户被禁止登录", c)
		return
	}
	g.TokenNext(c, *user)
	return
}

func (g *UserGroup) TokenNext(c *gin.Context, user user.User) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateUserClaims(request.BaseClaims{
		ID:    user.ID,
		QQID:  user.QQID,
		Phone: user.Phone,
	})
	token, err := j.CreateUserToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response2.LoginUserRsp{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功1", c)
		return
	}

	if jwtStr, err := userService.GetRedisJWT(user.QQID); err == redis.Nil {
		fmt.Println("jwtStr:", jwtStr, "err:", err)
		if err := userService.SetRedisJWT(token, user.QQID); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response2.LoginUserRsp{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功2", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := userService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := userService.SetRedisJWT(token, user.QQID); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response2.LoginUserRsp{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功3", c)
	}
}

func (g *UserGroup) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	err := userService.JsonInBlacklist(jwt)
	if err != nil {
		global.GVA_LOG.Error("jwt作废失败!", zap.Error(err))
		response.FailWithMessage("jwt作废失败", c)
		return
	}
	response.OkWithMessage("jwt作废成功", c)
}

func (g *UserGroup) GetUserInfoByToken(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	flag, qqID, err := j.GetUserIDFromUserToken(token)
	if err != "" {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("flag:", flag)
	fmt.Println("qqID:", qqID)
}
