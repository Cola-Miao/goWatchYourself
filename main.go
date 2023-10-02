package main

import (
	"goWatchYourself/global"
	"goWatchYourself/initialize"
	"goWatchYourself/utils"
	"log"
	"os/exec"
)

func main() {
	utils.Version()
	initialize.Initialize()
	log.Println(exec.Command(`cmd`, `/c`, `start`, global.DefaultAddr).Start())
	log.Fatalln(global.Engine.Run(global.Addr))
}
