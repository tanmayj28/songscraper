package main

import (
	"fmt"
	"log"
	config "songscraper/configs"
	scraper "songscraper/internal/scraper"
)

func main() {
	// Load Config
	configPath := "configs/config.yaml"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Error loading config: %w", err)
	}

	// Scrape Songs
	songs, err := scraper.ScrapeSongs(cfg.URL)
	if err != nil {
		log.Fatalf("Error scraping songs: %w", err)
	}

	// Print Songs
	for _, song := range songs {
		fmt.Printf("Rank: %d Title: %s Artist: %s\n", song.Rank, song.Title, song.Artist)
	}
}
