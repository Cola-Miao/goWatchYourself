package api

import (
	"fmt"
	"goWatchYourself/global"
	"goWatchYourself/models"
)

func GetUserInput() (input *models.Input) {
	input = new(models.Input)
	//fmt.Println("UserName:")
	//fmt.Scanln(&input.UserName)
	//fmt.Println("PassWord:")
	//fmt.Scanln(&input.Password)
	fmt.Println("CourseID:")
	fmt.Scanln(&global.CourseID)

	input.UserName = "15526811889"
	input.Password = "ab123456AB1234@"

	return
}
