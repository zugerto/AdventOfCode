package main

import (
	"bytes"
	"fmt"
)

func process(list []int, sequenceOfLengths []int, pointer int, skipSize int) ([]int, int, int) {
	for _, nrOfElementsToReverse := range sequenceOfLengths {

		list = reverse(list, pointer, nrOfElementsToReverse)

		pointer = (pointer + nrOfElementsToReverse + skipSize) % len(list)
		skipSize++
	}

	return list, pointer, skipSize
}

func reverse(list []int, startPosition int, nrOfElementsToReverse int) []int {
	for {
		if nrOfElementsToReverse <= 0 {
			break
		}

		i := (startPosition + nrOfElementsToReverse - 1) % len(list)
		j := startPosition % len(list)
		list[j], list[i] = list[i], list[j]

		startPosition++
		nrOfElementsToReverse = nrOfElementsToReverse - 2
	}

	return list
}

func hash(list []int, saltedSequence []int, iterations int) []int {
	pointer := 0
	skipSize := 0

	for iterations > 0 {
		list, pointer, skipSize = process(list, saltedSequence, pointer, skipSize)
		iterations--
	}

	return list
}

func salt(input string) []int {
	result := []int{}

	for _, v := range input {
		result = append(result, int(v))
	}

	for _, v := range []int{17, 31, 73, 47, 23} {
		result = append(result, v)
	}

	return result
}

func dense(list []int) []int {
	dense := make([]int, 16)

	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			dense[i] ^= list[i*16+j]
		}
	}

	return dense
}

func hexadecimal(list []int) string {
	var buffer bytes.Buffer

	for _, v := range list {
		buffer.WriteString(fmt.Sprintf("%02x", v))
	}

	return buffer.String()
}

func createList(size int) []int {
	list := make([]int, size)
	for i := range list {
		list[i] = i
	}
	return list
}

func main() {

	//Part A
	list := createList(256)
	sequenceOfLengths := []int{63, 144, 180, 149, 1, 255, 167, 84, 125, 65, 188, 0, 2, 254, 229, 24}
	processed, _, _ := process(list, sequenceOfLengths, 0, 0)
	fmt.Println(processed[0] * processed[1])

	//Part B
	list = createList(256)
	salted := salt("63,144,180,149,1,255,167,84,125,65,188,0,2,254,229,24")
	sparseHash := hash(list, salted, 64)
	denseHash := dense(sparseHash)
	fmt.Println(hexadecimal(denseHash))
}
