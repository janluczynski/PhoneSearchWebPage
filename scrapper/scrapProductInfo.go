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
		phoneInfo := Test2(link)
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
func KomputronikScrapProductInfo() {

}
func MediaMarktScrapProductInfo() {

}

func Test() {
	c := colly.NewCollector()
	baseURL := "https://www.x-kom.pl/p/1155251-smartfon-telefon-tecno-spark-10-nfc-4-128gb-meta-blue.html"

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
		fmt.Println(fullPhoneInfo)
		Brand := fullPhoneInfo[0]
		Model := ""
		slashRegex := regexp.MustCompile("/")

		for i := 1; i < len(fullPhoneInfo); i++ {
			if !slashRegex.MatchString(fullPhoneInfo[i]) {
				Model = Model + fullPhoneInfo[i] + " "
			} else {
				break
			}
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
		for _, element := range phoneElements {
			ramRegex := regexp.MustCompile(`^\b([1-9]|1\d|2[0-9])\b GB$`)
			storageRegex := regexp.MustCompile(`^(32|[4-9][0-9]|[1-4][0-9][0-9]|5[0-1][0-2]) GB$|1 TB$`)
			caleRegex := regexp.MustCompile(`^\d,\d{1,2}"$`)
			hertznRegex := regexp.MustCompile(`^\d{2,3} Hz$`)
			batteryRegex := regexp.MustCompile(`^\d{4} mAh$`)

			if ramRegex.MatchString(element) {
				RAM = element
			}
			if storageRegex.MatchString(element) {
				Storage = element
			}
			if caleRegex.MatchString(element) {
				Cale = element
			}
			if hertznRegex.MatchString(element) {
				Herce = element
			}
			if batteryRegex.MatchString(element) {
				Battery = element
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

		fmt.Printf("Brand: %s\n", Brand)
		fmt.Printf("Model: %s\n", Model)
		fmt.Printf("Price: %s\n", Price)
		fmt.Printf("Procesor: %s\n", Procesor)
		fmt.Printf("RAM: %s\n", RAM)
		fmt.Printf("Storage: %s\n", Storage)
		fmt.Printf("Battery: %s\n", Battery)
		fmt.Printf("Display: %s\n", Display)
	})

	c.Visit(baseURL)

}

func Test2(baseURL string) []string {
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
		slashRegex := regexp.MustCompile("/")

		for i := 1; i < len(fullPhoneInfo); i++ {
			if !slashRegex.MatchString(fullPhoneInfo[i]) {
				Model = Model + fullPhoneInfo[i] + " "
			} else {
				break
			}
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
		for _, element := range phoneElements {
			ramRegex := regexp.MustCompile(`^\b([1-9]|1\d|2[0-9])\b GB$`)
			storageRegex := regexp.MustCompile(`^(32|[4-9][0-9]|[1-4][0-9][0-9]|5[0-1][0-2]) GB$|1 TB$`)
			caleRegex := regexp.MustCompile(`^\d,\d{1,2}"$`)
			hertznRegex := regexp.MustCompile(`^\d{2,3} Hz$`)
			batteryRegex := regexp.MustCompile(`^\d{4} mAh$`)

			if ramRegex.MatchString(element) {
				RAM = element
			}
			if storageRegex.MatchString(element) {
				Storage = element
			}
			if caleRegex.MatchString(element) {
				Cale = element
			}
			if hertznRegex.MatchString(element) {
				Herce = element
			}
			if batteryRegex.MatchString(element) {
				Battery = element
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

		// fmt.Printf("Brand: %s\n", Brand)
		// fmt.Printf("Model: %s\n", Model)
		// fmt.Printf("Price: %s\n", Price)
		// fmt.Printf("Procesor: %s\n", Procesor)
		// fmt.Printf("RAM: %s\n", RAM)
		// fmt.Printf("Storage: %s\n", Storage)
		// fmt.Printf("Battery: %s\n", Battery)
		// fmt.Printf("Display: %s\n", Display)
	})

	c.Visit(baseURL)

	return fullProductInfo
}
