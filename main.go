package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type mine struct {
	Name string `json:"name"`
}

func main() {
	router := gin.Default()
	router.GET("/", hello)
	router.Run("localhost:3000")
}

func hello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello!!!")
}
