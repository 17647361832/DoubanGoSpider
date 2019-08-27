package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

var visited = map[string]bool{}

func main() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6)"+
			" AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36"),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*.douban.*", Parallelism: 255})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("访问中", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("发生错误：", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("访问结束", r.Request.URL)
	})

	c.OnHTML(".hd", func(e *colly.HTMLElement) {
		// https://movie.douban.com/subject/1292052/
		// 用'/'分割链接，取index是4的元素
		log.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
			strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("结束", r.Request.URL)
	})

	c.Visit("https://movie.douban.com/top250?start=0&filter=")
	c.Wait()
}
