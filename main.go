package main

import (
	"goWatchYourself/initialize"
	"goWatchYourself/utils"
	"goWatchYourself/watcher"
)

func main() {
	utils.Version()
	initialize.InitDefault()
	watcher.Watcher()

	utils.WaitAnyKey()
}
