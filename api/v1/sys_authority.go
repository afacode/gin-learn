package v1

import (
	"gin-learn/model/response"

	"github.com/gin-gonic/gin"
)

func CreateAuthority(c *gin.Context) {
	response.FailWithMessage("创建角色", c)
}
func DeleteAuthority(c *gin.Context) {
	response.FailWithMessage("删除角色", c)
}
func UpdateAuthority(c *gin.Context) {
	response.FailWithMessage("更新角色", c)
}
func CopyAuthority(c *gin.Context) {
	response.FailWithMessage("复制角色", c)
}
func GetAuthorityList(c *gin.Context) {
	response.FailWithMessage("获取角色列表", c)
}
func SetDataAuthority(c *gin.Context) {
	response.FailWithMessage("设置角色资源权限", c)
}
