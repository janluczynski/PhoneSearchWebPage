package mongodb

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"main.go/commons"
)

//package od metod związanych z łącznością z bazą danych

type MongoDB struct {
	Client            *mongo.Client
	ProductCollection *mongo.Collection
	PhoneCollection   *mongo.Collection
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

	phoneCollection, err := CreateIfNotExists(db, "phone")
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		Client:            client,
		ProductCollection: productCollection,
		PhoneCollection:   phoneCollection,
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

// function to check if there is product with given ID in database
func (m *MongoDB) CheckIfProductInDB(productID string) bool {
	var product commons.Product

	filter := bson.M{"product_id": productID}

	err := m.ProductCollection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return false
	}
	return true
}

// function to get product data from database
func (m *MongoDB) GetProductData(productID string) (commons.Product, error) {
	var product commons.Product

	filter := bson.M{"product_id": productID}

	err := m.ProductCollection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return commons.Product{}, err
	}

	return product, nil
}

// function to get product data from database
func (m *MongoDB) GetSameProductData(productID string) (map[string][]interface{}, error) {
	var product commons.Product

	filter := bson.M{"product_id": productID}

	err := m.ProductCollection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return nil, err
	}

	patternBrand := bson.M{"$regex": primitive.Regex{Pattern: product.Brand, Options: "i"}}
	filterSame := bson.M{"brand": patternBrand, "model": product.Model, "ram": product.RAM, "storage": product.Storage}

	var sameProducts []commons.Product

	cursor, err := m.ProductCollection.Find(context.Background(), filterSame)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &sameProducts)
	if err != nil {
		return nil, err
	}

	// creating map with product_id as key and [link,price] as value
	sameProductsMap := make(map[string][]interface{})
	for _, sameProduct := range sameProducts {
		pattern := `//([^/]+)`
		re := regexp.MustCompile(pattern)

		match := re.FindStringSubmatch(sameProduct.ProductURL)

		sameProductsMap[sameProduct.ProductID] = []interface{}{sameProduct.ProductURL, match[1], sameProduct.Price}
	}

	return sameProductsMap, nil
}

func (m *MongoDB) GetProductsByBrandOrModel(searchedPhrase, sortByField string, sortOrder int) ([]commons.Product, error) {
	filter := bson.M{"name": primitive.Regex{Pattern: searchedPhrase, Options: "i"}}

	var products []commons.Product

	options := options.Find()
	options.SetSort(bson.D{{Key: sortByField, Value: sortOrder}})

	cursor, err := m.ProductCollection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &products)
	if err != nil {

		return nil, err
	}

	return products, nil
}
func (m *MongoDB) GetProductsWithoutSorting(searchedPhrase string) ([]commons.Product, error) {
	filter := bson.M{"name": primitive.Regex{Pattern: searchedPhrase, Options: "i"}}

	var products []commons.Product

	cursor, err := m.ProductCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &products)
	if err != nil {

		return nil, err
	}

	return products, nil
}
func (m *MongoDB) FindSimilarPhones(name string, ram, storage int) ([]commons.Product, error) {
	filter := bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}}, "ram": ram, "storage": storage}
	var products []commons.Product

	cursor, err := m.ProductCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &products)
	if err != nil {

		return nil, err
	}

	return products, nil
}
