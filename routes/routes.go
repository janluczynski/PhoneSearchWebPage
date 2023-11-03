package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	commons "main.go/commons"
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
func SearchProductsFromSearchBar(r *gin.Engine, m *mongodb.MongoDB) {
	r.POST("/search", func(c *gin.Context) {
		searchedPhrase, err := c.Cookie("searchedPhrase")
		if err != nil || searchedPhrase == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No variable in cookies files"})

			return
		}
		filter := bson.M{
			"$or": []bson.M{
				{"product_name": bson.M{"$regex": primitive.Regex{Pattern: searchedPhrase, Options: ""}}},
			},
		}

		var products []commons.Product
		cursor, err := m.DatabaseCollection.Find(context.Background(), filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		defer cursor.Close(context.Background())

		err = cursor.All(context.Background(), &products)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		// Sending data as JSON response
		c.JSON(http.StatusOK, products)
	})
}
