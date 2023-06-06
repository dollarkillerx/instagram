package tiktok

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestTiktokGet(t *testing.T) {
	get, strings, err := TiktokGet("https://www.tiktok.com/@yy._.101/video/7237045637030743301", true)
	if err != nil {
		panic(err)
	}

	fmt.Println(get)
	fmt.Println(strings)
}

func TestQueryParse(t *testing.T) {
	file, err := os.ReadFile("idx.html")
	if err != nil {
		log.Fatalln(err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	img := doc.Find(".result_author").AttrOr("src", "")
	fmt.Println(img)

	doc.Find(".flex-1").Find("a").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(i, " ", selection.AttrOr("href", ""))
	})
}
