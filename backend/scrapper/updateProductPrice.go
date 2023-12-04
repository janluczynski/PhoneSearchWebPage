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

var c = colly.NewCollector()

func UpdateProductPrice() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Some error occured. Err: %s \n", err)
	}
	m, err := mongodb.InitDB()
	if err != nil {
		fmt.Println("Error:", err)
	}

	filter := bson.M{"product_url": bson.M{"$regex": "https://www.komputronik.pl/product/"}}

	var products []commons.Product

	cur, err := m.ProductCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &products); err != nil {
		log.Fatal(err)
	}
	xkomRegex := regexp.MustCompile(`https://www.x-kom.pl/p/\d+`)
	mediaMarktRegex := regexp.MustCompile(`https://mediamarkt.pl/telefony-i-smartfony/\d+`)
	komputronikRegex := regexp.MustCompile(`https://www.komputronik.pl/product/\d+`)
	for _, product := range products {
		link := product.ProductURL
		price := product.Price
		updatedPrice := ""
		if xkomRegex.MatchString(link) {

		} else if mediaMarktRegex.MatchString(link) {
			updatedPrice = mediaMarktPriceScraper(link)
		} else if komputronikRegex.MatchString(link) {
			var t3 []string
			c.OnHTML("div.font-bold.leading-8.text-3xl", func(e *colly.HTMLElement) {
				elemnt := e.Text
				t1 := strings.ReplaceAll(elemnt, " ", "")
				t2 := strings.Split(t1, "\n")
				t3 = append(t3, t2[0])
			})
			c.Visit(link)
			c.OnScraped(func(r *colly.Response) {
				fmt.Println("Visiting: ", link, "Price: ", t3)
				fmt.Println("Updated price: ", t3[0])
			})
			updatedPrice = t3[0]
		}
		if updatedPrice != price {
			update := bson.M{"$set": bson.M{"sale_price": updatedPrice}}
			_, err := m.ProductCollection.UpdateOne(context.Background(), bson.M{"product_url": link}, update)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
func xkomPriceScraper(link string) string {
	var updatedPrice string
	c.OnHTML(".sc-n4n86h-1.hYfBFq", func(e *colly.HTMLElement) {
		updatedPrice = e.Text
	})
	c.Visit(link)
	return updatedPrice
}
func mediaMarktPriceScraper(link string) string {
	var updatedPrice string
	c.OnHTML("div.main-price.is-big span.whole", func(e *colly.HTMLElement) {
		elemnt := e.Text
		re := regexp.MustCompile(`\b\d+\b`)
		matches := re.FindAllString(elemnt, 1)
		updatedPrice = matches[0] + " zł"
	})
	c.Visit(link)
	return updatedPrice
}