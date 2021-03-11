package v1

import (
	"gin-learn/model/request"
	"gin-learn/model/response"
	"gin-learn/service"
	"gin-learn/utils"

	"github.com/gin-gonic/gin"
)

// 获取用户菜单
func GetMenu(c *gin.Context) {
	err, menus := service.GetMenuTree(getUserAuthorityId(c))
	if err != nil {
		response.FailWithMessage("获取失败", c)

	} else {
		response.OkWithDetailed(response.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// 获取菜单（分页）目前全部请求到 部分页
func GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menuList, total := service.GetInfoList(); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// 新增菜单
func AddBaseMenu(c *gin.Context) {
	response.FailWithMessage("token黑名单", c)
}

// 获取用户动态路由
func GetBaseMenuTree(c *gin.Context) {
	response.FailWithMessage("token黑名单", c)
}

// 增加menu和角色关联关系
func AddMenuAuthority(c *gin.Context) {
	response.FailWithMessage("token黑名单", c)
}

// 获取指定角色menu
func GetMenuAuthority(c *gin.Context) {
	response.FailWithMessage("token黑名单", c)
}

// 删除菜单
func DeleteBaseMenu(c *gin.Context) {
	response.FailWithMessage("token黑名单", c)
}

// 更新菜单
func UpdateBaseMenu(c *gin.Context) {
	response.FailWithMessage("token黑名单", c)
}

// 根据id获取菜单
func GetBaseMenuById(c *gin.Context) {
	response.FailWithMessage("token黑名单", c)
}
