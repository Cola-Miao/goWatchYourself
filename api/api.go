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

	if input.UserName == "debug" {
		input = &models.Input{
			UserName: global.DebugUserName,
			Password: global.DebugPassword,
		}
	} else {
		fmt.Println("PassWord:")
		fmt.Scanln(&input.Password)
	}
	fmt.Println("Model:")
	fmt.Scanln(&global.Model)

	switch global.Model {
	case 1:
		fmt.Println("PreviewID:")
		fmt.Scanln(&global.PreviewID)
	case 2:
		fmt.Println("CourseID:")
		fmt.Scanln(&global.CourseID)
	}

	return
}
