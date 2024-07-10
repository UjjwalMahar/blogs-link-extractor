package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/xuri/excelize/v2"
)

func main() {
	// Create a new Excel file
	f := excelize.NewFile()

	// Create a new sheet
	sheetName := "Links"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		log.Fatal("Error creating new sheet:", err)
	}

	// Set headers for the columns
	f.SetCellValue(sheetName, "A1", "Name")
	f.SetCellValue(sheetName, "B1", "URL")

	row := 2

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

			// Write the link and its name to the Excel file
			cellName := fmt.Sprintf("A%d", row)
			cellURL := fmt.Sprintf("B%d", row)
			f.SetCellValue(sheetName, cellName, name)
			f.SetCellValue(sheetName, cellURL, href)
			row++
		})

		fmt.Printf("Scraped page %d\n", pageNumber)
	}

	// Set the active sheet
	f.SetActiveSheet(index)

	// Save the Excel file
	if err := f.SaveAs("links.xlsx"); err != nil {
		log.Fatal("Error saving file:", err)
	}

	fmt.Println("Scraping completed. Links saved to links.xlsx file.")
}
