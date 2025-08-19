package main

import (
	"fmt"
	"jwt-go/initializers"

	"github.com/gin-gonic/gin"
)



func init(){
	initializers.LoadEnvVariables()
}

func main(){
	fmt.Println("hello")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}