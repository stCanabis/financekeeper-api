package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=15204352 sslmode=disable")
	if err != nil {
		fmt.Printf("db.Prepare error: %v\n",err)
		panic("failed to connect database")
	}
	return db

}

type Action struct {
	gorm.Model
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
}

type TransformedAction struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
}

func autoMigration() {
	db := Database()
	db.AutoMigrate(&Action{})
}

func CreateAction(c *gin.Context) {
	autoMigration()

	var action Action
	c.BindJSON(&action)

	db := Database()
	db.Save(&action)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Action item created successfully!", "resourceId": action.ID})
}
func FetchAllAction(c *gin.Context) {
	autoMigration()
	var action []Action
	var _action []TransformedAction

	db := Database()
	db.Find(&action)

	if len(action) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range action {
		_action = append(
			_action, TransformedAction{
				ID:     item.ID,
				Name:   item.Name,
				Ticker: item.Ticker,
			})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _action})
}

//func FetchSingleProduct(c *gin.Context) {
//	autoMigration()
//	var product Product
//	productId := c.Param("id")
//
//	db := Database()
//	db.First(&product, productId)
//
//	if product.ID == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})
//		return
//	}
//	_product := TransformedProduct{
//		ID:          product.ID,
//		Name:        product.Name,
//		Description: product.Description,
//		Images:      product.Images,
//		Price:       product.Price,
//	}
//	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _product})
//}
//
//func UpdateProduct(c *gin.Context) {
//	autoMigration()
//	var product Product
//	tproductId := c.Param("id")
//	db := Database()
//	db.First(&product, tproductId)
//
//	if product.ID == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})
//		return
//	}
//
//	db.Model(&product).Update("name", c.PostForm("name"))
//	db.Model(&product).Update("description", c.PostForm("description"))
//	db.Model(&product).Update("images", c.PostForm("images"))
//	db.Model(&product).Update("price", c.PostForm("price"))
//	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Product updated successfully!"})
//}
//
//func DeleteProduct(c *gin.Context) {
//	autoMigration()
//	var product Product
//	productId := c.Param("id")
//	db := Database()
//	db.First(&product, productId)
//
//	if product.ID == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product found!"})
//		return
//	}
//
//	db.Delete(&product)
//	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "product deleted successfully!"})
//}
