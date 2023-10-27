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
	scrap "main.go/photoscrapper"
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
	products, err := getDatabaseItems("../Dataset/data.json")
	if err != nil {
		log.Fatal(err)
	}

	// Insert products into MongoDB collection
	err = m.AddProducts(products)
	if err != nil {
		log.Fatal(err)
	}
}

func getDatabaseItems(filePath string) ([]interface{}, error) {
	// Read the JSON file into a byte slice
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var products []commons.Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}

	u := launcher.New().Bin("./bin/chrome").MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()

	defer browser.MustClose()

	// Add UUID to each product
	for i := range products {
		products[i].ProductID = uuid.New().String()
		//TODO: add scrapping images function to scram images from the web and add them to a []string and apped to the products
		products[i].Imagetab = scrap.Scrap(browser, products[i].ProductURL)
	}

	// Convert products to []interface{}
	var interfaceProducts []interface{}
	for _, p := range products {
		interfaceProducts = append(interfaceProducts, p)
	}

	return interfaceProducts, nil

}

// func printProduct(product commons.Product) {
// 	fmt.Println("Product ID:", product.ProductID)
// 	fmt.Println("Product Name:", product.ProductName)
// 	fmt.Println("Brand:", product.Brand)
// 	fmt.Println("Image URL:", product.ImageURL)
// 	fmt.Println("Sale Price:", product.SalePrice)
// 	fmt.Println("Colour:", product.Colour)
// 	fmt.Println("Description:", product.Description)
// 	fmt.Println("Category:", product.Category)
// 	fmt.Println("Material:", product.Material)
// }
