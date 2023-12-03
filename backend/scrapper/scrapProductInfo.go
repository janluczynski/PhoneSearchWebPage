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
		bracketsReegx := regexp.MustCompile(`\([^)]*\)`)
		Brand := phoneInfo[0]
		Model := bracketsReegx.ReplaceAllString(phoneInfo[1], "")
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
					Cale = strings.ReplaceAll(element, ",", ".")
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
		if Cale != "" && Herce != "" {
			Display = fmt.Sprintf("%s, %s", Cale, Herce)
		} else if Cale != "" && Herce == "" {
			Display = Cale
		} else {
			Display = "N/A"
		}
		if RAM == "" {
			RAM = "N/A"
		}
		if Battery == "" {
			Battery = "N/A"
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

	phoneElements := make([]string, 0)
	phoneInfo := make([]string, 0)
	fullProductInfo := make([]string, 0)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.OnHTML("div.p-1 p", func(e *colly.HTMLElement) {
		elements := e.Text
		// fmt.Println(elements)
		test := strings.ReplaceAll(elements, " ", "")
		xd := strings.Split(test, "\n")
		for _, element := range xd {
			phoneElements = append(phoneElements, element)
		}

	})

	/////////// DZIALA
	c.OnHTML("h1.tests-product-name", func(e *colly.HTMLElement) {
		fullPhone := e.Text
		// fmt.Println(fullPhone)
		fullPhoneInfo := strings.Fields(fullPhone)

		Brand := fullPhoneInfo[0]
		Model := ""
		for i := 1; i < len(fullPhoneInfo); i++ {
			Model = Model + fullPhoneInfo[i] + " "
		}
		phoneInfo = append(phoneInfo, Brand)
		phoneInfo = append(phoneInfo, Model)

	})
	/////////// DZIALA
	c.OnHTML("div.font-bold.leading-8.text-3xl", func(e *colly.HTMLElement) {
		elemnt := e.Text
		refactoring := strings.ReplaceAll(elemnt, " ", "")
		phonePrice := strings.Split(refactoring, "\n")
		phoneInfo = append(phoneInfo, phonePrice[0])
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

		Cale := ""

		fmt.Println("Finished", r.Request.URL)
		storageRegex := regexp.MustCompile(`GB$`)
		caleRegex := regexp.MustCompile(`^\d.\d{1,2}cale$`)
		batteryRegex := regexp.MustCompile(`^\d{4}mAh$`)
		for i := 0; i < len(phoneElements); i++ {
			if storageRegex.MatchString(phoneElements[i]) {
				if Storage == "" {
					Storage = phoneElements[i]
				} else if RAM == "" {
					RAM = phoneElements[i]
					if storageRegex.MatchString(phoneElements[i+2]) {
						Procesor = phoneElements[i+4]
					} else {
						Procesor = phoneElements[i+2]
					}
				}
			} else if caleRegex.MatchString(phoneElements[i]) {
				if Cale == "" {
					Cale = strings.ReplaceAll(phoneElements[i], "cale", `"`)
				}
			} else if batteryRegex.MatchString(phoneElements[i]) {
				if Battery == "" {
					Battery = phoneElements[i]
				}
			}
		}
		if Cale == "" {
			Cale = "N/A"
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

		fullProductInfo = append(fullProductInfo, Brand)
		fullProductInfo = append(fullProductInfo, Model)
		fullProductInfo = append(fullProductInfo, Price)
		fullProductInfo = append(fullProductInfo, Procesor)
		fullProductInfo = append(fullProductInfo, RAM)
		fullProductInfo = append(fullProductInfo, Storage)
		fullProductInfo = append(fullProductInfo, Battery)
		fullProductInfo = append(fullProductInfo, Cale)
	})
	c.Visit(baseURL)

	return fullProductInfo
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
		if Cale != "" && Herce != "" {
			Display = fmt.Sprintf("%s, %s", Cale, Herce)
		} else if Cale != "" && Herce == "" {
			Display = Cale
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
