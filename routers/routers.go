package routers

import (
	"financekeeper-api/model"
	"time"
)

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

func Routers()  {
	r := gin.Default()
	r.Use(cors.Default())
	v1 := r.Group("/api/v1/")
	v1.GET("date", func(c *gin.Context) {
		t := time.Now()
		c.JSON(200, gin.H{
			"Date": t.Format(time.RFC3339),
		})
	})
	{
		v1.POST("action", model.CreateAction)
		v1.GET("action", model.FetchAllAction)
		//v1.GET("action/:id", model.FetchSingleAction)
		//v1.PUT("action/:id", model.UpdateAction)
		//v1.DELETE("action/:id", model.DeleteAction)
	}





	r.Run(":7700")
}
