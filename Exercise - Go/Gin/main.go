package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode) //modalità di esecuzione di Gin
	r := gin.Default()         //nuovo router

	r.GET("/example", func(c *gin.Context) { //definisce una rotta
		c.String(http.StatusOK, fmt.Sprintln("Demo endpoint ..."))
	})

	err := r.Run(":8080") //avvia il server sulla porta 8080
	if err != nil {
		panic(err)
	}
}
