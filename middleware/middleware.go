package middleware

import (
	"Girl/utlis"
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoginStatus(c *gin.Context) {
	fmt.Printf("utlis.MD5Encrypt(\"123456\"): %v\n", utlis.MD5Encrypt("123456"))
}
