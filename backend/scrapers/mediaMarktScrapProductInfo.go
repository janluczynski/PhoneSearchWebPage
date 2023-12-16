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
		phone := mediaMarktScrapHelper(link)
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

func mediaMarktScrapHelper(baseURL string) commons.Product {
	c := colly.NewCollector()
	//baseURL := "https://mediamarkt.pl/telefony-i-smartfony/smartfon-samsung-galaxy-a14-lte-4-128gb-czarny-sm-a145rzkveue"

	var Specification []string
	var Product commons.Product

	c.OnHTML("div.product-menu-specification li.attribute", func(e *colly.HTMLElement) {
		Specification = append(Specification, e.Text)
	})
	c.OnHTML("h1.title.is-heading", func(e *colly.HTMLElement) {
		PhoneName := strings.Split(e.Text, " ")
		GBRegex := regexp.MustCompile(`(GB|1TB)$`)
		Product.Brand = PhoneName[1]
		Model := ""
		for i := 2; i < len(PhoneName); i++ {
			if GBRegex.MatchString(PhoneName[i]) || strings.Contains(PhoneName[i], "5G") || strings.Contains(PhoneName[i], "LTE") {
				break
			} else {
				Model += PhoneName[i] + " "
			}
		}
		Product.Model = Model
	})
	c.OnHTML("div.price-box span.whole", func(e *colly.HTMLElement) {
		Price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(e.Text, "\n", ""), " ", ""), 64)
		if err != nil {
			fmt.Println(err)
		}
		Product.Price = Price
	})
	c.OnHTML(".spark-image.image img.is-loaded", func(e *colly.HTMLElement) {
		Product.ImageURL = e.Attr("src")
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		for _, specification := range Specification {
			if strings.Contains(specification, "Model procesora") {
				Product.Processor = strings.ReplaceAll(specification, " Model procesora   Określa nazwę i model procesora/układu SoC.    ", "")
			} else if strings.Contains(specification, "Pamięć RAM") {
				ram := strings.ReplaceAll(specification, " Pamięć RAM   Informuje o ilość pamięci RAM.    ", "")
				if strings.Contains(ram, "GB") {
					ram = strings.ReplaceAll(ram, " GB ", "")
					ramInt, err := strconv.Atoi(ram)
					if err != nil {
						fmt.Println(err)
					}
					Product.RAM = ramInt * 1024
				}
			} else if strings.Contains(specification, "Pamięć wbudowana") {
				Storage := strings.ReplaceAll(specification, " Pamięć wbudowana   Pamięć wewnętrzna jest to wbudowana pamięć przeznaczona do zapisywania danych użytkownika. Im więcej pamięci tym więcej aplikacji i danych można zapisać. Wielkość pamięci dostępnej dla użytkownika może być mniejsza ze względu na zainstalowany system i aplikacje.    ", "")
				if strings.Contains(Storage, "GB") {
					Storage = strings.ReplaceAll(Storage, " GB ", "")
					StorageInt, err := strconv.Atoi(Storage)
					if err != nil {
						fmt.Println(err)
					}
					Product.Storage = StorageInt * 1024
				} else if strings.Contains(Storage, "TB") {
					Storage = strings.ReplaceAll(Storage, " TB ", "")
					StorageInt, err := strconv.Atoi(Storage)
					if err != nil {
						fmt.Println(err)
					}
					Product.Storage = StorageInt * 1024 * 1024
				}
			} else if strings.Contains(specification, "Pojemność [mAh]") {
				BatteryInt, err := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(specification, " Pojemność [mAh]   Informuje o pojemności akumulatora zastosowanego w telefonie. Wartość podawana w miliamperogodzinach.    ", ""), " ", ""))
				if err != nil {
					fmt.Println(err)
				}
				Product.Battery = BatteryInt
			} else if strings.Contains(specification, "Przekątna ekranu") {
				Product.Display = strings.ReplaceAll(strings.ReplaceAll(specification, " Przekątna ekranu [cal]   Rozmiar przekątnej ekranu podawany w calach. Im większa wartość (przekątna) tym większy i bardziej szczegółowy obraz.    ", ""), " ", "") + `"`
			}
		}
	})
	c.Visit(baseURL)

	return Product
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
