package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//package od metod związanych z łącznością z bazą danych

type MongoDB struct {
	Client            *mongo.Client
	ProductCollection *mongo.Collection
}

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

	clientOptions := options.Client().ApplyURI(uri)

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
func (m *MongoDB) AddProducts(product []interface{}) error {
	_, err := m.ProductCollection.InsertMany(context.Background(), product)
	if err != nil {
		return err
	}
	return nil
}
func (m *MongoDB) DeleteAllProducts() error {
	// Specify an empty filter to match all documents
	filter := bson.M{} // bson.M{} represents an empty BSON document

	// Perform the deletion operation
	result, err := m.ProductCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	// Output the number of deleted documents
	fmt.Printf("Deleted %v products.\n", result.DeletedCount)
	return nil
}
func (m *MongoDB) AddProduct(product interface{}) error {
	_, err := m.ProductCollection.InsertOne(context.Background(), product)
	if err != nil {
		return err
	}
	return nil
}
