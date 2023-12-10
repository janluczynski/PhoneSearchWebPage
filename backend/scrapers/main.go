package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"main.go/commons"
	scrapersInfo "main.go/scrapers/scrapingInfo"
)

func main() {
	// scrapersLinks.KomputronikScrap()
	// scraperInfo.KomputronikScrapProductInfo()
	scrapersInfo.FakeXKomRequest()
	Test()
}

func Test() {
	c := colly.NewCollector()
	baseURL := "https://www.x-kom.pl/p/1127067-smartfon-telefon-xiaomi-redmi-note-12-4-128gb-ice-blue.html"

	var Specification []string
	var Product commons.Product

	c.OnHTML(".sc-13p5mv-2.fxqQxb .sc-1s1zksu-0.sc-1s1zksu-1.hHQkLn.sc-13p5mv-0.VGBov", func(e *colly.HTMLElement) {
		Specification = append(Specification, e.Text)
	})
	c.OnHTML(".sc-1bker4h-10.kHPtVn h1", func(e *colly.HTMLElement) {
		PhoneName := strings.Split(e.Text, " ")
		GBRegex := regexp.MustCompile(`GB$`)
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
		fmt.Println(Product)

	})

	c.Visit(baseURL)
}
