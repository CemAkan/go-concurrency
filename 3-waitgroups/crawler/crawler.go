package crawler

import (
	"exercise3/models"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

func Fetch(url string) (PageResult, error) {
	// HTTP isteği
	response, err := http.Get(url)
	if err != nil {
		return PageResult{}, err
	}

	// in normal conditions defer func allows only func calls but there i want to catch the errors which comes from http body's built-in closer func
	defer func(Body io.ReadCloser) {

		if err := Body.Close(); err != nil {
			log.Println("ERROR: Response body closing error. URL:", url)
		}
	}(response.Body)

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		return PageResult{}, err
	}

	var elements []models.Element

	doc.Find("h1,h2,h3,p,a,img").Each(func(i int, s *goquery.Selection) {
		tag := goquery.NodeName(s)
		text := strings.TrimSpace(s.Text())
		attr := ""

		if tag == "a" {
			attr, _ = s.Attr("href")
		} else if tag == "img" {
			attr, _ = s.Attr("src")
		}

		elements = append(elements, models.Element{
			ElementType: tag,
			Content:     text,
			Attribute:   attr,
			Position:    i,
		})
	})

	return PageResult{
		URL:      url,
		Elements: elements,
	}, nil
}

func Crawl(urls []string) []PageResult {
	var wg sync.WaitGroup
	results := make([]PageResult, len(urls))

	for i, url := range urls {
		wg.Add(1)

		go func(i int, url string) {
			defer wg.Done()

			res, err := Fetch(url)
			if err == nil {
				results[i] = res
			} else {
				log.Println("Error fetching URL:", url, err)
			}
		}(i, url)
	}

	wg.Wait()

	return results
}
