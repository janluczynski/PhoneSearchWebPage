package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mongodb "main.go/mongoDB"
)

func PostProductInfo(r *gin.Engine, m *mongodb.MongoDB) {
	r.POST("/parse/product", func(c *gin.Context) {
		productID, err := c.Cookie("productID")
		if err != nil || productID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No variable in cookies files"})

			return
		}
		// Checking if product exists in DB
		if !m.CheckIfProductInDB(productID) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Worker doesn't exist"})

			return
		}

		// Getting product data from DB
		productData, err := m.GetProductData(productID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting product data"})
			return
		}

		// Sending data as JSON response
		c.JSON(http.StatusOK, productData)
	})
}
