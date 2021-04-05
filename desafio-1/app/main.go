package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()
	routes.Static("/", "./public")
	routes.Run()
}
