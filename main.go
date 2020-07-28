package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
//import "net/http"
import "time"


func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("api/v1/date", func(c *gin.Context) {
		t:= time.Now()

		c.JSON(200, gin.H{
			"Date": t.Format(time.RFC3339),
		})
	})
	r.Run(":7500") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
