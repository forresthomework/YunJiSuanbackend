package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ",2)
		if len(parts) == 2 {
			docID := parts[0]
			text := parts[1]
			words := strings.Split(text, " ")
			for _, word := range words {
				fmt.Printf("%s %s\n", word, docID)
			}
		}
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", scanner.Err())
		os.Exit(1)
	}
}

