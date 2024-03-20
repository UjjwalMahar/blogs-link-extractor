package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Create or open a text file to save the links
	file, err := os.Create("links.txt")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	// Iterate over page numbers from 1 to 4
	for pageNumber := 1; pageNumber <= 4; pageNumber++ {
		url := fmt.Sprintf("https://cloudzenia.com/blog/page/%d", pageNumber)

		// Make HTTP GET request to the website
		response, err := http.Get(url)
		if err != nil {
			log.Printf("Error fetching URL %s: %s", url, err)
			continue
		}
		defer response.Body.Close()

		// Parse the HTML response
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Printf("Error parsing HTML from URL %s: %s", url, err)
			continue
		}

		// Find all links within the desired div and extract their names and URLs
		doc.Find(".theimran-post-layout-one__title").Each(func(i int, s *goquery.Selection) {
			link := s.Find("h3 a")
			href, _ := link.Attr("href")
			name := link.Text()

			// Write the link and its name to the text file
			fmt.Fprintf(file, "%s\t%s\n", name, href)
		})

		fmt.Printf("Scraped page %d\n", pageNumber)
	}

	fmt.Println("Scraping completed. Links saved to links.txt file.")
}
