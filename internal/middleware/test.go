package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func TestMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Printf("test middle ware")
		context.Abort()
		return
	}
}
