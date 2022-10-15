package main

import (
	"encoding/csv"
	"log"
	"os"
)

func readFile() [][]string {
	file, err := os.Open("weather.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(file)

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	rawData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return rawData
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
