package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Takes a string with the file path and returns a file descriptor to open the file
	csvfile, err := os.Open("test.csv")
	if err != nil {
		log.Fatalln("Couldn't open the dsv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	// r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the recods
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("From: %s To: %s Cost: %s\n", record[0], record[1], record[2])
	}
}
