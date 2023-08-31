package models

type Input struct {
	UserName string
	Password string
}

type Login struct {
	AutomaticLogon string
	UserName       string
	Password       string
}
