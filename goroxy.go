package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func getHelloAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, authorized!"})
}

func main() {
	r := gin.Default()

	r.GET("/", getHello)

	accounts := gin.Accounts{
		"goran": "password",
		"lahi":  "password",
	}

	authorized := r.Group("/", gin.BasicAuth(accounts))
	authorized.GET("/auth", getHelloAuth)

	r.Run()
}
