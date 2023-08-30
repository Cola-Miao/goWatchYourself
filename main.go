package main

import (
	"fmt"
	"goWatchYourself/initialize"
	"goWatchYourself/watcher"
)

func main() {
	initialize.InitDefault()
	watcher.Watcher()

	fmt.Scanln()
}
