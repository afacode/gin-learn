package initialize

import (
	"fmt"
	"gin-learn/middleware"
	"gin-learn/router"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	var Router = gin.Default()

	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 登录等基础不做鉴权
	}

	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		router.InitUserRouter(PrivateGroup)               // 用户路由
		router.InitMenuRouter(PrivateGroup)               // menu路由
		router.InitApiRouter(PrivateGroup)                // API管理路由
		router.InitSystemRouter(PrivateGroup)             // system路由
		router.InitJwtRouter(PrivateGroup)                // jwt黑名单
		router.InitCasbinRouter(PrivateGroup)             // 权限路由
		router.InitAuthorityRouter(PrivateGroup)          // 角色路由
		router.InitSysOperationRecordRouter(PrivateGroup) // 操作记录
	}
	fmt.Println("路由注册成功")
	return Router
}
