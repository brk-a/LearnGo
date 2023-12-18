package main

import (
	"github.com/gin-gonic/gin"
	"src/go-crud/initialisers"
)

func init()  {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDb()
}

func main()  {
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "success"
		})
	})
	r.Run()
}