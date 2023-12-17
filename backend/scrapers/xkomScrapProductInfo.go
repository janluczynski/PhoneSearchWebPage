package scrapers

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"main.go/commons"
	mongodb "main.go/mongoDB"
)

func XkomScrapProductInfo() {
	err := godotenv.Load("../.env")
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
				"name":      phone.Brand + " " + phone.Model,
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
	//baseURL := "https://www.x-kom.pl/p/1079444-smartfon-telefon-asus-rog-phone-6d-12g-256g-space-grey.html"

	var Specification []string
	var Product commons.Product

	c.OnHTML(".sc-13p5mv-2.fxqQxb .sc-1s1zksu-0.sc-1s1zksu-1.hHQkLn.sc-13p5mv-0.VGBov", func(e *colly.HTMLElement) {
		Specification = append(Specification, e.Text)
	})
	c.OnHTML(".sc-1bker4h-10.kHPtVn h1", func(e *colly.HTMLElement) {
		PhoneName := strings.Split(e.Text, " ")
		GBRegex := regexp.MustCompile(`(GB|1TB)$`)
		Product.Brand = PhoneName[0]
		Model := ""
		for i := 1; i < len(PhoneName); i++ {
			if GBRegex.MatchString(PhoneName[i]) || strings.Contains(PhoneName[i], "5G") || strings.Contains(PhoneName[i], "/") {
				break
			} else {
				Model += PhoneName[i] + " "
			}
		}
		Product.Model = Model
	})
	c.OnHTML(".sc-n4n86h-1.hYfBFq", func(e *colly.HTMLElement) {
		Price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(e.Text, ",00 zł", ""), " ", ""), 64)
		if err != nil {
			fmt.Println(err)
		}
		Product.Price = Price
	})
	c.OnHTML(".sc-1tblmgq-0.sc-1tblmgq-3.cIswgX.sc-jiiyfe-2.jGSlBb img", func(e *colly.HTMLElement) {
		Product.ImageURL = e.Attr("src")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		for _, specification := range Specification {
			if strings.Contains(specification, "Procesor") {
				Product.Processor = strings.ReplaceAll(specification, "Procesor", "")
			} else if strings.Contains(specification, "Pamięć RAM") {
				ram := strings.ReplaceAll(specification, "Pamięć RAM", "")
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
			} else if strings.Contains(specification, "Pamięć wbudowana") {
				Storage := strings.ReplaceAll(specification, "Pamięć wbudowana", "")
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
				} else if strings.Contains(Storage, "TB") {
					Storage = strings.ReplaceAll(Storage, " TB", "")
					StorageInt, err := strconv.Atoi(Storage)
					if err != nil {
						fmt.Println(err)
					}
					Product.Storage = StorageInt * 1024 * 1024
				}
			} else if strings.Contains(specification, "Pojemność baterii") {
				BatteryInt, err := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(specification, "Pojemność baterii", ""), " mAh", ""))
				if err != nil {
					fmt.Println(err)
				}
				Product.Battery = BatteryInt
			} else if strings.Contains(specification, "Przekątna ekranu") {
				Product.Display = strings.ReplaceAll(strings.ReplaceAll(specification, "Przekątna ekranu", ""), ",", ".")
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
