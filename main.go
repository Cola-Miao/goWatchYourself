package main

import (
	"goWatchYourself/global"
	"goWatchYourself/initialize"
	"goWatchYourself/utils"
	"log"
)

func main() {
	var err error

	if err = initialize.Initialize(); err != nil {
		log.Fatalln(err)
	}
	go func() {
		if err = utils.OpenBrowser(); err != nil {
			log.Println(err)
		}
	}()
	if err = global.Engine.Run(global.Addr); err != nil {
		log.Fatalln(err)
	}
}
