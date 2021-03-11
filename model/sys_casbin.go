package model

type CasbinModel struct {
	ID          int    `json:"id" gorm:"column:id"`
	Ptype       string `json:"p_type" gorm:"column:p_type"`
	AuthorityId string `json:"rolename" gorm:"column:v0"`
	Path        string `json:"path" gorm:"column:v1"`
	Method      string `json:"method" gorm:"column:v2"`
}
