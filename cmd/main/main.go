package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type Weather struct {
	dayNumber int
	maxTemp   int
	minTemp   int
}

func main() {
	sliceOfSliceOfStr := readFile()

	for _, record := range sliceOfSliceOfStr {
		newStr := strings.Join(strings.Fields(record[0]), " ")
		splittedStr := strings.Split(newStr, " ")
		log.Printf("Splitted string is: %T\n", splittedStr)
		// TODO: convert splittedStr to int values
		// TODO: create a Weather struct with the values
		// TODO: find the day with the smallest temperature spread
	}
}

func readFile() [][]string {
	file, err := os.Open("weather.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// strings.Join(strings.Fields(file), " ")

	log.Printf("File type is: %T", file)

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	rawData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return rawData
}
