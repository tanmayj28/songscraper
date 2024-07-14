package scraper

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Song struct {
	Title  string
	Artist string
	Rank   int
}

func ScrapeSongs(url string) ([]Song, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Status Code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var songs []Song
	doc.Find(".chart-items .chart-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".description .chart-name span").Text()
		artist := s.Find(".description .chart-artist").Text()
		rank := i + 1

		songs = append(songs, Song{Title: title, Artist: artist, Rank: rank})
	})

	return songs, nil
}
