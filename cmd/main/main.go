package main

import (
	"log"
	"regexp"
	"strings"
)

type Weather struct {
	dayNumber int
	maxTemp   int
	minTemp   int
}

func main() {
	rawData := readFile()
	parseData(rawData)

}

func parseData(twoDimSlice [][]string) []Weather {
	// Get heading from first row
	h := twoDimSlice[0]
	// Make a slice of the headings
	h[0] = strings.Join(strings.Fields(h[0]), " ")
	// Use just the first three headings
	heading := strings.Split(h[0], " ")[:3]
	// Print out the headings for reference
	log.Println(heading)
	// Remove the last row which is precalculated averages
	twoDimSlice = twoDimSlice[1 : len(twoDimSlice)-1]
	for _, record := range twoDimSlice {
		cleandStr := strings.Join(strings.Fields(record[0]), " ")
		slicedStr := strings.Split(cleandStr, " ")
		slicedStr = slicedStr[0:3]
		for i, str := range slicedStr {
			slicedStr[i] = removeNonNumericalChars(str)
			// log.Println(slicedStr[i])
			// TODO: Convert to int
		}
		log.Println(slicedStr)
		// log.Printf("Splitted string is: %T\n", splittedStr)
		// TODO: create a Weather struct with the values
		// TODO: find the day with the smallest temperature spread
	}
	// TODO: return weather struct
	return nil
}

func removeNonNumericalChars(s string) string {
	// Check if string starts with numerical value
	if numVal, _ := regexp.MatchString(`^\d`, s); numVal {
		reg, err := regexp.Compile("[^0-9]+")
		if err != nil {
			log.Fatal(err)
		}
	}
	return s
}
