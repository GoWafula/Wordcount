package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordCount(text string) map[string]int {
	words := strings.Fields(text)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}
	return counts
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	counts := wordCount(text)

	fmt.Println("Wordcount: ")
	for word, count := range counts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
