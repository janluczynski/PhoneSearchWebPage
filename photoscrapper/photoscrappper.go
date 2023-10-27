package scrapper

import (
	"strings"

	"github.com/go-rod/rod"
)

// func main() {

// 	//download https://download-chromium.appspot.com/ , unzip it and put it in the ~/photoscrapper/bin folder

// 	u := launcher.New().Bin("./bin/chrome").MustLaunch()
// 	browser := rod.New().ControlURL(u).MustConnect()

// 	defer browser.MustClose()

// 	var linksToPhotos [][]string

// 	links := readCSV("./../DataSet/dataset.csv")
// 	for _, link := range links[1:] {
// 		if arr := Scrap(browser, link); len(arr) > 0 {
// 			linksToPhotos = append(linksToPhotos, arr)
// 		}
// 	}

// 	fmt.Println(linksToPhotos)
// }

func Scrap(browser *rod.Browser, link string) []string {
	page := browser.MustPage(link)
	photosContainer := page.MustElementX("/html/body/div/div/div[3]/div[1]/div[1]/div/div[2]/div[1]")
	photosDivs := photosContainer.MustElements("div")

	var productPhotos []string

	for _, photoDiv := range photosDivs {
		linksToPhotos := photoDiv.MustElements("source")
		if !linksToPhotos.Empty() {
			linkToPhoto := linksToPhotos[0].MustAttribute("srcset")
			link = strings.Split(*linkToPhoto, ",")[0]
			productPhotos = append(productPhotos, link)
		}

	}
	return productPhotos
}

// func readCSV(path string) []string {
// 	//open file
// 	f, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//close file
// 	defer f.Close()

// 	//read file
// 	r := csv.NewReader(f)

// 	//read all lines
// 	lines, err := r.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//create array of strings
// 	var result []string

// 	//iterate over lines
// 	for _, line := range lines {
// 		//append first column to array
// 		result = append(result, line[0])
// 	}

// 	return result
// }
