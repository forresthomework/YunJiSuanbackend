package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	currentWord := ""
	currentCounts := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			word := parts[0]
			docID := parts[1]
			if currentWord != "" && currentWord != word {
				printCounts(currentWord, currentCounts)
				currentCounts = make(map[string]int)
			}
			currentWord = word
			currentCounts[docID]++
		}
	}
	if currentWord != "" {
		printCounts(currentWord, currentCounts)
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", scanner.Err())
		os.Exit(1)
	}
}

func printCounts(word string, counts map[string]int) {
	var pairs []string
	for docID, count := range counts {
		docID = strings.TrimSpace(docID)
		pairs = append(pairs, fmt.Sprintf("(%s,%d)", docID, count))
	}
	fmt.Printf("%s %s\n", word, strings.Join(pairs, ","))
}

