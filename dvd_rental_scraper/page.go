package dvd_rental_scraper

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

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
