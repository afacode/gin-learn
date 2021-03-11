package service

import (
	"errors"
	"gin-learn/global"
	"gin-learn/model"
	"gin-learn/model/request"
	"strconv"

	"gorm.io/gorm"
)

// 根据 authorityId
func GetMenuTreeMap(authorityId string) (err error, treeMap map[string][]model.SysMenu) {
	var allMenus []model.SysMenu
	treeMap = make(map[string][]model.SysMenu)
	err = global.G_DB.Where("authority_id = ?", authorityId).Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

func GetMenuTree(authorityId string) (err error, menus []model.SysMenu) {
	err, menuTree := GetMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

// 获取子菜单
func getChildrenList(menu *model.SysMenu, treeMap map[string][]model.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// 获取路由分页
func GetInfoList() (err error, list interface{}, total int64) {
	var menuList []model.SysBaseMenu
	err, treeMap := getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

// 获取路由总数 map
func getBaseMenuTreeMap() (err error, treeMap map[string][]model.SysBaseMenu) {
	var allMenus []model.SysBaseMenu
	treeMap = make(map[string][]model.SysBaseMenu)
	err = global.G_DB.Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

// 获取子菜单
func getBaseChildrenList(menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// 添加基础路由
func AddBaseMenu(menu model.SysBaseMenu) error {
	if !errors.Is(global.G_DB.Where("name= ?", menu.Name).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复的name, 请修改name")
	}
	return global.G_DB.Create(&menu).Error
}

// 获取基础路由树
func GetBaseMenuTree() (err error, menus []model.SysBaseMenu) {
	err, treeMap := getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

// 对角色增加menu树
func AddMenuAuthority(menus []model.SysBaseMenu, authorityId string) (err error) {
	var auth model.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	// err = SetMenuAuthority(&auth)
	return err
}

// 获取当前角色树
func GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []model.SysMenu) {
	err = global.G_DB.Where("authority_id", info.AuthorityId).Order("sort").Find(&menus).Error
	return err, menus
}
