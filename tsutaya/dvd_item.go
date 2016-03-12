package tsutaya

import (
	"github.com/PuerkitoBio/goquery"
	"sync"
)

func GetDVDItem(itemUrl string) ResultItem {
	doc, _ := goquery.NewDocument(itemUrl)
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
func GoGetDVDItemUrls(pages []string) []string {
	itemUrls := []string{}
	var wg sync.WaitGroup
	for _, page := range pages {
		wg.Add(1)
		go func(page string) {
			defer wg.Done()
			itemUrls = append(itemUrls, GetDVDItemUrls(page)...)
		}(page)
	}
	wg.Wait()
	return itemUrls
}

// Get dvd item url list
func GetDVDItemUrls(url string) []string {
	urls := []string{}
	doc, _ := goquery.NewDocument(url)
	doc.Find(".itemGroup .imageBlock a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		urls = append(urls, url)
	})
	return urls
}
