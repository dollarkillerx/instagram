package tiktok

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"strings"
)

var v1 = "https://ssstik.io/abc?url=dl"
var v2 = "https://tiktokdownload.online/abc?url=dl"

func TiktokGet(url string, alternative bool) (string, []string, error) {
	var img string
	var videos []string
	post, err := resty.New().R().
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36").
		SetFormData(map[string]string{
			"id":     url,
			"locale": "en",
			"tt":     "OGkyekxl",
		}).Post(func() string {
		if !alternative {
			return v1
		}
		return v2
	}())
	if err != nil {
		return "", nil, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(post.Body()))
	if err != nil {
		return "", nil, err
	}

	// Find the review items
	img = doc.Find(".result_author").AttrOr("src", "")

	doc.Find(".flex-1").Find("a").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			videos = append(videos, strings.TrimSpace(selection.AttrOr("href", "")))
		} else if i == 1 {
			videos = append(videos, strings.TrimSpace(selection.AttrOr("href", "")))
		}
	})

	return img, videos, nil
}
