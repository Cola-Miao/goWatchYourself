package utils

import (
	"goWatchYourself/global"
	"net/http"
	"time"
)

func RandomSleep() {
	t := global.Rand.Intn(1500)
	time.Sleep(time.Duration(t) * time.Millisecond)
}

func SetHeader(req *http.Request) {
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/117.0")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}
