package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})

	r.POST("/callback", func(c *gin.Context) {
		b, _ := ioutil.ReadAll(c.Request.Body)
		c.JSON(http.StatusOK, gin.H{
			"received": string(b),
		})
	})

	r.Run(":8080") // listen and serve on localhost:8080
}
