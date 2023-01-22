package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rizqonmaulana/go-restapi-gin/controllers/productController"
	"github.com/rizqonmaulana/go-restapi-gin/models"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()

	r.GET("/api/products", productController.Index)
	r.GET("/api/product/:id", productController.Show)
	r.POST("/api/product", productController.Create)
	r.PUT("/api/product/:id", productController.Update)
	r.DELETE("/api/product", productController.Delete)

	r.Run()
}