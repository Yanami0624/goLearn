package mygin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	username string `binding:"required" json:"username" form:"username" uri:"username"`
	password string `binding:"required" json:"password" form:"password" uri:"password"`
}

func Login(ctx *gin.Context) {
	var login LoginUser
	if ctx.ShouldBind(&login) != nil && login.username != "" {
		ctx.String(http.StatusOK, "login successfully")
	} else {
		ctx.String(http.StatusBadRequest, "login failed")
	}
}

func FindUser(ctx *gin.Context) {
	username := ctx.Param("username")
	ctx.String(http.StatusOK, "found user %s", username)
}

func UserPage(ctx *gin.Context) {
	filepath := ctx.Param("filepath")
	ctx.String(http.StatusOK, "filapath is %s", filepath)
}
