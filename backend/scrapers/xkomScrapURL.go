package scrapers

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	common "main.go/commons"
	mongodb "main.go/mongoDB"
)

func XkomScrap() {
	c := colly.NewCollector()

	baseURL := "https://www.x-kom.pl/g-4/c/1590-smartfony-i-telefony.html?page="
	var productLinks []string
	visitedLinks := make(map[string]bool)

	productLinkRegex := regexp.MustCompile(`/p/\d+-[a-z0-9-]+\.html`)

	for i := 1; i <= 26; i++ {
		scrapeURL := baseURL + fmt.Sprintf("%d", i) + "&hide_unavailable=1"

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if productLinkRegex.MatchString(link) && !visitedLinks[link] && !strings.HasSuffix(link, "#Opinie") && !strings.HasPrefix(link, "https://www.x-kom.pl") {
				productLinks = append(productLinks, link)
				visitedLinks[link] = true
			}
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		c.OnError(func(_ *colly.Response, err error) {
			fmt.Println("Something went wrong:", err)
		})

		c.Visit(scrapeURL)
	}

	var linksOnly []string
	for link := range visitedLinks {
		correctLink := "https://www.x-kom.pl" + link
		linksOnly = append(linksOnly, correctLink)
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
