package initialize

import (
	"github.com/gin-gonic/gin"
	"goWatchYourself/global"
	"goWatchYourself/models"
	"goWatchYourself/watcher"
	"net/http"
)

func initEngine() {
	e := gin.Default()

	e.LoadHTMLGlob("template/*")

	v1 := e.Group("/")
	{
		v1.GET("/index", renderIndex)
		v1.POST("/index", formHandle)
	}

	global.Engine = e
}

func outputErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"error": err.Error(),
	})
}

func formHandle(c *gin.Context) {
	var form models.Form
	var err error

	if err = c.ShouldBind(&form); err != nil {
		outputErr(c, err)
		return
	}
	if err = watcher.Watcher(&form); err != nil {
		outputErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful(?)",
	})
}

func renderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", nil)
}
