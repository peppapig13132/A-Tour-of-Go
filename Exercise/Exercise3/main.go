package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)           // Split string into words
	wordCountMap := make(map[string]int) // Create an empty map

	for _, word := range words {
		wordCountMap[word]++ // Count occurrences of each word
	}

	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
