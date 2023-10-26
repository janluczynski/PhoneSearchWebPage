package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//package od metod związanych z łącznością z bazą danych

type MongoDB struct {
	Client            *mongo.Client
	ProductCollection *mongo.Collection
}

func CreateIfNotExists(d *mongo.Database, collection string) (*mongo.Collection, error) {
	collectionObj := d.Collection(collection)

	collectionExists, err := collectionObj.EstimatedDocumentCount(context.Background())
	if err != nil {
		return nil, err
	}

	if collectionExists == 0 {
		err = d.CreateCollection(context.Background(), collection)
		if err != nil {
			return nil, err
		}
	}

	return collectionObj, nil
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
func (m *MongoDB) AddProducts(products []interface{}) error {
	_, err := m.ProductCollection.InsertMany(context.Background(), products)
	if err != nil {
		return err
	}
	return nil
}
