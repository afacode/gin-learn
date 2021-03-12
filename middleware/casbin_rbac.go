package middleware

import (
	"fmt"
	"gin-learn/model/request"
	"gin-learn/model/response"
	"gin-learn/service"

	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 权限校验
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := service.Casbin()
		fmt.Println("*casbin.Enforcer")
		fmt.Println(e)
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)

		if success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
