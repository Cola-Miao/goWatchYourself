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

type Replay struct {
	Duration float64
	CourseId float64
	VideoId  float64
}
