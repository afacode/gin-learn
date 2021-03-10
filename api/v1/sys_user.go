package v1

import (
	"gin-app/model/request"

	"github.com/gin-gonic/gin"
)

// 登录
func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
}
