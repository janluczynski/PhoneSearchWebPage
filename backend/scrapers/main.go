package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"main.go/commons"
	scrapersInfo "main.go/scrapers/scrapingInfo"
)

func main() {
	scrapersInfo.FakeXKomRequest()
	// scrapersLinks.XkomScrap()
	scrapersInfo.XkomScrapProductInfo()
	// Test()
}

func Test() {
	c := colly.NewCollector()
	baseURL := "https://www.x-kom.pl/p/671031-smartfon-telefon-ulefone-armor-8-pro-8-128gb-czerwony.html"

	var Specification []string
	var Product commons.Product

	c.OnHTML(".sc-13p5mv-2.fxqQxb .sc-1s1zksu-0.sc-1s1zksu-1.hHQkLn.sc-13p5mv-0.VGBov", func(e *colly.HTMLElement) {
		Specification = append(Specification, e.Text)
	})
	c.OnHTML(".sc-1bker4h-10.kHPtVn h1", func(e *colly.HTMLElement) {
		Product.Brand = strings.Split(e.Text, " ")[0]
		Product.Model = strings.Join(strings.Split(e.Text, " ")[1:], " ")
	})
	c.OnHTML(".sc-n4n86h-1.hYfBFq", func(e *colly.HTMLElement) {
		Price, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(e.Text, ",00 zł", ""), " ", ""), 32)
		if err != nil {
			fmt.Println(err)
		}
		Product.Price = float32(Price)
	})
	c.OnHTML(".sc-1tblmgq-0.sc-1tblmgq-3.cIswgX.sc-jiiyfe-2.jGSlBb img", func(e *colly.HTMLElement) {
		Product.ImageURL = e.Attr("src")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		for _, element := range Specification {
			if strings.Contains(element, "Procesor") {
				Product.Processor = strings.ReplaceAll(element, "Procesor", "")
			} else if strings.Contains(element, "Pamięć RAM") {
				ram := strings.ReplaceAll(element, "Pamięć RAM", "")
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
			} else if strings.Contains(element, "Pamięć wbudowana") {
				Storage := strings.ReplaceAll(element, "Pamięć wbudowana", "")
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
			} else if strings.Contains(element, "Pojemność baterii") {
				BatteryInt, err := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(element, "Pojemność baterii", ""), " mAh", ""))
				if err != nil {
					fmt.Println(err)
				}
				Product.Battery = BatteryInt
			} else if strings.Contains(element, "Przekątna ekranu") {
				Product.Display = strings.ReplaceAll(strings.ReplaceAll(element, "Przekątna ekranu", ""), ",", ".")
			}
		}

		fmt.Println(Product)
	})
	c.Visit(baseURL)
}

// func Test() {
// 	c := colly.NewCollector()
// 	baseURL := "https://www.x-kom.pl/p/640021-smartfon-telefon-myphone-2220-czarny.html"

// 	phoneSpecification := make([]string, 0)
// 	phoneInfo := make([]string, 0)

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})

// 	c.OnError(func(_ *colly.Response, err error) {
// 		fmt.Println("Something went wrong:", err)
// 	})
// 	c.OnHTML(".sc-13p5mv-3", func(e *colly.HTMLElement) { // scraping specyfication
// 		element := e.Text
// 		dividedPhoneSpecification := strings.Split(element, "\n")
// 		for _, specification := range dividedPhoneSpecification {
// 			phoneSpecification = append(phoneSpecification, specification)
// 		}
// 	})
// 	c.OnHTML(".sc-1bker4h-10.kHPtVn", func(e *colly.HTMLElement) { //scraping brand + model
// 		element := e.Text
// 		fullPhoneInfo := strings.Split(element, " ")

// 		Brand := fullPhoneInfo[0]
// 		Model := strings.Join(fullPhoneInfo[1:], " ")
// 		phoneInfo = append(phoneInfo, Brand, Model)

// 	})
// 	c.OnHTML(".sc-n4n86h-1.hYfBFq", func(e *colly.HTMLElement) { //scraping price
// 		phoneInfo = append(phoneInfo, e.Text)
// 	})

