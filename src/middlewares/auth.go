package middlewares

import (
	"batara/src/helpers/env"
	"batara/src/helpers/response"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
)


func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := env.GetEnv("API_KEY")
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}
	return func(c *gin.Context) {
		reqToken := c.GetHeader("Authorization")
		if reqToken == "" {
			response.ResponseJson(c, 401, "API token required")
			return
		}
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		//verify
		token, timeOut := RequireTokenAuthentication(reqToken)
		if token && timeOut > 0 {
			c.Next()
		}else{
			response.ResponseJson(c, 401, "not authorized")
			return
		}
	}
}

func RequireTokenAuthentication(tokenstring string) (bool, int) {
	token, _ := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return []byte(env.GetEnv("API_KEY")), nil
	})
	return token.Valid, getTokenRemainingValidity(token.Claims.(jwt.MapClaims)["exp"])
}

func getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Before(time.Now())
		if !remainer  {
			return 1
		}
	}
	return -1
}

