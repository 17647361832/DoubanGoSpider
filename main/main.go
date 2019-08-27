package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func SpiderDouban(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {
	var array [10]int
	var start = 0
	var add = 25
	for i := 0; i < len(array); i++ {
		array[i] = start
		start += add
	}
	const url string = "https://movie.douban.com/top250?start="
	for _, num := range array {
		var newurl string
		newurl = url + strconv.Itoa(num)
		var result = SpiderDouban(newurl)
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(result))
		dom.Find(".grid_view > li > div > .info").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Find(".hd > a > .title").Text())
		})
	}

}
