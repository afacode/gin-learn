package v1

import (
	"gin-learn/model/response"

	"github.com/gin-gonic/gin"
)

// 发送测试邮件
func EmailTest(c *gin.Context) {
	response.OkWithData("发送成功", c)
}
