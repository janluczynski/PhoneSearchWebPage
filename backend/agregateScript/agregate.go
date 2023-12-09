package main

import (
	"context"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main.go/commons"
	mongodb "main.go/mongoDB"
)

//TODO
//make collection of phones
//get data from one and compare to others and get them into threes
//make a map of them

type Phone struct {
	Products []commons.Product `json:"products" bson:"products"`
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Some error occured. Err: %s \n", err)
	}
	m, err := mongodb.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{
		"product_url": bson.M{"$regex": primitive.Regex{Pattern: "x-kom", Options: ""}},
	}
	cur, err := m.ProductCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var products []commons.Product

	err = cur.All(context.Background(), &products)
	if err != nil {
		log.Fatal(err)
	}

	var sortedProducts [][]commons.Product

	for _, product := range products {
		filter := bson.M{
			"brand": bson.M{"$regex": primitive.Regex{Pattern: product.Brand, Options: "i"}},
			"model": bson.M{"$regex": primitive.Regex{Pattern: strings.Split(product.Model, " ")[0] + ".*" + strings.Split(product.Model, " ")[1], Options: "i"}},
			// "ram":       bson.M{"$regex": primitive.Regex{Pattern: strings.Split(product.RAM, " ")[0] + ".*", Options: "i"}},
			// "storage":   bson.M{"$regex": primitive.Regex{Pattern: strings.Split(product.Storage, " ")[0] + ".*", Options: "i"}},
			// "battery":   bson.M{"$regex": primitive.Regex{Pattern: strings.Split(product.Battery, " ")[0] + ".*", Options: "i"}},
			"processor": bson.M{"$regex": primitive.Regex{Pattern: strings.Split(product.Processor, " ")[0] + ".*" + strings.Split(product.Processor, " ")[1], Options: "i"}},
			"display":   bson.M{"$regex": primitive.Regex{Pattern: strings.ReplaceAll(strings.Split(product.Display, " ")[0], ",", "") + ".*", Options: "i"}},
		}
		cur, err := m.ProductCollection.Find(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(context.Background())

		var products2 []commons.Product

		err = cur.All(context.Background(), &products2)
		if err != nil {
			log.Fatal(err)
		}

		sortedProducts = append(sortedProducts, products2)
	}

	var phones []Phone

	for _, product := range sortedProducts {
		var phone Phone
		phone.Products = product
		phones = append(phones, phone)
	}

	for _, phone := range phones {
		_, err = m.PhoneCollection.InsertOne(context.Background(), phone)
		if err != nil {
			log.Fatal(err)
		}
	}
}
