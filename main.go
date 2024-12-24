/*
Roboform to Kaspersky Password Manager CSV converter

Usage:

	go run main.go

App Versions:
  - Roboform: Version 9.6.3.3 (en-english, RUS) Oct 16 2024 09:12:00
  - Kaspersky Password Manager: 24.3.0.327 / 1731938737_7632 / 1.1.0.2_1
*/

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

const (
	INPUT_FILENAME  = "Roboform_Backup_Export_Logins_Only.csv"
	OUTPUT_FILENAME = "output.csv"
)

var KASPERSKY_CSV_HEADERS = []string{"url", "username", "password", "name", "extra", "folder"}

func main() {

	// open input csv file
	inputCSV, err := os.Open(INPUT_FILENAME)
	if err != nil {
		log.Fatalf("Cannot open input csv file: %v\n", err)
	}
	defer inputCSV.Close()

	// create output csv file
	outputCSV, err := os.Create(OUTPUT_FILENAME)
	if err != nil {
		log.Fatalf("Cannot create output csv file: %v\n", err)
	}
	defer outputCSV.Close()

	// setup csv parser custom cases
	reader := csv.NewReader(inputCSV)
	reader.LazyQuotes = true    // Allows lazy quotes
	reader.FieldsPerRecord = -1 // Allows variable number of fields
	// reader.TrimLeadingSpace = true // Allows leading spaces

	// Roboform csv format (exported from Roboform)
	// Name,Url,MatchUrl,Login,Pwd,Note,Folder,RfFieldsV2

	// Kaspersky Password manager csv format
	// url,username,password,name,extra
	// or
	// "Account","Login Name","Password","Web Site","Comments"??

	// write headers into output file
	writer := csv.NewWriter(outputCSV)
	err = writer.Write(KASPERSKY_CSV_HEADERS)
	if err != nil {
		log.Fatalf("Error writing headers: %v\n", err)
	}
	defer writer.Flush()

	// read input file
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}

	// write rows into output file
	// ignore header
	for _, row := range rows[1:] {

		// rows mapping
		url := row[1]
		username := row[3]
		password := row[4]
		name := row[0]
		note := row[5]
		folder := row[6]
		// matchurl := row[2]
		// rffieldsv2 := row[7]

		// DEBUG LOG:
		// fmt.Printf("URL: %s, Username: %s, Password: %s, Name: %s, Note: %s, Folder: %s\n", url, username, password, name, note, folder)

		// write row
		err = writer.Write([]string{url, username, password, name, note, folder})
		if err != nil {
			fmt.Printf("Error writing row: %v\n", err)
		}
	}

	if err := writer.Error(); err != nil {
		log.Fatalf("Error flushing writer: %v\n", err)
	}

	// success log
	fmt.Println("✅ Converting has been successfully done!")
}
