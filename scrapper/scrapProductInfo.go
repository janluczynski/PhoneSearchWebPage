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

	cur, err := m.DatabaseCollection.Find(context.Background(), filter)
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
		_, err := m.DatabaseCollection.UpdateOne(context.Background(), bson.M{"product_url": link}, update)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func xkomScrapHelper(baseURL string) []string {
	c := colly.NewCollector()
	// baseURL := "https://www.x-kom.pl/p/1160711-smartfon-telefon-samsung-galaxy-a23-4-128gb-black-25w-120hz.html"

	phoneElements := make([]string, 0)
	phoneInfo := make([]string, 0)

	fullProductInfo := make([]string, 0)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.OnHTML(".sc-13p5mv-3", func(e *colly.HTMLElement) {
		element := e.Text
		// fmt.Println(element)
		test := strings.Split(element, "\n")
		for _, element := range test {
			phoneElements = append(phoneElements, element)
		}
	})
	c.OnHTML(".sc-1bker4h-10.kHPtVn", func(e *colly.HTMLElement) {
		fullPhone := e.Text
		fullPhoneInfo := strings.Split(fullPhone, " ")

		Brand := fullPhoneInfo[0]
		Model := ""
		for i := 1; i < len(fullPhoneInfo); i++ {
			Model = Model + fullPhoneInfo[i] + " "
		}
		phoneInfo = append(phoneInfo, Brand)
		phoneInfo = append(phoneInfo, Model)

	})
	c.OnHTML(".sc-n4n86h-1.hYfBFq", func(e *colly.HTMLElement) {
		phoneInfo = append(phoneInfo, e.Text)
	})

	c.OnHTML(".sc-13p5mv-3.UfEQd", func(e *colly.HTMLElement) {
		phoneElements = append(phoneElements, e.Text)
	})
	c.OnScraped(func(r *colly.Response) {
		Brand := phoneInfo[0]
		Model := phoneInfo[1]
		Price := phoneInfo[2]
		Display := ""
		Procesor := phoneElements[0]
		RAM := ""
		Storage := ""
		Battery := ""

		Cale := ""
		Herce := ""

		fmt.Println("Finished", r.Request.URL)
		ramRegex := regexp.MustCompile(`^\b([1-9]|1\d|2[0-9])\b GB$`)
		storageRegex := regexp.MustCompile(`^(32|[4-9][0-9]|[1-4][0-9][0-9]|5[0-1][0-2]) GB$|1 TB$`)
		caleRegex := regexp.MustCompile(`^\d,\d{1,2}"$`)
		hertznRegex := regexp.MustCompile(`^\d{2,3} Hz$`)
		batteryRegex := regexp.MustCompile(`^\d{4} mAh$`)
		for _, element := range phoneElements {

			if ramRegex.MatchString(element) {
				if RAM == "" {
					RAM = element
				}
			} else if storageRegex.MatchString(element) {
				if Storage == "" {
					Storage = element
				}
			} else if caleRegex.MatchString(element) {
				if Cale == "" {
					Cale = element
				}
			} else if hertznRegex.MatchString(element) {
				if Herce == "" {
					Herce = element
				}
			} else if batteryRegex.MatchString(element) {
				if Battery == "" {
					Battery = element
				}
			}
		}
		Display = fmt.Sprintf("%s, %s", Cale, Herce)

		fullProductInfo = append(fullProductInfo, Brand)
		fullProductInfo = append(fullProductInfo, Model)
		fullProductInfo = append(fullProductInfo, Price)
		fullProductInfo = append(fullProductInfo, Procesor)
		fullProductInfo = append(fullProductInfo, RAM)
		fullProductInfo = append(fullProductInfo, Storage)
		fullProductInfo = append(fullProductInfo, Battery)
		fullProductInfo = append(fullProductInfo, Display)

	})

	c.Visit(baseURL)

	return fullProductInfo
}
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

	cur, err := m.DatabaseCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	if err := cur.All(context.TODO(), &products); err != nil {
		log.Fatal(err)
	}
	for _, product := range products {
		link := product.ProductURL
		phoneInfo := mediaMarktScrapHelper(link) // CHANGE TO KOMPUTRONIKSCRAP HELPER
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
		_, err := m.DatabaseCollection.UpdateOne(context.Background(), bson.M{"product_url": link}, update)
		if err != nil {
			log.Fatal(err)
		}
	}
}
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

	cur, err := m.DatabaseCollection.Find(context.Background(), filter)
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
		_, err := m.DatabaseCollection.UpdateOne(context.Background(), bson.M{"product_url": link}, update)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func mediaMarktScrapHelper(baseURL string) []string {
	c := colly.NewCollector()
	// baseURL := "https://mediamarkt.pl/telefony-i-smartfony/smartfon-samsung-galaxy-s23-8-256gb-czarny-sm-s916bzkdeue"

	phoneElements := make([]string, 0)
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
		// fmt.Println(element)
		test := strings.Split(element, "\n")
		for _, element := range test {
			phoneElements = append(phoneElements, element)
		}

	})
	c.OnHTML("h1.title.is-heading", func(e *colly.HTMLElement) {
		fullPhone := e.Text
		fullPhoneInfo := strings.Split(fullPhone, " ")

		Brand := fullPhoneInfo[1]
		Model := ""
		for i := 2; i < len(fullPhoneInfo); i++ {
			Model = Model + fullPhoneInfo[i] + " "
		}
		phoneInfo = append(phoneInfo, Brand)
		phoneInfo = append(phoneInfo, Model)

	})
	c.OnHTML("div.main-price.is-big span.whole", func(e *colly.HTMLElement) {
		elemnt := e.Text
		re := regexp.MustCompile(`\b\d+\b`)
		matches := re.FindAllString(elemnt, 1)
		phoneInfo = append(phoneInfo, matches[0])
	})

	c.OnScraped(func(r *colly.Response) {
		Brand := phoneInfo[0]
		Model := phoneInfo[1]
		Price := phoneInfo[2] + " zł"
		Display := ""
		Procesor := ""
		RAM := ""
		Storage := ""
		Battery := ""

		Cale := ""
		Herce := ""
		fmt.Println("Finished", r.Request.URL)

		for i := 0; i < len(phoneElements); i++ {
			if strings.HasPrefix(phoneElements[i], "Rozmiar przekątnej") {
				Cale = strings.TrimLeft(phoneElements[i+1], " ") + `"`
			} else if strings.HasPrefix(phoneElements[i], "Częstotliwość odświeżania") {
				Herce = phoneElements[i+1] + "Hz"
			} else if strings.HasPrefix(phoneElements[i], "Pamięć wewnętrzna") {
				Storage = strings.TrimLeft(phoneElements[i+1], " ")
			} else if strings.HasPrefix(phoneElements[i], "Informuje o ilość") {
				RAM = strings.TrimLeft(phoneElements[i+1], " ")
			} else if strings.HasPrefix(phoneElements[i], "Określa nazwę i model") {
				Procesor = strings.TrimLeft(phoneElements[i+1], " ")
			} else if strings.HasPrefix(phoneElements[i], "Informuje o pojemności akumulatora") {
				Battery = strings.TrimLeft(phoneElements[i+1], " ")
			}
		}
		Display = fmt.Sprintf("%s,%s", Cale, Herce)
		if Storage != "" {
			Storage = Storage + " mAh"
		}

		fullProductInfo = append(fullProductInfo, Brand)
		fullProductInfo = append(fullProductInfo, Model)
		fullProductInfo = append(fullProductInfo, Price)
		fullProductInfo = append(fullProductInfo, Procesor)
		fullProductInfo = append(fullProductInfo, RAM)
		fullProductInfo = append(fullProductInfo, Storage)
		fullProductInfo = append(fullProductInfo, Battery)
		fullProductInfo = append(fullProductInfo, Display)

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

func Test2() {
	c := colly.NewCollector()
	baseURL := "https://www.komputronik.pl/product/826977/xiaomi-redmi-note-12s-8-256gb-czarny-onyx-black-.htmla"

	phoneElements := make([]string, 0)
	// phoneInfo := make([]string, 0)
	// fullProductInfo := make([]string, 0)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.OnHTML("div.p-1 p", func(e *colly.HTMLElement) {
		elements := e.Text
		// fmt.Println(element)
		test := strings.ReplaceAll(elements, " ", "")
		// test = strings.ReplaceAll(elements, "\n", "")
		// test := strings.Split(elements, "\n")
		test = strings.ReplaceAll(test, " ", "")
		xd := strings.Split(test, "\n")
		for _, element := range xd {
			phoneElements = append(phoneElements, element)
		}

	})
	c.OnHTML("div.relative.flex span", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	/////////// DZIALA
	// c.OnHTML("h1.tests-product-name", func(e *colly.HTMLElement) {
	// 	fullPhone := e.Text
	// 	fmt.Println(fullPhone)
	// 	// fullPhoneInfo := strings.Split(fullPhone, " ")

	// 	// Brand := fullPhoneInfo[1]
	// 	// Model := ""
	// 	// for i := 2; i < len(fullPhoneInfo); i++ {
	// 	// 	Model = Model + fullPhoneInfo[i] + " "
	// 	// }
	// 	// phoneInfo = append(phoneInfo, Brand)
	// 	// phoneInfo = append(phoneInfo, Model)

	// })
	/////////// DZIALA
	// c.OnHTML("div.font-bold.leading-8.text-3xl", func(e *colly.HTMLElement) {
	// 	elemnt := e.Text
	// 	fmt.Println(elemnt)
	// 	// re := regexp.MustCompile(`\b\d+\b`)
	// 	// matches := re.FindAllString(elemnt, 1)
	// 	// phoneInfo = append(phoneInfo, matches[0])
	// })

	c.OnScraped(func(r *colly.Response) {
		// for _, element := range phoneElements {
		// 	fmt.Println(element)
		// }
		// Brand := phoneInfo[0]
		// Model := phoneInfo[1]
		// Price := phoneInfo[2]
		// Display := ""
		// Procesor := ""
		// RAM := ""
		// Storage := ""
		// Battery := ""

		// Cale := ""
		// Herce := ""
		// fmt.Println("Finished", r.Request.URL)

		// for i := 0; i < len(phoneElements); i++ {
		// 	if strings.HasPrefix(phoneElements[i], "Rozmiar przekątnej") {
		// 		Cale = phoneElements[i+1] + `"`
		// 	} else if strings.HasPrefix(phoneElements[i], "Częstotliwość odświeżania") {
		// 		Herce = phoneElements[i+1] + "Hz"
		// 	} else if strings.HasPrefix(phoneElements[i], "Pamięć wewnętrzna") {
		// 		Storage = phoneElements[i+1]
		// 	} else if strings.HasPrefix(phoneElements[i], "Informuje o ilość") {
		// 		RAM = phoneElements[i+1]
		// 	} else if strings.HasPrefix(phoneElements[i], "Określa nazwę i model") {
		// 		Procesor = phoneElements[i+1]
		// 	} else if strings.HasPrefix(phoneElements[i], "Informuje o pojemności akumulatora") {
		// 		Battery = phoneElements[i+1]
		// 	}
		// }
		// Display = fmt.Sprintf("%s,%s", Cale, Herce)

		// fullProductInfo = append(fullProductInfo, Brand)
		// fullProductInfo = append(fullProductInfo, Model)
		// fullProductInfo = append(fullProductInfo, Price)
		// fullProductInfo = append(fullProductInfo, Procesor)
		// fullProductInfo = append(fullProductInfo, RAM)
		// fullProductInfo = append(fullProductInfo, Storage)
		// fullProductInfo = append(fullProductInfo, Battery)
		// fullProductInfo = append(fullProductInfo, Display)

	})
	c.Visit(baseURL)

}
