package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	println("start")
	http.HandleFunc("/login", screen)
	http.ListenAndServe(":8080", nil)
}

func screen(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("assets/index.html")
	if err != nil {
		panic(err.Error())
	}

	if r.Method == "POST" {
		r.ParseForm()
		doc, err := goquery.NewDocument(r.Form["url"][0])
		if err != nil {
			log.Fatal(err)
		}
		title := doc.Find("title").Text()
		urls := []string{}

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			url := s.AttrOr("href", "")
			urls = append(urls, url)
		})

		data := struct {
			Title string
			URLs  []string
		}{
			Title: title,
			URLs:  urls,
		}

		// urls.Each(func(i int, s *goquery.Selection) {
		// 	fmt.Println("URL: ", s.AttrOr("href", ""))
		// })

		t.Execute(w, data)
	} else if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}
