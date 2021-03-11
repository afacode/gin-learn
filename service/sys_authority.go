package service

import (
	"gin-learn/global"
	"gin-learn/model"
)

// 菜单与角色绑定
func SetMenuAuthority(auth model.SysAuthority) (err error) {
	var s model.SysAuthority
	global.G_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	// err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

// 设置角色资源权限
func SetDataAuthority(auth model.SysAuthority) (err error) {
	// var s model.SysAuthority
	// global.G_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	// err := global.G_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

// 查询子角色
func findChildrenAuthority(authority *model.SysAuthority) (err error) {
	err = global.G_DB.Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
