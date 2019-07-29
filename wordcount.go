package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		m[word]++
	}

	return m
}

func main() {
	fmt.Println(WordCount("Ahmet Mehmet Ahmet Cemil Muco Muco Mehmet"))
}
