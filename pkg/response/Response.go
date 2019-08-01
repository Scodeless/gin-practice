package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Response struct {
	G *gin.Context
	Time time.Time
}

func (res *Response) Response(code int64, message string, data interface{}) {
	fmt.Println(res.Time)
	res.G.JSON(200, gin.H{
		"Runtime": time.Since(res.Time).Seconds(),
		"Code": code,
		"Error": message,
		"Data": data,
	})
}