// 	c.OnScraped(func(r *colly.Response) {
// 		bracketsReegx := regexp.MustCompile(`\([^)]*\)`)
// 		Brand := phoneInfo[0]
// 		Model := bracketsReegx.ReplaceAllString(phoneInfo[1], "")
// 		var Price float32
// 		Display := ""
// 		Processor := phoneSpecification[0]
// 		var RAM int
// 		var Storage int
// 		var Battery int
// 		Inches := ""
// 		Hertz := ""

// 		fmt.Println("Finished", r.Request.URL)
// 		ramRegex := regexp.MustCompile(`^\b([1-9]|1\d|2[0-9])\b\s*(MB|GB)$`)
// 		storageRegex := regexp.MustCompile(`^(32|[4-9][0-9]|[1-4][0-9][0-9]|5[0-1][0-2]) MB|GB$|1 TB$`)
// 		inchesRegex := regexp.MustCompile(`^\d,\d{1,2}"$`)
// 		hertznRegex := regexp.MustCompile(`^\d{2,3} Hz$`)
// 		batteryRegex := regexp.MustCompile(`mAh$`)
// 		for _, element := range phoneSpecification {

// 			if ramRegex.MatchString(element) {
// 				if RAM == 0 {
// 					if strings.Contains(element, "GB") {
// 						element = strings.TrimSpace(strings.ReplaceAll(element, "GB", ""))
// 						parsedRAM, err := strconv.Atoi(element)
// 						if err != nil {
// 							fmt.Println(err)
// 						}
// 						RAM = parsedRAM * 1024
// 					} else if strings.Contains(element, "MB") {
// 						element = strings.TrimSpace(strings.ReplaceAll(element, "MB", ""))
// 						parsedRAM, err := strconv.Atoi(element)
// 						if err != nil {
// 							fmt.Println(err)
// 						}
// 						RAM = parsedRAM
// 					}
// 				}
// 			} else if storageRegex.MatchString(element) {
// 				if Storage == 0 {
// 					if strings.Contains(element, "GB") {
// 						element = strings.TrimSpace(strings.ReplaceAll(element, "GB", ""))
// 						parsedStorage, err := strconv.Atoi(element)
// 						if err != nil {
// 							fmt.Println(err)
// 						}
// 						Storage = parsedStorage * 1024
// 					} else if strings.Contains(element, "MB") {
// 						element = strings.TrimSpace(strings.ReplaceAll(element, "MB", ""))
// 						parsedStorage, err := strconv.Atoi(element)
// 						if err != nil {
// 							fmt.Println(err)
// 						}
// 						Storage = parsedStorage
// 					}
// 				}
// 			} else if inchesRegex.MatchString(element) {
// 				if Inches == "" {
// 					Inches = strings.ReplaceAll(element, ",", ".")
// 				}
// 			} else if hertznRegex.MatchString(element) {
// 				if Hertz == "" {
// 					Hertz = element
// 				}
// 			} else if batteryRegex.MatchString(element) {
// 				if Battery == 0 {
// 					element = strings.TrimSpace(strings.ReplaceAll(element, "mAh", ""))
// 					parsedBattery, err := strconv.Atoi(element)
// 					if err != nil {
// 						fmt.Println(err)
// 					}
// 					Battery = parsedBattery
// 				}
// 			}
// 		}
// 		if Inches != "" && Hertz != "" {
// 			Display = fmt.Sprintf("%s, %s", Inches, Hertz)
// 		} else if Inches != "" && Hertz == "" {
// 			Display = Inches
// 		} else {
// 			Display = "N/A"
// 		}
// 		// Price, err := strconv.ParseFloat(strings.ReplaceAll(phoneInfo[2], " zł", ""), 32)
// 		// if err != nil {
// 		// 	fmt.Println(err)
// 		// }

// 		//fullProductInfo = append(fullProductInfo, Brand, Model, Price, Procesor, RAM, Storage, Battery, Display)
// 		fullProduct := commons.Product{
// 			Brand:     Brand,
// 			Model:     Model,
// 			Price:     Price,
// 			Processor: Processor,
// 			RAM:       RAM,
// 			Storage:   Storage,
// 			Battery:   Battery,
// 			Display:   Display,
// 		}
// 		fmt.Println(fullProduct)
// 	})

// 	c.Visit(baseURL)
// }
