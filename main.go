package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func getHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello world"})
}

func posts(c *gin.Context) {
	// var u User
	var u User
	c.BindJSON(&u)
	c.JSON(http.StatusOK, gin.H{"user": u})
}

func putpathParameter(c *gin.Context) {
	user := c.Param("user")
	c.JSON(http.StatusOK, gin.H{"message": user})
}

func deleteQuery(c *gin.Context) {
	user := c.Query("user")
	c.JSON(http.StatusOK, gin.H{"message": user})
}

func main() {
	router := gin.Default()
	router.GET("/", getHelloWorld)
	router.POST("/", posts)
	router.PUT("/", putpathParameter)
	router.DELETE("/", deleteQuery)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
