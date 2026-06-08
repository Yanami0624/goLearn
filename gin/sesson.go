package mygin

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Cookie(ctx *gin.Context) {
	username := ctx.Param("username")
	password := ctx.Param("password")
	cookie, err := ctx.Cookie(username)
	if err != nil {
		cookie = "Not set"
		ctx.SetCookie(username, password, 30, "/", "", false, true)
	}
	fmt.Println("cookie value: ", cookie)
}

const jwtSecret = "gocodes-dev-secret"

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func SignJWT(ctx *gin.Context) {
	username := ctx.Param("username")
	password := ctx.Param("password")
	if username == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	claims := JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "gocodes",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(jwtSecret))

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func ParseJWT(ctx *gin.Context) {
	tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
	if tokenString == "" {
		tokenString = ctx.Query("token")
	}
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
		return
	}

	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"username":   claims.Username,
		"issuer":     claims.Issuer,
		"expires_at": claims.ExpiresAt.Time.Format(time.RFC3339),
	})
}
