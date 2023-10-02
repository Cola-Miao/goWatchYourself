package initialize

import (
	"log"
	"os"
)

func initLog() (err error) {
	fp, err := os.OpenFile("gws.log", os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	log.SetOutput(fp)
	log.SetFlags(log.Llongfile)

	return
}
