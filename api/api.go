package api

import (
	"fmt"
	"goWatchYourself/global"
	"goWatchYourself/models"
)

func GetUserInput() (input *models.Input) {
	input = new(models.Input)
	fmt.Println("UserName:")
	fmt.Scanln(&input.UserName)
	fmt.Println("PassWord:")
	fmt.Scanln(&input.Password)
	fmt.Println("CourseID:")
	fmt.Scanln(&global.CourseID)

	return
}
