package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	http.HandleFunc("/", Scraping)
	http.ListenAndServe(":8080", nil)
}

func Scraping(w http.ResponseWriter, r *http.Request) {
	url := "#"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("title").Text()
	fmt.Println("タイトル:", title)

	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		fmt.Println("見出し：", s.Text())
	})

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		fmt.Println("リンク：", link)
	})
}
