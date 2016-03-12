package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"sync"
)

type ResultItem struct {
	Title      string
	ReleasedAt string
}

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

func GetDVDPages(url string) []string {
	urls := []string{url}
	doc, _ := goquery.NewDocument(url)
	doc.Find("ul.pageList").First().Find("li").Each(func(_ int, li *goquery.Selection) {
		a := li.Find("a").First()
		aClass, _ := a.Attr("class")
		liClass, _ := li.Attr("class")

		if liClass != "last" && aClass != "active" {
			href, _ := a.Attr("href")
			fullUrl := GenerateUrlWithPath(url, href)
			urls = append(urls, fullUrl)
		}
	})
	return urls
}

func GenerateUrlWithPath(pageUrl string, path string) string {
	u, _ := url.Parse(pageUrl)
	fullUrl := u.Scheme + "://" + u.Host + path
	return fullUrl
}

func main() {
	url := "http://store-tsutaya.tsite.jp/top/rels/dvd_rental.html"
	pages := GetDVDPages(url)

	itemUrls := []string{}
	for _, page := range pages {
		itemUrls = append(itemUrls, GetDVDItemUrls(page)...)
	}
	results := GoGetDVDItems(itemUrls)

	for _, result := range results {
		fmt.Println(result.Title + " : " + result.ReleasedAt)
	}
}
