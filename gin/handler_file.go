package mygin

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "%+v", err)
		return
	}
	err = ctx.SaveUploadedFile(file, "./"+file.Filename)
	if err != nil {
		ctx.String(http.StatusBadRequest, "%+v", err)
		return
	}
	ctx.String(http.StatusOK, "upload %s, total %d bytes", file.Filename, file.Size)
}

func Download(ctx *gin.Context) {
	filename := filepath.Base(ctx.Param("filename"))
	ctx.FileAttachment(filepath.Join("static", filename), filename)
}
