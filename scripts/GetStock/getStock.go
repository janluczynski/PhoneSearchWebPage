package main

import (
	"context"
	"log"
	"math/rand"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	commons "main.go/commons"
	mongodb "main.go/mongoDB"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Some error occured. Err: %s \n", err)
	}
	// Initialize the database connection
	m, err := mongodb.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	var products []commons.Product
	var productStocks []commons.ProductStock

	cursor, err := m.DatabaseCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product commons.Product
		if err := cursor.Decode(&product); err != nil {
			log.Println(err)
			continue
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	for _, product := range products {
		var productStock commons.ProductStock
		productStock.ProductID = product.ProductID
		productStock.Count = rand.Intn(15) + 10
		productStocks = append(productStocks, productStock)
	}
	var stockList []interface{}
	for _, productStock := range productStocks {
		stockList = append(stockList, productStock)
	}
	m.Client.Database("Projekt").Collection("stock").DeleteMany(context.Background(), bson.M{})
	m.Client.Database("Projekt").Collection("stock").InsertMany(context.Background(), stockList)
}
