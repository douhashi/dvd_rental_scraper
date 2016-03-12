package main

import (
	"./tsutaya"
	"fmt"
)

func main() {
	url := "http://store-tsutaya.tsite.jp/top/rels/dvd_rental.html"
	pages := tsutaya.GetDVDPages(url)

	itemUrls := tsutaya.GoGetDVDItemUrls(pages)
	results := tsutaya.GoGetDVDItems(itemUrls)

	for _, result := range results {
		fmt.Println(result.Title + " : " + result.ReleasedAt)
	}
}
