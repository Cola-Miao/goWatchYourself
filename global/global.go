package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Version = "0.65"

	UserAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/117.0"
	Accept         = "application/json, text/plain, */*"
	AcceptLanguage = "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2"
	ContentType    = "application/x-www-form-urlencoded"

	WaitTime = 2

	Addr        = "localhost:5912"
	DefaultAddr = "http://localhost:5912/index"
)

var (
	Engine    *gin.Engine
	Client    *http.Client
	CourseID  string
	PreviewID string
)
