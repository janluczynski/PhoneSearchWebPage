package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	commons "main.go/commons"
	mongodb "main.go/mongoDB"
	scrapper "main.go/photoscrapper"
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
	m.DeleteAllProducts()

	// Call the function to get database items from JSON file
	err = getDatabaseItems(m, "../Dataset/data.json")
	if err != nil {
		log.Fatal(err)
	}

	// // Insert products into MongoDB collection
	// err = m.AddProducts(products)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func getDatabaseItems(m *mongodb.MongoDB, filePath string) error {
	// Read the JSON file into a byte slice
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var products []commons.Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return err
	}

	u := launcher.New().Bin("./bin/chrome").MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()

	defer browser.MustClose()

	// Add UUID and images to each product
	for i := range products {
		var product interface{}
		products[i].ProductID = uuid.New().String()
		//TODO: add scrapping images function to scram images from the web and add them to a []string and apped to the products
		products[i].Imagetab = scrapper.Scrap(browser, products[i].ProductURL)
		product = products[i]
		m.AddProduct(product)
	}

	return nil
}
