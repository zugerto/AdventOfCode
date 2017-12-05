package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func SolutionA(digits []string) int {
	sum := 0

	for i, v := range digits {
		if i > 0 && digits[i-1] == v {
			i, _ := strconv.Atoi(v)
			sum += i
		}
	}

	if digits[0] == digits[len(digits)-1] {
		i, _ := strconv.Atoi(digits[0])
		sum += i
	}

	return sum
}

func SolutionB(digits []string) int {
	sum := 0

	for i, v := range digits {
		if digits[(i+(len(digits)/2))%len(digits)] == v {
			i, _ := strconv.Atoi(v)
			sum += i
		}
	}

	return sum
}

func main() {
	fileContent, err := ioutil.ReadFile("input1.txt")
	if err != nil {
		fmt.Print(err)
	}

	digits := strings.Split(string(fileContent), "")

	fmt.Println(SolutionA(digits))
	fmt.Println(SolutionB(digits))
}
