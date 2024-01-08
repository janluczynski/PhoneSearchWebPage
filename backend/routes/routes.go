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
	SortBy         string `json:"sortBy"`
	Order          int    `json:"order"` //1 ascending, -1 descending
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Empty search phrase"})
			return
		}

		products, err := m.GetProductsByBrandOrModel(phrase.SearchedPhrase, phrase.SortBy, phrase.Order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting product data"})
			return
		}

		// Sending data as JSON response
		c.JSON(http.StatusOK, products)
	})
}
func SearchProducts(r *gin.Engine, m *mongodb.MongoDB) {
	r.POST("/searchbar", func(c *gin.Context) {

		type searchedPhrase struct {
			SearchedPhrase string `json:"searchedPhrase"`
		}
		var phrase searchedPhrase
		err := c.BindJSON(&phrase)

		if err != nil || phrase.SearchedPhrase == "" {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Empty search phrase"})
			return
		}

		products, err := m.GetProductsWithoutSorting(phrase.SearchedPhrase)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting product data"})
			return
		}

		// Sending data as JSON response
		c.JSON(http.StatusOK, products)
	})
}
func GetSimilarProducts(r *gin.Engine, m *mongodb.MongoDB) {
	r.POST("/similar", func(c *gin.Context) {

		type similarProduct struct {
			Name    string `json:"name" bson:"name"`
			Ram     int    `json:"ram" bson:"ram"`
			Storage int    `json:"storage" bson:"storage"`
		}

		var productInfo similarProduct

		if err := c.ShouldBindJSON(&productInfo); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}

		products, err := m.FindSimilarPhones(productInfo.Name, productInfo.Ram, productInfo.Storage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting product data"})
			return
		}

		// Sending data as JSON response
		c.JSON(http.StatusOK, products)
	})
}
