package response

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Response struct {
}

func (res *Response) Response(g *gin.Context, code int64, message string, data interface{}) {
	g.JSON(200, gin.H{
		"Runtime": time.Since(time.Now()).Seconds(),
		"Code": code,
		"Error": message,
		"Data": data,
	})
}