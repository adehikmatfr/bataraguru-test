package response

import "github.com/gin-gonic/gin"

func ResponseJson(c *gin.Context, code int, message interface{}) {
	if code>200{
		c.AbortWithStatusJSON(code, gin.H{"status_code":code,"error": message})
	}else{
		c.AbortWithStatusJSON(code, gin.H{"status_code":code,"data": message})
	}
}
