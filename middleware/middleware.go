package middleware

import (
	"fmt"
	"net/http"
	"os"
	"plugspot/db"
	"plugspot/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func MiddlewareForAllRole(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": tokenString})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || float64(time.Now().Unix()) > claims["exp"].(float64) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "your token was expired"})
		return
	}

	var user model.UserAccount
	db.Connection.First(&user, claims["userid"])

	if user.ID == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "not match with any user id"})
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}

func MiddlewareForCustomerRole(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": tokenString})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || float64(time.Now().Unix()) > claims["exp"].(float64) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "your token was expired"})
		return
	}
	if !ok || claims["role"] != "customer" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "your role is not customer"})
		return
	}

	var user model.UserAccount
	db.Connection.First(&user, claims["userid"])

	if user.ID == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "not match with any user id"})
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}

func MiddlewareForProviderRole(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": tokenString})
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || float64(time.Now().Unix()) > claims["exp"].(float64) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "your token was expired"})
		return
	}
	if !ok || claims["role"] != "provider" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "your role is not provider"})
		return
	}

	var user model.UserAccount
	db.Connection.First(&user, claims["userid"])

	if user.ID == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "not match with any user id"})
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}
