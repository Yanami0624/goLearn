package mygin

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	
	v1 := e.Group("")
	{
		v1.GET("/hello", Hello)
	}
	v2 := e.Group("")
	{
		v2.GET("/login", Login)
		v2.GET("/finduser/:username", FindUser)
		v2.GET("/userpage/*filepath", UserPage)
	}
	v3 := e.Group("") 
	{
		v3.POST("/upload", Upload)
		v3.GET("/download/:filename", Download)
	}
	e.Run("127.0.0.1:8080")
}
