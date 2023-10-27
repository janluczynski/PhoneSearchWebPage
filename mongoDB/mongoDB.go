package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"main.go/commons"
)

//package od metod związanych z łącznością z bazą danych

type MongoDB struct {
	Client            *mongo.Client
	ProductCollection *mongo.Collection
}

// func check if collection exists and create it if not
func CreateIfNotExists(db *mongo.Database, collectionName string) (*mongo.Collection, error) {
	collection := db.Collection(collectionName)

	if err := db.CreateCollection(context.Background(), collectionName); err != nil {
		if _, ok := err.(mongo.CommandError); !ok {
			return nil, err
		}
	}

	return collection, nil
}

func InitDB() (*MongoDB, error) {
	uri := os.Getenv("DBURL")
	if uri == "" {
		return nil, fmt.Errorf("DBURL env is empty")
	}

	database := os.Getenv("DBNAME")
	if database == "" {
		return nil, fmt.Errorf("DBNAME env is empty")
	}
	clientOptions := options.Client().ApplyURI("mongodb+srv://projektzespolowy73:esFPWrGpjtdsYkCM@projekt.cch4qp1.mongodb.net/")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	db := client.Database(database)

	productCollection, err := CreateIfNotExists(db, "product")
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		Client:            client,
		ProductCollection: productCollection,
	}, nil
}

// function to check if there is product with given ID in database
func (m *MongoDB) CheckIfProductInDB(productID string) bool {
	var product commons.Product
	err := m.ProductCollection.FindOne(context.Background(), commons.Product{ProductID: productID}).Decode(&product)
	if err != nil {
		return false
	}
	return true
}

// function to get product data from database
func (m *MongoDB) GetProductData(productID string) (commons.Product, error) {
	var product commons.Product
	err := m.ProductCollection.FindOne(context.Background(), commons.Product{ProductID: productID}).Decode(&product)
	if err != nil {
		return commons.Product{}, err
	}
	return product, nil
}
