package scrapers

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"main.go/commons"
	mongodb "main.go/mongoDB"
)

func KomputronikScrapProductInfo() {
	err := godotenv.Load("../.env")
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
		phone := komputronikScrapHelper(link)
		update := bson.M{
			"$set": bson.M{
				"name":      phone.Brand + " " + phone.Model,
				"site_name": phone.SiteName,
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
		time.Sleep(1 * time.Second)
	}
}
func komputronikScrapHelper(baseURL string) commons.Product {
	c := colly.NewCollector()
	//baseURL := "https://www.komputronik.pl/product/860865/telefon-apple-iphone-15-pro-max-1tb-tytan-czarny.html"

	var Specification []string
	var Product commons.Product
	var StringPrice []string

	c.OnHTML(".tests-full-specification.wrap-text .grid.grid-cols-2.text-sm", func(e *colly.HTMLElement) {
		Specification = append(Specification, e.Text)
	})
	c.OnHTML("div.inline-flex.items-center.mt-2.flex-wrap div.font-bold.leading-8.text-3xl", func(e *colly.HTMLElement) {
		StringPrice = append(StringPrice, e.Text)
	})
	c.OnHTML("div.overflow-hidden.flex.justify-center.items-center.w-80.h-80 img", func(e *colly.HTMLElement) {
		Product.ImageURL = e.Attr("src")
	})
	c.OnHTML("h1.tests-product-name.font-headline.text-lg.font-bold.leading-8.line-clamp-2", func(e *colly.HTMLElement) {
		trimRegex := regexp.MustCompile(` {2,}`)
		PhoneName := strings.Split(strings.ReplaceAll(trimRegex.ReplaceAllString(e.Text, ""), "\n", ""), " ")
		Product.SiteName = strings.Join(PhoneName, " ")
		GBRegex := regexp.MustCompile(`(GB|1TB)$`)
		Product.Brand = PhoneName[0]
		Model := ""
		for i := 1; i < len(PhoneName); i++ {
			if GBRegex.MatchString(PhoneName[i]) || strings.Contains(PhoneName[i], "5G") {
				break
			} else {
				Model += PhoneName[i] + " "
			}
		}
		Product.Model = Model
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		trimRegex := regexp.MustCompile(` {2,}`)
		Price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(StringPrice[0], ",", "."), " zł", ""), "\u00a0", ""), 64)
		if err != nil {
			fmt.Println(err)
		}
		Product.Price = Price
		for _, specification := range Specification {
			specification = trimRegex.ReplaceAllString(strings.ReplaceAll(specification, "\n", ""), "")
			if strings.Contains(specification, "Zastosowany procesor") {
				Product.Processor = strings.ReplaceAll(specification, "Zastosowany procesor", "")
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
					Product.RAM = ramInt
				}
			} else if strings.Contains(specification, "Pamięć Flash") {
				Storage := strings.ReplaceAll(specification, "Pamięć Flash", "")
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
			} else if strings.Contains(specification, "Pojemność akumulatora") {
				BatteryInt, err := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(specification, "Pojemność akumulatora", ""), " mAh", ""))
				if err != nil {
					fmt.Println(err)
				}
				Product.Battery = BatteryInt
			} else if strings.Contains(specification, "Przekątna wyświetlacza") {
				Product.Display = strings.ReplaceAll(strings.ReplaceAll(specification, "Przekątna wyświetlacza", ""), " cale", "\"")
			}
		}
	})

	c.Visit(baseURL)

	return Product
}
func Test1(baseURL string) {
	c := colly.NewCollector()
	//baseURL := "https://www.komputronik.pl/product/860865/telefon-apple-iphone-15-pro-max-1tb-tytan-czarny.html"

	var Specification []string
	var Product commons.Product
	var StringPrice []string

	c.OnHTML(".tests-full-specification.wrap-text .grid.grid-cols-2.text-sm", func(e *colly.HTMLElement) {
		Specification = append(Specification, e.Text)
	})
	c.OnHTML("div.inline-flex.items-center.mt-2.flex-wrap div.font-bold.leading-8.text-3xl", func(e *colly.HTMLElement) {
		StringPrice = append(StringPrice, e.Text)
	})
	c.OnHTML("div.overflow-hidden.flex.justify-center.items-center.w-80.h-80 img", func(e *colly.HTMLElement) {
		Product.ImageURL = e.Attr("src")
	})
	c.OnHTML("h1.tests-product-name.font-headline.text-lg.font-bold.leading-8.line-clamp-2", func(e *colly.HTMLElement) {
		trimRegex := regexp.MustCompile(` {2,}`)
		PhoneName := strings.Split(strings.ReplaceAll(trimRegex.ReplaceAllString(e.Text, ""), "\n", ""), " ")
		Product.SiteName = strings.Join(PhoneName, " ")
		GBRegex := regexp.MustCompile(`(GB|1TB)$`)
		Product.Brand = PhoneName[0]
		Model := ""
		for i := 1; i < len(PhoneName); i++ {
			if GBRegex.MatchString(PhoneName[i]) || strings.Contains(PhoneName[i], "5G") {
				break
			} else {
				Model += PhoneName[i] + " "
			}
		}
		Product.Model = Model
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		trimRegex := regexp.MustCompile(` {2,}`)
		Price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(StringPrice[0], ",", "."), " zł", ""), "\u00a0", ""), 64)
		if err != nil {
			fmt.Println(err)
		}
		Product.Price = Price
		for _, specification := range Specification {
			specification = trimRegex.ReplaceAllString(strings.ReplaceAll(specification, "\n", ""), "")
			if strings.Contains(specification, "Zastosowany procesor") {
				Product.Processor = strings.ReplaceAll(specification, "Zastosowany procesor", "")
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
					Product.RAM = ramInt
				}
			} else if strings.Contains(specification, "Pamięć Flash") {
				Storage := strings.ReplaceAll(specification, "Pamięć Flash", "")
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
			} else if strings.Contains(specification, "Pojemność akumulatora") {
				BatteryInt, err := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(specification, "Pojemność akumulatora", ""), " mAh", ""))
				if err != nil {
					fmt.Println(err)
				}
				Product.Battery = BatteryInt
			} else if strings.Contains(specification, "Przekątna wyświetlacza") {
				Product.Display = strings.ReplaceAll(strings.ReplaceAll(specification, "Przekątna wyświetlacza", ""), " cale", "\"")
			}
		}
	})

	c.Visit(baseURL)

	fmt.Println(Product)
}
