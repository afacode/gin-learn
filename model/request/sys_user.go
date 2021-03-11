package request

import uuid "github.com/satori/go.uuid"

type Login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Register struct {
}

type ChangePasswordStruct struct {
}

type SetUserAuth struct {
	UUID uuid.UUID `json:"uuid"`
}
