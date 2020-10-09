package controlles

import (
	"batara/src/helpers/env"
	"batara/src/helpers/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)


func GetToken(c *gin.Context){
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute*4).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString([]byte(env.GetEnv("API_KEY")))
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	tokenR, error := sign.SignedString([]byte(env.GetEnv("API_REFRESH")))
	if err != nil || error != nil {
		response.ResponseJson(c,201,err)
	}

	response.ResponseJson(c,200,gin.H{"api_token":tokenstring,"token_refresh":tokenR})
}