package mygin

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.Static("/static", "./static")

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
	v3 := e.Group("", TimeMiddleware())
	{
		v3.POST("/upload", Upload)
		v3.GET("/download/:filename", Download)
	}
	v4 := e.Group("")
	{
		v4.GET("/cookie/:username/:password", Cookie)
		v4.GET("/jwt/sign/:username/:password", SignJWT)
		v4.GET("/jwt/parse", ParseJWT)
	}
	e.Run("127.0.0.1:8080")
}
