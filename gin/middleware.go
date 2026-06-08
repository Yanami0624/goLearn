package mygin

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeMiddleware() gin.HandlerFunc {
	const header string = "[time MW]"
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)
		fmt.Println(header, "request cost", duration)
	}
}