package main

import (
    // import standard libraries
    "fmt"
    "log"

    // import third party libraries
    "github.com/PuerkitoBio/goquery"
)

func postScrape() {
	//doc, err := goquery.NewDocument("http://jonathanmh.com")
	//doc, err := goquery.NewDocument("https://www.bi.go.id/id/tentang-bi/pesan-gubernur/Contents/Default.aspx")
	doc, err := goquery.NewDocument("https://www.bca.co.id/id/Individu/Sarana/Kurs-dan-Suku-Bunga/Kurs-dan-Kalkulator")
    if err != nil {
        log.Fatal(err)
    }

    // use CSS selector found with the browser inspector
	// for each, use index and item
	/*
    doc.Find("#main article .entry-title").Each(func(index int, item *goquery.Selection) {
        title := item.Text()
        linkTag := item.Find("a")
        link, _ := linkTag.Attr("href")
        fmt.Printf("Post #%d: %s - %s\n", index, title, link)
	})
	
	doc.Find(".menu-item-type-custom").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		link, _ := item.Find("a").Attr("href")
		fmt.Printf("Post #%d: %s - %s\n", index, title, link)
	})
	
	doc.Find(".ms-rteThemeFontFace-1").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		fmt.Printf("Post #%d: %s\n", index, title)
	})
	*/

	doc.Find(".ms-rteThemeFontFace-1").Each(func(index int, item *goquery.Selection) {
		link, _ := item.Find("td").Attr("style")
		fmt.Printf("Post #%d: %s\n", index, link)
	})
}

func main() {
    postScrape()
}