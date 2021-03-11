package v1

import (
	"gin-learn/model/response"

	"github.com/gin-gonic/gin"
)

// 系统配置
func GetSystemConfig(c *gin.Context) {
	response.FailWithMessage("系统配置", c)
}
func SetSystemConfig(c *gin.Context) {
	response.FailWithMessage("创建失败", c)
}
func GetServerInfo(c *gin.Context) {
	response.FailWithMessage("创建失败", c)
}
func ReloadSystem(c *gin.Context) {
	response.FailWithMessage("创建失败", c)
}
