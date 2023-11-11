package main

import scrappers "main.go/scrapper"

func main() {
	scrappers.FakeMediaMarktRequest()
	// scrappers.MediaMarktScrap()
	scrappers.MediaMarktScrapProductInfo()
}
