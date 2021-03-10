package model

import (
	"gorm.io/gorm"
)

type Sysuser struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	NickName string    `json:"nickName" gorm:"default:后台用户"`
}
