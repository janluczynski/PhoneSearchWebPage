package scrapers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"main.go/commons"
	mongodb "main.go/mongoDB"
)

func XkomScrapProductInfo() {
	err := godotenv.Load("C:/Users/lepar/VSdev/Projekt3rok/backend/.env")
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
		phone := xkomScrapHelper(link)
		update := bson.M{
			"$set": bson.M{
				"brand":     phone.Brand,
				"model":     phone.Model,
				"image":     phone.ImageURL,
				"price":     phone.Price,
				"processor": phone.Processor,
				"ram":       phone.RAM,
				"storage":   phone.Storage,
				"battery":   phone.Battery,
				"display":   phone.Display,
			},
		}
		_, err := m.ProductCollection.UpdateOne(context.Background(), bson.M{"product_url": link}, update)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func xkomScrapHelper(baseURL string) commons.Product {
	c := colly.NewCollector()
	//baseURL := "https://www.x-kom.pl/p/1160711-smartfon-telefon-samsung-galaxy-a23-4-128gb-black-25w-120hz.html"

	var Specification []string
	var Product commons.Product

	c.OnHTML(".sc-13p5mv-2.fxqQxb .sc-1s1zksu-0.sc-1s1zksu-1.hHQkLn.sc-13p5mv-0.VGBov", func(e *colly.HTMLElement) {
		Specification = append(Specification, e.Text)
	})
	c.OnHTML(".sc-1bker4h-10.kHPtVn h1", func(e *colly.HTMLElement) {
		Product.Brand = strings.Split(e.Text, " ")[0]
		Product.Model = strings.Join(strings.Split(e.Text, " ")[1:], " ")
	})
	c.OnHTML(".sc-n4n86h-1.hYfBFq", func(e *colly.HTMLElement) {
		Price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(e.Text, ",00 zł", ""), " ", ""), 32)
		if err != nil {
			fmt.Println(err)
		}
		Product.Price = float32(Price)
	})
	c.OnHTML(".sc-1tblmgq-0.sc-1tblmgq-3.cIswgX.sc-jiiyfe-2.jGSlBb img", func(e *colly.HTMLElement) {
		Product.ImageURL = e.Attr("src")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		for _, element := range Specification {
			if strings.Contains(element, "Procesor") {
				Product.Processor = strings.ReplaceAll(element, "Procesor", "")
			} else if strings.Contains(element, "Pamięć RAM") {
				ram := strings.ReplaceAll(element, "Pamięć RAM", "")
				if strings.Contains(ram, "GB") {
					ram = strings.ReplaceAll(ram, " GB", "")
					ramInt, err := strconv.Atoi(ram)
					if err != nil {
						fmt.Println(err)
					}
					Product.RAM = ramInt * 1024
				} else if strings.Contains(ram, "MB") {
					ram = strings.ReplaceAll(ram, " MB", "")
					ramInt, err := strconv.Atoi(ram)
					if err != nil {
						fmt.Println(err)
					}
					Product.RAM = ramInt * 1024
				}
			} else if strings.Contains(element, "Pamięć wbudowana") {
				Storage := strings.ReplaceAll(element, "Pamięć wbudowana", "")
				if strings.Contains(Storage, "GB") {
					Storage = strings.ReplaceAll(Storage, " GB", "")
					StorageInt, err := strconv.Atoi(Storage)
					if err != nil {
						fmt.Println(err)
					}
					Product.Storage = StorageInt * 1024
				} else if strings.Contains(Storage, "MB") {
					Storage = strings.ReplaceAll(Storage, " MB", "")
					StorageInt, err := strconv.Atoi(Storage)
					if err != nil {
						fmt.Println(err)
					}
					Product.Storage = StorageInt * 1024
				}
			} else if strings.Contains(element, "Pojemność baterii") {
				BatteryInt, err := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(element, "Pojemność baterii", ""), " mAh", ""))
				if err != nil {
					fmt.Println(err)
				}
				Product.Battery = BatteryInt
			} else if strings.Contains(element, "Przekątna ekranu") {
				Product.Display = strings.ReplaceAll(strings.ReplaceAll(element, "Przekątna ekranu", ""), ",", ".")
			}
		}
	})

	c.Visit(baseURL)

	return Product
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
