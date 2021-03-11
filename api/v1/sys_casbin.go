package v1

import (
	"gin-learn/model/response"

	"github.com/gin-gonic/gin"
)

// 跟新角色API权限
func UpdateCasbin(c *gin.Context) {
	response.FailWithMessage("跟新角色API权限", c)
}

// 获取权限列表
func GetPolicyPathByAuthorityId(c *gin.Context) {
	response.FailWithMessage("获取权限列表", c)
}
