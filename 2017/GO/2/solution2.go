package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Checksum(digits []string) int {
	highest := 0
	lowest := math.MaxInt32

	for _, v := range digits {
		digit, _ := strconv.Atoi(v)

		if digit < lowest {
			lowest = digit
		}

		if digit > highest {
			highest = digit
		}
	}

	return highest - lowest
}

func EvenlyDivide(digits []string) int {
	result := 0

	for i, v := range digits {
		digit, _ := strconv.Atoi(v)
		for k, z := range digits {
			anotherDigit, _ := strconv.Atoi(z)
			if digit%anotherDigit == 0 && k != i {
				result = int(digit / anotherDigit)
			}
		}
	}

	return result
}

func main() {

	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	sumChecksum := 0
	sumEvenlyDivide := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		digits := strings.Fields(scanner.Text())
		sumChecksum += Checksum(digits)
		sumEvenlyDivide += EvenlyDivide(digits)
	}

	fmt.Println(sumChecksum)
	fmt.Println(sumEvenlyDivide)
}
