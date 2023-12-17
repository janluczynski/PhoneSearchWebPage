package scrapers

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	common "main.go/commons"
	mongodb "main.go/mongoDB"
)

func KomputronikScrap() {
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{
		Delay:       1 * time.Second,
		Parallelism: 2,
	})

	baseURL := "https://www.komputronik.pl/category/1596/telefony.html?showBuyActiveOnly=0&p="
	var productLinks []string
	visitedLinks := make(map[string]bool)

	productLinkRegex := regexp.MustCompile(`https:\/\/www\.komputronik\.pl\/product`)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if productLinkRegex.MatchString(link) && !visitedLinks[link] {
			productLinks = append(productLinks, link)
			visitedLinks[link] = true
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	for i := 1; i <= 42; i++ {
		scrapeURL := baseURL + fmt.Sprintf("%d", i)
		c.Visit(scrapeURL)
		time.Sleep(3 * time.Second)
	}

	var linksOnly []string
	for link := range visitedLinks {
		linksOnly = append(linksOnly, link)
	}

	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Some error occured. Err: %s \n", err)
	}
	m, err := mongodb.InitDB()
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, link := range linksOnly {
		product := common.Product{
			ProductURL: link,
			ProductID:  uuid.New().String(),
		}
		m.AddProduct(product)
	}

}
