package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
	fmt.Println("ok")
}

func Little(c *gin.Context) {
	c.HTML(http.StatusOK, "little.html", nil)

}

func Leaving(c *gin.Context) {
	c.HTML(http.StatusOK, "little.html", nil)

}

func Photo(c *gin.Context) {
	c.HTML(http.StatusOK, "little.html", nil)

}

func TodoList(c *gin.Context) {
	c.HTML(http.StatusOK, "little.html", nil)

}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "little.html", nil)

}
