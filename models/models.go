package models

type User struct {
	UserName string `form:"userName"`
	Password string `form:"password"`
}

type Class struct {
	Model int    `form:"model"`
	ID    string `form:"classID"`
}

type Form struct {
	User
	Class
}
