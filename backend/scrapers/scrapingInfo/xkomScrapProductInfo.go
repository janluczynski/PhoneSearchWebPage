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

func XkomScrapProductInfo() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Some error occured. Err: %s \n", err)
	}
	m, err := mongodb.InitDB()
	if err != nil {
		fmt.Println("Error:", err)
	}

	filter := bson.M{"product_url": bson.M{"$regex": "https://www.x-kom.pl/p"}}

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
		phoneInfo := xkomScrapHelper(link)
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
func xkomScrapHelper(baseURL string) []string {
	c := colly.NewCollector()
	// baseURL := "https://www.x-kom.pl/p/1160711-smartfon-telefon-samsung-galaxy-a23-4-128gb-black-25w-120hz.html"

	phoneSpecification := make([]string, 0)
	phoneInfo := make([]string, 0)
	fullProductInfo := make([]string, 0)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.OnHTML(".sc-13p5mv-3", func(e *colly.HTMLElement) { // scraping specyfication
		element := e.Text
		dividedPhoneSpecification := strings.Split(element, "\n")
		for _, specification := range dividedPhoneSpecification {
			phoneSpecification = append(phoneSpecification, specification)
		}
	})
	c.OnHTML(".sc-1bker4h-10.kHPtVn", func(e *colly.HTMLElement) { //scraping brand + model
		element := e.Text
		fullPhoneInfo := strings.Split(element, " ")

		Brand := fullPhoneInfo[0]
		Model := strings.Join(fullPhoneInfo[1:], " ")
		phoneInfo = append(phoneInfo, Brand, Model)

	})
	c.OnHTML(".sc-n4n86h-1.hYfBFq", func(e *colly.HTMLElement) { //scraping price
		phoneInfo = append(phoneInfo, e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		bracketsReegx := regexp.MustCompile(`\([^)]*\)`)
		Brand := phoneInfo[0]
		Model := bracketsReegx.ReplaceAllString(phoneInfo[1], "")
		Price := phoneInfo[2]
		Display := ""
		Procesor := phoneSpecification[0]
		RAM := ""
		Storage := ""
		Battery := ""
		Inches := ""
		Hertz := ""

		fmt.Println("Finished", r.Request.URL)
		ramRegex := regexp.MustCompile(`^\b([1-9]|1\d|2[0-9])\b GB$`)
		storageRegex := regexp.MustCompile(`^(32|[4-9][0-9]|[1-4][0-9][0-9]|5[0-1][0-2]) GB$|1 TB$`)
		inchesRegex := regexp.MustCompile(`^\d,\d{1,2}"$`)
		hertznRegex := regexp.MustCompile(`^\d{2,3} Hz$`)
		batteryRegex := regexp.MustCompile(`^\d{4} mAh$`)
		for _, element := range phoneSpecification {

			if ramRegex.MatchString(element) {
				if RAM == "" {
					RAM = element
				}
			} else if storageRegex.MatchString(element) {
				if Storage == "" {
					Storage = element
				}
			} else if inchesRegex.MatchString(element) {
				if Inches == "" {
					Inches = strings.ReplaceAll(element, ",", ".")
				}
			} else if hertznRegex.MatchString(element) {
				if Hertz == "" {
					Hertz = element
				}
			} else if batteryRegex.MatchString(element) {
				if Battery == "" {
					Battery = element
				}
			}
		}
		if Inches != "" && Hertz != "" {
			Display = fmt.Sprintf("%s, %s", Inches, Hertz)
		} else if Inches != "" && Hertz == "" {
			Display = Inches
		} else {
			Display = "N/A"
		}
		if RAM == "" {
			RAM = "N/A"
		}
		if Battery == "" {
			Battery = "N/A"
		}

		fullProductInfo = append(fullProductInfo, Brand, Model, Price, Procesor, RAM, Storage, Battery, Display)
	})

	c.Visit(baseURL)

	return fullProductInfo
}

func FakeXKomRequest() {
	c := colly.NewCollector()
	baseURL := "https://www.x-kom.pl/p/1165774-smartfon-telefon-nokia-2660-4g-flip-rozowy-stacja-ladujaca.html"
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.Visit(baseURL)
}
