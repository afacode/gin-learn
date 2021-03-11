package middleware

import (
	"gin-learn/model/response"

	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 权限校验
		success := true
		if success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
