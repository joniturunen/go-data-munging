package main

import (
	"log"
	"strconv"
	"strings"
)

func parseData(twoDimStrSlice [][]string) []Weather {
	// Remove the last row which is precalculated averages, we dont need them
	// Create a slice of slice of ints to hold the cleaned data
	// Create a slice of Weather structs for the final output
	twoDimStrSlice = twoDimStrSlice[1 : len(twoDimStrSlice)-1]
	var twoDimIntSlice [][]int
	var weatherSlice []Weather

	// Loop through the slice of slices to clean the data (spaces, non-numerical characters)
	for _, record := range twoDimStrSlice {
		cleandStr := strings.Join(strings.Fields(record[0]), " ")
		slicedStr := strings.Split(cleandStr, " ")
		// We only want the first three columns
		// Make a slice of ints to hold the cleaned data
		slicedStr = slicedStr[0:3]
		slicedInts := make([]int, len(slicedStr))
		// Loop through the slice of strings to convert to integers
		for i, str := range slicedStr {
			slicedStr[i] = removeNonNumericalChars(str)
			intVal, err := strconv.Atoi(slicedStr[i])
			if err != nil {
				log.Fatal(err)
			}
			slicedInts[i] = intVal
		}
		// Append the slice of integers to the slice of slices of integers
		twoDimIntSlice = append(twoDimIntSlice, slicedInts)
	}
	// Create a slice of Weather structs
	for _, record := range twoDimIntSlice {
		w := Weather{
			dayNumber: record[0],
			maxTemp:   record[1],
			minTemp:   record[2],
			tempDiff:  record[1] - record[2],
		}
		weatherSlice = append(weatherSlice, w)

	}
	return weatherSlice
}
