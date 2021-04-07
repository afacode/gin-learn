package v1

import (
	"gin-learn/model/response"

	"github.com/gin-gonic/gin"
)

// 生成验证码
func Captcha(c *gin.Context) {
	response.FailWithMessage("验证码获取成功", c)
}
