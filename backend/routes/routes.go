package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	r.GET("/parse/product", func(c *gin.Context) {
		ProductID := c.Query("product_id")

		if ProductID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No variable in data"})
			return
		}
		// Checking if product exists in DB
		if !m.CheckIfProductInDB(ProductID) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product doesn't exist"})

			return
		}

		// Getting product data from DB
		productData, err := m.GetProductData(ProductID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting product data"})
			return
		}

		// Sending data as JSON response
		c.JSON(http.StatusOK, productData)
	})
}

func GetSamePhones(r *gin.Engine, m *mongodb.MongoDB) {

	r.GET("/same/product", func(c *gin.Context) {
		ProductID := c.Query("product_id")

		if ProductID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No variable in data"})
			return
		}
		// Checking if product exists in DB
		if !m.CheckIfProductInDB(ProductID) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product doesn't exist"})

			return
		}

		// Getting product data from DB
		sameProducts, err := m.GetSameProductData(ProductID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting product data"})
			return
		}

		// Sending data as JSON response
		c.JSON(http.StatusOK, sameProducts)
	})
}
func SearchProductsFromSearchBar(r *gin.Engine, m *mongodb.MongoDB) {
	r.GET("/search", func(c *gin.Context) {

		searchedPhrase := c.Query("searchedPhrase")
		sortBy := c.Query("sortBy")
		orderInt, err := strconv.Atoi(c.Query("order"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error with converting order to int"})
			return
		}
		value, err := strconv.Atoi(c.Query("value"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error with converting value to int"})
			return
		}
		if searchedPhrase == "" || sortBy == "" || orderInt != 1 && orderInt != -1 || value < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Empty search phrase"})
			return
		}

		products, err := m.GetProductsByBrandOrModel(searchedPhrase, sortBy, orderInt, value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting product data"})
			return
		}

		// Sending data as JSON response
		c.JSON(http.StatusOK, products)
	})
}
func SearchProducts(r *gin.Engine, m *mongodb.MongoDB) {
	r.GET("/searchbar", func(c *gin.Context) {

		name := c.Query("name")

		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Empty search phrase"})
			return
		}

		products, err := m.GetProductsWithoutSorting(name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting product data"})
			return
		}

		// Sending data as JSON response
		c.JSON(http.StatusOK, products)
	})
}
func GetSimilarProducts(r *gin.Engine, m *mongodb.MongoDB) {
	r.GET("/similar", func(c *gin.Context) {

		name := c.Query("name")
		ram, err := strconv.Atoi(c.Query("ram"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error with converting ram to int"})
			return
		}
		storage, err := strconv.Atoi(c.Query("storage"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error with converting storage to int"})
			return
		}
		products, err := m.FindSimilarPhones(name, ram, storage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting product data"})
			return
		}
		// Sending data as JSON response
		c.JSON(http.StatusOK, products)
	})
}
func IncrementField(r *gin.Engine, m *mongodb.MongoDB) {
	r.PUT("/increment", func(c *gin.Context) {
		id := c.Query("id") // get the id from the query parameters
		println(id)
		filter := bson.M{"product_id": id}                // filter the documents by id
		update := bson.M{"$inc": bson.M{"popularity": 1}} // increment yourField by 1

		// update the document
		_, err := m.ProductCollection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating the field"})
			return
		}

		// Sending response
		c.JSON(http.StatusOK, gin.H{"message": "Field incremented successfully"})
	})
}
func GetTopProducts(r *gin.Engine, m *mongodb.MongoDB) {
	r.GET("/top-products", func(c *gin.Context) {
		// Define the sort option
		sortOption := options.Find()
		sortOption.SetSort(bson.D{{"popularity", -1}}) // sort by popularity in descending order
		sortOption.SetLimit(3)                         // limit the result to 3 documents

		// Find the top 3 products
		cursor, err := m.ProductCollection.Find(context.TODO(), bson.D{{}}, sortOption)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting top products"})
			return
		}

		var products []bson.M
		if err = cursor.All(context.TODO(), &products); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding top products"})
			return
		}

		// Sending response
		c.JSON(http.StatusOK, products)
	})
}
