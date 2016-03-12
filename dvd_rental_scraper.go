package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"sync"
)

type ResultItem struct {
	Title      string
	ReleasedAt string
}

func GetDVDItem(url string) ResultItem {
	doc, _ := goquery.NewDocument(url)
	title := doc.Find(".header h2 span").First().Text()
	releasedAt := doc.Find(".detailBox li").First().Next().Text()

	result := ResultItem{title, releasedAt}
	return result
}
func GoGetDVDItems(urls []string) []ResultItem {
	results := []ResultItem{}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			results = append(results, GetDVDItem(url))
		}(url)
	}
	wg.Wait()
	return results
}

// Get dvd item url list
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
	results := GoGetDVDItems(list)

	for _, result := range results {
		fmt.Println(result.Title + " : " + result.ReleasedAt)
	}
}
