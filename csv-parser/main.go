package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	records := csvReader()
	csvWriter("", records)
}

func csvReader() [][]string {
	// Takes a string with the file path and returns a file descriptor to open the file
	csvfile, err := os.Open("./files/test.csv")
	if err != nil {
		log.Fatalln("Couldn't open the CSV file :: ", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	// r := csv.NewReader(bufio.NewReader(csvfile))
	// r.Comma = "\t" // when records are separated by tab
	// r.Comment = "#" // ignores the line starting with '#'

	// Iterate through the records
	// for {
	// 	// Read each record from csv
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("From: %s To: %s Cost: %s\n", record[0], record[1], record[2])
	// }

	// Read all records at once
	records, _ := r.ReadAll()
	fmt.Println(records)
	return records
}

func csvWriter(path string, csvData [][]string) {
	dataVals := [][]string{
		{"LDB", "GRU", "30"},
		{"GRU", "LDB", "60"},
	}

	if path == "" {
		path = "./files/routes.csv"
	}

	if len(csvData) == 0 {
		csvData = dataVals
	} else {
		for _, item := range dataVals {
			csvData = append(csvData, item)
		}
	}

	csvfile, err := os.Create(path)
	if err != nil {
		log.Fatalln("Couldn't create the CSV file :: ", err)
	}

	w := csv.NewWriter(csvfile)

	err = w.WriteAll(csvData)
	if err != nil {
		log.Fatalln("Couldn't write the CSV file :: ", err)
	}
}
