package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID    `json:"uuid"`
	Username    string       `json:"username"`
	Password    string       `json:"-"`
	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Authority   SysAuthority `json:"-" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}
