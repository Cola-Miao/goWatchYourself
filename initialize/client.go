package initialize

import (
	"goWatchYourself/global"
	"net/http"
	"time"
)

func initClient() {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	global.Client = c
}
