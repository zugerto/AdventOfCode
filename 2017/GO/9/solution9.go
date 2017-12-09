package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Calculate(characters []string) (int, int) {
	score := 0
	depth := 0
	insideGarbage := false
	garbage := 0

	i := 0
	for i < len(characters) {
		character := characters[i]

		switch character {
		case "{":
			if insideGarbage {
				i++
				garbage++
			} else {
				depth++
				i++
			}
			continue
		case "}":
			if insideGarbage {
				i++
				garbage++
			} else {
				score += depth
				depth--
				i++
			}
			continue
		case "!":
			i += 2
			continue
		case "<":
			if insideGarbage {
				garbage++
			}
			insideGarbage = true
			i++
			continue
		case ">":
			insideGarbage = false
			i++
			continue
		}

		if insideGarbage {
			garbage++
		}

		i++
	}

	return score, garbage
}

func main() {
	fileContent, err := ioutil.ReadFile("input9.txt")
	if err != nil {
		fmt.Print(err)
	}

	characters := strings.Split(string(fileContent), "")
	score, garbage := Calculate(characters)
	fmt.Printf("Score: %d, Garbage: %d\n", score, garbage)
}
