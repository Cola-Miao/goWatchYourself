package watcher

import (
	"fmt"
	"goWatchYourself/global"
	"goWatchYourself/method"
	"goWatchYourself/models"
	"goWatchYourself/text"
)

func Watcher(form *models.Form) (err error) {
	cookies, err := method.Login(&form.User)
	if err != nil {
		return
	}

	switch form.Model {
	case 1:
		global.PreviewID = form.Class.ID
		err = method.PreviewClass{}.Watch(cookies)
		if err != nil {
			return
		}
	case 2:
		global.CourseID = form.Class.ID
		err = method.FreeClass{}.Watch(cookies)
		if err != nil {
			return
		}
	default:
		fmt.Println(text.WrongInput)
	}

	return
}
