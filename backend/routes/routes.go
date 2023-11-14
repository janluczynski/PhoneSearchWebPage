package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	mongodb "main.go/mongoDB"
)

type productID struct {
	ProductID string `json:"product_id"`
}

type sphrase struct {
	SearchedPhrase string `json:"searchedPhrase"`
}

func PostProductInfo(r *gin.Engine, m *mongodb.MongoDB) {
	r.POST("/parse/product", func(c *gin.Context) {
		var product productID
		err := c.BindJSON(&product)

		if err != nil || product.ProductID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No variable in data"})

			return
		}
		// Checking if product exists in DB
		if !m.CheckIfProductInDB(product.ProductID) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product doesn't exist"})

			return
		}

		// Getting product data from DB
		productData, err := m.GetProductData(product.ProductID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting product data"})
			return
		}

		// Sending data as JSON response
		c.JSON(http.StatusOK, productData)
	})
}
func SearchProductsFromSearchBar(r *gin.Engine, m *mongodb.MongoDB) {
	r.POST("/search", func(c *gin.Context) {

		var phrase sphrase
		err := c.BindJSON(&phrase)

		if err != nil || phrase.SearchedPhrase == "" {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "No variable in cookies files"})
			return
		}

		products, err := m.GetProductsByBrandOrModel(phrase.SearchedPhrase)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting product data"})
			return
		}

		log.Println(products)
		// Sending data as JSON response
		c.JSON(http.StatusOK, products)
	})
}
