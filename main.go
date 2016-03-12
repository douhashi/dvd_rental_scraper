package main

import (
	"./dvd_rental_scraper"
	"fmt"
)

func main() {
	url := "http://store-tsutaya.tsite.jp/top/rels/dvd_rental.html"
	pages := dvd_rental_scraper.GetDVDPages(url)

	itemUrls := dvd_rental_scraper.GoGetDVDItemUrls(pages)
	results := dvd_rental_scraper.GoGetDVDItems(itemUrls)

	for _, result := range results {
		fmt.Println(result.Title + " : " + result.ReleasedAt)
	}
}
