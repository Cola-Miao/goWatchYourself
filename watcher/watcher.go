package watcher

import (
	"fmt"
	"goWatchYourself/api"
	"goWatchYourself/global"
	"goWatchYourself/method"
	"goWatchYourself/text"
)

func Watcher() {
	input := api.GetUserInput()
	cookies, err := method.Login(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch global.Model {
	case 1:
		method.PreviewClass{}.Watch(cookies)
	case 2:
		method.FreeClass{}.Watch(cookies)
	default:
		fmt.Println(text.WrongInput)
	}
}
