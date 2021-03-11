package v1

import (
	"fmt"
	"gin-learn/middleware"
	"gin-learn/model"
	"gin-learn/model/request"
	"gin-learn/model/response"
	"gin-learn/service"
	"gin-learn/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	// _ = c.ShouldBind(&L) // form
	if err := utils.Verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 数据库验证 用户
	u := &model.SysUser{Username: L.Username, Password: L.Password}
	err, user := service.Login(u)
	if err != nil {
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		fmt.Println(user)
		tokenNext(c, *user)
	}
}

// 登录成功签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{SigningKey: []byte("bgbiao.top")}

	claims := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),      // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600*24*7), // 过期时间 7天
			Issuer:    "cfun",                               //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)

	if err != nil {
		response.FailWithMessage("设置登录状态失败", c)
		return
	}

	response.OkWithDetailed(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
	return
}

func Register(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "注册",
		"data": L,
	})
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func getUserAuthorityId(c *gin.Context) string {
	claims, exists := c.Get("claims")
	if !exists {
		fmt.Println("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.AuthorityId
	}
}

// 修改密码
func ChangePassword(c *gin.Context) {
	response.FailWithMessage("创建失败", c)
}

// 分页获取用户列表
func GetUserList(c *gin.Context) {
	response.FailWithMessage("创建失败", c)
}

// 设置用户角色 权限
func SetUserAuthority(c *gin.Context) {
	response.FailWithMessage("创建失败", c)
}

// 删除用户
func DeleteUser(c *gin.Context) {
	response.FailWithMessage("创建失败", c)
}

// 设置用户信息
func SetUserInfo(c *gin.Context) {
	response.FailWithMessage("创建失败", c)
}
