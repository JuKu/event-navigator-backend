package middlewares

import (
	"github.com/JuKu/event-navigator-backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token found."})
		return
	}

	// remove "Bearer ", if neccessary
	token = strings.Replace(token, "Bearer ", "", 1)

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
