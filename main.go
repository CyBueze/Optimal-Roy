package main

import (
	"github.com/CyBueze/org/views"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		views.InventoryPage().Render(c.Request.Context(), c.Writer)
	})
	
	r.GET("/sales", func(c *gin.Context) {
		views.Calculator().Render(c.Request.Context(), c.Writer)
	})
	
	r.GET("/cart", func(c *gin.Context){
	  views.CartPage().Render(c.Request.Context(), c.Writer)
	})

	r.Run(":8080")
}