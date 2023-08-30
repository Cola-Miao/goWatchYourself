package main

import (
	"fmt"
	"goWatchYourself/api"
	"goWatchYourself/global"
	"goWatchYourself/initialize"
	"goWatchYourself/method"
)

func main() {
	initialize.InitDefault()
	input := api.GetUserInput()
	cookies, err := method.Login(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	videos, err := method.GetVideos(global.CourseID)
	if err != nil {
		fmt.Println(err)
		return
	}
	method.ReplayAttack(videos, cookies)
}
