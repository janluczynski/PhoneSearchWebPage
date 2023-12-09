package scrapers

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

func MediaMarktScrapProductInfo() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Some error occured. Err: %s \n", err)
	}
	m, err := mongodb.InitDB()
	if err != nil {
		fmt.Println("Error:", err)
	}

	filter := bson.M{"product_url": bson.M{"$regex": "https://mediamarkt.pl/"}}

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
		phoneInfo := mediaMarktScrapHelper(link)
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

func mediaMarktScrapHelper(baseURL string) []string {
	c := colly.NewCollector()
	//baseURL := "https://mediamarkt.pl/telefony-i-smartfony/smartfon-samsung-galaxy-a14-lte-4-128gb-czarny-sm-a145rzkveue"

	phoneSpecification := make([]string, 0)
	phoneInfo := make([]string, 0)
	fullProductInfo := make([]string, 0)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.OnHTML("span.product-show-specification-item span", func(e *colly.HTMLElement) {
		element := e.Text
		dividedPhoneSpecification := strings.Split(element, "\n")
		for _, specification := range dividedPhoneSpecification {
			phoneSpecification = append(phoneSpecification, specification)
		}

	})
	c.OnHTML("h1.title.is-heading", func(e *colly.HTMLElement) {
		element := e.Text
		fullPhoneInfo := strings.Split(element, " ")

		Brand := fullPhoneInfo[0]
		Model := strings.Join(fullPhoneInfo[1:], " ")
		phoneInfo = append(phoneInfo, Brand, Model)

	})
	c.OnHTML("div.main-price.is-big span.whole", func(e *colly.HTMLElement) {
		element := e.Text
		phonePrice := strings.Fields(element)[0]
		phoneInfo = append(phoneInfo, phonePrice)
	})

	c.OnScraped(func(r *colly.Response) {
		bracketsReegx := regexp.MustCompile(`\([^)]*\)`)
		Brand := phoneInfo[0]
		Model := bracketsReegx.ReplaceAllString(phoneInfo[1], "")
		Price := ""
		if len(phoneInfo) == 2 {
			Price = "N/A"
		} else {
			Price = phoneInfo[2] + " zł"
		}
		Display := ""
		Procesor := ""
		RAM := ""
		Storage := ""
		Battery := ""

		Inches := ""
		Hertz := ""
		fmt.Println("Finished", r.Request.URL)

		for i := 0; i < len(phoneSpecification); i++ {
			if strings.HasPrefix(phoneSpecification[i], "Rozmiar przekątnej") {
				Inches = strings.TrimLeft(phoneSpecification[i+1], " ") + `"`
			} else if strings.HasPrefix(phoneSpecification[i], "Częstotliwość odświeżania") {
				Hertz = phoneSpecification[i+1] + "Hz"
			} else if strings.HasPrefix(phoneSpecification[i], "Pamięć wewnętrzna") {
				Storage = strings.TrimLeft(phoneSpecification[i+1], " ")
			} else if strings.HasPrefix(phoneSpecification[i], "Informuje o ilość") {
				RAM = strings.TrimLeft(phoneSpecification[i+1], " ")
			} else if strings.HasPrefix(phoneSpecification[i], "Określa nazwę i model") {
				Procesor = strings.TrimLeft(phoneSpecification[i+1], " ")
			} else if strings.HasPrefix(phoneSpecification[i], "Informuje o pojemności akumulatora") {
				Battery = strings.TrimLeft(phoneSpecification[i+1], " ")
			}
		}
		if Inches != "" && Hertz != "" {
			Display = fmt.Sprintf("%s,%s", Inches, Hertz)
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
		if Battery != "" {
			Battery = Battery + " mAh"
		}

		fullProductInfo = append(fullProductInfo, Brand, Model, Price, Procesor, RAM, Storage, Battery, Display)
	})
	c.Visit(baseURL)

	return fullProductInfo
}

func FakeMediaMarktRequest() {
	c := colly.NewCollector()
	baseURL := "https://mediamarkt.pl/telefony-i-smartfony/smartfon-samsung-galaxy-s23-8-256gb-czarny-sm-s916bzkdeue"
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.Visit(baseURL)
}
