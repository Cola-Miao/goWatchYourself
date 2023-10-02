package initialize

import (
	"embed"
	"github.com/gin-gonic/gin"
	"goWatchYourself/global"
	"goWatchYourself/models"
	"goWatchYourself/watcher"
	"html/template"
	"net/http"
)

//go:embed template/*
var fs embed.FS

func initEngine() {
	e := gin.Default()
	tpl := template.Must(template.New("").ParseFS(fs, "template/*.gohtml"))
	e.SetHTMLTemplate(tpl)
	//e.LoadHTMLGlob("template/*.gohtml")

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
