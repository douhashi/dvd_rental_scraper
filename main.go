package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type ItemResult struct {
	Title      string
	ReleasedAt string
}

func GetDVDItem(url string) ItemResult {
	doc, _ := goquery.NewDocument(url)
	title := doc.Find(".header h2 span").First().Text()
	releasedAt := doc.Find(".detailBox li").First().Next().Text()

	result := ItemResult{title, releasedAt}
	return result
}

func GetDVDList(url string) []string {
	list := []string{}
	doc, _ := goquery.NewDocument(url)
	doc.Find(".itemGroup .imageBlock a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		list = append(list, url)
	})
	return list
}

func main() {
	url := "http://store-tsutaya.tsite.jp/top/rels/dvd_rental.html"
	list := GetDVDList(url)
	for _, v := range list {
		result := GetDVDItem(v)
		fmt.Println(result.Title + " : " + result.ReleasedAt)
	}
}
