package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		searchedPhrase := c.PostForm("searchedPhrase")
		if searchedPhrase == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No variable in cookies files"})

			return
		}
		filter := bson.M{
			"$or": []bson.M{
				{"product_name": bson.M{"$regex": primitive.Regex{Pattern: searchedPhrase, Options: ""}}},
			},
		}

		var products []string
		cursor, err := m.DatabaseCollection.Find(context.Background(), filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var result bson.M
			err := cursor.Decode(&result)
			if err != nil {
				log.Printf("Error decoding document: %v\n", err)
				continue
			}

			// Extract product_name field from the document and append it to the products slice
			productName, ok := result["product_name"].(string)
			if ok {
				products = append(products, productName)
			}
		}

		// Sending data as JSON response
		c.JSON(http.StatusOK, products)
	})
}
