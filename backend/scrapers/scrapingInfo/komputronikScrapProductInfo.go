package scrappers

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"main.go/commons"
	mongodb "main.go/mongoDB"
)

func KomputronikScrapProductInfo() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Some error occured. Err: %s \n", err)
	}
	m, err := mongodb.InitDB()
	if err != nil {
		fmt.Println("Error:", err)
	}

	filter := bson.M{"product_url": bson.M{"$regex": "https://www.komputronik.pl/"}}

	var products []commons.Product

	cur, err := m.ProductCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &products); err != nil {
		log.Fatal(err)
	}
	for _, product := range products {
		link := product.ProductURL
		phoneInfo := komputronikScrapHelper(link) // CHANGE TO KOMPUTRONIKSCRAP HELPER
		update := bson.M{
			"$set": bson.M{
				"brand":      phoneInfo[0],
				"model":      phoneInfo[1],
				"sale_price": phoneInfo[2],
				"processor":  phoneInfo[3],
				"ram":        phoneInfo[4],
				"storage":    phoneInfo[5],
				"battery":    phoneInfo[6],
				"display":    phoneInfo[7],
			},
		}
		_, err := m.ProductCollection.UpdateOne(context.Background(), bson.M{"product_url": link}, update)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func komputronikScrapHelper(baseURL string) []string {
	c := colly.NewCollector()
	//baseURL := "https://www.komputronik.pl/product/826982/poco-f5-pro-12-256gb-czarny.html"

	phoneSpecification := make([]string, 0)
	phoneInfo := make([]string, 0)
	fullProductInfo := make([]string, 0)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.OnHTML("div.p-1 p", func(e *colly.HTMLElement) { //scraping specification
		elements := e.Text
		dividePhoneSpecification := strings.Split(strings.ReplaceAll(elements, " ", ""), "\n")
		for _, specification := range dividePhoneSpecification {
			phoneSpecification = append(phoneSpecification, specification)
		}

	})
	c.OnHTML("h1.tests-product-name", func(e *colly.HTMLElement) { // scraping brand + model
		element := e.Text
		fullPhoneInfo := strings.Fields(element)

		Brand := fullPhoneInfo[0]
		Model := strings.Join(fullPhoneInfo[1:], " ")
		phoneInfo = append(phoneInfo, Brand, Model)

	})
	c.OnHTML("div.font-bold.leading-8.text-3xl", func(e *colly.HTMLElement) { //scraping price
		phoneInfo = append(phoneInfo, e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		bracketsReegx := regexp.MustCompile(`\([^)]*\)`)
		Brand := phoneInfo[0]
		Model := bracketsReegx.ReplaceAllString(phoneInfo[1], "")
		Price := phoneInfo[2]
		Procesor := ""
		RAM := ""
		Storage := ""
		Battery := ""
		Inches := ""
		fmt.Println("Finished", r.Request.URL)
		storageRegex := regexp.MustCompile(`GB$`)
		inchesRegex := regexp.MustCompile(`^\d.\d{1,2}cale$`)
		batteryRegex := regexp.MustCompile(`^\d{4}mAh$`)

		for i := 0; i < len(phoneSpecification); i++ {
			if storageRegex.MatchString(phoneSpecification[i]) {
				if Storage == "" {
					Storage = phoneSpecification[i]
				} else if RAM == "" {
					RAM = phoneSpecification[i]
					if storageRegex.MatchString(phoneSpecification[i+2]) {
						Procesor = phoneSpecification[i+4]
					} else {
						Procesor = phoneSpecification[i+2]
					}
				}
			} else if inchesRegex.MatchString(phoneSpecification[i]) {
				if Inches == "" {
					Inches = strings.ReplaceAll(phoneSpecification[i], "cale", `"`)
				}
			} else if batteryRegex.MatchString(phoneSpecification[i]) {
				if Battery == "" {
					Battery = phoneSpecification[i]
				}
			}
		}
		if Inches == "" {
			Inches = "N/A"
		}
		if RAM == "" {
			RAM = "N/A"
		}
		if Battery == "" {
			Battery = "N/A"
		}
		if Procesor == "" {
			Procesor = "N/A"
		}

		fullProductInfo = append(fullProductInfo, Brand, Model, Price, Procesor, RAM, Storage, Battery, Inches)
	})
	c.Visit(baseURL)

	return fullProductInfo
}
