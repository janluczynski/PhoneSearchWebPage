package main

import (
	scrapers "main.go/scrapers/scrapingInfo"
)

func main() {
	// scrapers.FakeXKomRequest()
	// scrapersLinks.KomputronikScrap()
	// scrapers.KomputronikScrapProductInfo()
	scrapers.FakeMediaMarktRequest()
	// scrapersLinks.MediaMarktScrap()
	scrapers.MediaMarktScrapProductInfo()
}
