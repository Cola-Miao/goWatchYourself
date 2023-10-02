package initialize

import "github.com/gin-gonic/gin"

func initEngine() {
	e := gin.Default()
	e.GET("/")
	e.POST("/")
}
