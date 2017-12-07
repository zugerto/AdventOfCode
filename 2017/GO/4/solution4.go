package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func ValidPassphrase(wordsInRow []string) int {

	seenWords := make(map[string]int)

	for _, word := range wordsInRow {
		seenWords[word]++
		if seenWords[word] > 1 {
			return 0
		}
	}

	return 1
}

func ValidPassphraseAlsoConsiderAnagrams(wordsInRow []string) int {

	seenWords := make(map[string]int)

	for _, word := range wordsInRow {
		sortedWord := SortWord(word)
		seenWords[sortedWord]++
		if seenWords[sortedWord] > 1 {
			return 0
		}
	}

	return 1
}

func SortWord(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {

	file, err := os.Open("input4.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	sumValidPassphrases := 0
	sumValidPassphrasesAlsoConsiderAnagrams := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		wordsInRow := strings.Fields(scanner.Text())
		sumValidPassphrases += ValidPassphrase(wordsInRow)
		sumValidPassphrasesAlsoConsiderAnagrams += ValidPassphraseAlsoConsiderAnagrams(wordsInRow)
	}

	fmt.Println(sumValidPassphrases)
	fmt.Println(sumValidPassphrasesAlsoConsiderAnagrams)
}
