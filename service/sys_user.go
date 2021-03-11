package service

import (
	"fmt"
	"gin-learn/global"
	"gin-learn/model"
	"gin-learn/utils"
)

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V(u.Password)
	err = global.G_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	fmt.Println(&user)
	return err, &user
}
