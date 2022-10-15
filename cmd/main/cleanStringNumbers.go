package main

import (
	"log"
	"regexp"
)

// removeNonNumericalChars removes all non numerical characters from a string
func removeNonNumericalChars(s string) string {
	// Check if string starts with numerical value
	if numVal, _ := regexp.MatchString(`^\d`, s); numVal {
		reg, err := regexp.Compile("[^0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		return reg.ReplaceAllString(s, "")
	}
	return s
}
