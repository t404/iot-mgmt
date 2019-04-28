package middlewares

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Evaluation() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("EVA Allow")
		c.Next()
	}
}