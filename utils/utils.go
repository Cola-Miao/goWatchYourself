package utils

import (
	"fmt"
	"goWatchYourself/global"
	"net/http"
	"time"
)

func WaitTime() {
	time.Sleep(global.WaitTime * time.Millisecond)
}

func SetHeader(req *http.Request) {
	req.Header.Add("User-Agent", global.UserAgent)
	req.Header.Add("Accept", global.Accept)
	req.Header.Add("Accept-Language", global.AcceptLanguage)
	req.Header.Add("Content-Type", global.ContentType)
}

func Version() {
	fmt.Println("Version: ", global.Version)
}

func WaitAnyKey() {
	fmt.Scanln()
}
