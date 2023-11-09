package scrappers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func XkomScrap() {
	c := colly.NewCollector()

	baseURL := "https://www.x-kom.pl/g-4/c/1590-smartfony-i-telefony.html?page="
	var productLinks []string
	visitedLinks := make(map[string]bool)

	// Regular expression to match product links
	productLinkRegex := regexp.MustCompile(`/p/\d+-[a-z0-9-]+\.html`)

	for i := 1; i <= 37; i++ {
		scrapeURL := baseURL + fmt.Sprintf("%d", i)

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if productLinkRegex.MatchString(link) && !visitedLinks[link] && !strings.HasSuffix(link, "#Opinie") && !strings.HasPrefix(link, "https://www.x-kom.pl") {
				productLinks = append(productLinks, link)
				visitedLinks[link] = true
				// fmt.Println(modifiedLink)
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

	// Save linksOnly to a new JSON file
	err := saveToJSON(linksOnly, "xkomLinks.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func KomputronikScrap() {
	c := colly.NewCollector()

	baseURL := "https://www.komputronik.pl/category/1596/telefony.html?showBuyActiveOnly=0&p="
	var productLinks []string
	visitedLinks := make(map[string]bool)

	productLinkRegex := regexp.MustCompile(`https:\/\/www\.komputronik\.pl\/product`)

	for i := 1; i <= 42; i++ {
		scrapeURL := baseURL + fmt.Sprintf("%d", i)

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if productLinkRegex.MatchString(link) && !visitedLinks[link] {
				productLinks = append(productLinks, link)
				visitedLinks[link] = true
				// fmt.Println(modifiedLink)
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
		// correctLink := "https://www.x-kom.pl" + link
		linksOnly = append(linksOnly, link)
	}

	// Save linksOnly to a new JSON file
	err := saveToJSON(linksOnly, "komputronikLinks.json")
	if err != nil {
		fmt.Println("Error:", err)
	}

}

func MediaMarktScrap() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	baseURL := "https://mediamarkt.pl/telefony-i-smartfony/smartfony/wszystkie-smartfony?page=" //".bhtml"
	var productLinks []string
	visitedLinks := make(map[string]bool)

	productLinkRegex := regexp.MustCompile(`/telefony-i-smartfony/smartfon-`)

	rand.Seed(time.Now().UnixNano()) // Seed for random number generation

	for i := 1; i <= 43; i++ {
		scrapeURL := baseURL + fmt.Sprintf("%d", i) //+ ".bhtml"

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if productLinkRegex.MatchString(link) && !visitedLinks[link] && !strings.HasSuffix(link, "#reviews") && !strings.HasPrefix(link, "https://mediamarkt.pl") {
				productLinks = append(productLinks, link)
				visitedLinks[link] = true
				fmt.Println(link)
			}
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		c.OnError(func(_ *colly.Response, err error) {
			fmt.Println("Something went wrong:", err)
		})

		err := c.Visit(scrapeURL)
		if err != nil {
			fmt.Println("Error visiting", scrapeURL, ":", err)
		}
	}

	var linksOnly []string
	for link := range visitedLinks {
		correctLink := "https://mediamarkt.pl" + link
		linksOnly = append(linksOnly, correctLink)
	}

	// Save linksOnly to a new JSON file
	err := saveToJSON(linksOnly, "mediamarktLinks.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
func saveToJSON(data interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	fmt.Println("Data saved to", filename)
	return nil
}
