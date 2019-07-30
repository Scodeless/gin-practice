package response

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Response struct {
	G *gin.Context
}

func (res *Response) Response(code int64, message string, data interface{}) {
	res.G.JSON(200, gin.H{
		"Runtime": time.Since(time.Now()).Seconds(),
		"Code": code,
		"Error": message,
		"Data": data,
	})
}