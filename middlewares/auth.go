package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gteruithi.com/demo-rest-api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	_, userId, err := utils.VerifyToken(token)

	if err != nil {
		fmt.Println(err)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	context.Set("userId", userId)
	context.Next()
}
