package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

//import "net/http"

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("api/v1/date", func(c *gin.Context) {
		t := time.Now()

		c.JSON(200, gin.H{
			"Date": t.Format(time.RFC3339),
		})
	})
	r.Run(":7700") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
