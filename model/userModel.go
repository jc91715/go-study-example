package model

type UserModel struct {
	Model
}

func NewUserModel() *UserModel {
	m := UserModel{}
	m.table = "userinfo"
	return &m
}
