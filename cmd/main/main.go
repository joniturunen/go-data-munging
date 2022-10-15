package main

import (
	"log"
)

type Weather struct {
	dayNumber int
	maxTemp   int
	minTemp   int
	tempDiff  int
}

func main() {
	// Read the data from the file, returns a slice of slices of strings
	rawData := readFile()
	// Parse the data, returns a slice of Weather structs
	w := parseData(rawData)

	resultDay, resultTempDiff := getLowestTempDiff(w)
	log.Printf("The day with the lowest temperature difference is day %d with a difference of %d", resultDay, resultTempDiff)
}

// Get the lowest temperature difference and return the day number
func getLowestTempDiff(w []Weather) (int, int) {
	var lowestTempDiff int
	var lowestTempDiffDay int
	for i, record := range w {
		if i == 0 {
			lowestTempDiff = record.tempDiff
			lowestTempDiffDay = record.dayNumber
		} else if record.tempDiff < lowestTempDiff {
			lowestTempDiff = record.tempDiff
			lowestTempDiffDay = record.dayNumber
		}
	}
	return lowestTempDiffDay, lowestTempDiff
}
