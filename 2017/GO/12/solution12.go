package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func ParseInputRow(row string) (int, []int) {
	split := strings.Split(row, " <-> ")

	program, _ := strconv.Atoi(strings.Trim(split[0], " "))

	foo := []int{}
	programs := strings.Split(split[1], ",")

	for _, v := range programs {
		bar, _ := strconv.Atoi(strings.Trim(v, " "))
		foo = append(foo, bar)
	}

	return program, foo
}

func alreadySeenIn(program int, programs []int) bool {
	for _, p := range programs {
		if p == program {
			return true
		}
	}
	return false
}

func Connect(programs map[int][]int, program int, connectedPrograms []int) {
	_, ok := programs[program]
	if ok {
		programs[program] = append(programs[program], connectedPrograms...)
	} else {
		programs[program] = connectedPrograms
	}
}

func CountA(programs map[int][]int) int {
	queue := []int{}
	connected := []int{}

	current := 0
	queue = append(queue, current)

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if !alreadySeenIn(current, connected) {
			connected = append(connected, current)
		}

		for _, program := range programs[current] {
			if !alreadySeenIn(program, connected) {
				queue = append(queue, program)
			}
		}
	}

	return len(connected)
}

func CountB(programs map[int][]int) int {

	groups := []int{}
	nrOfGroups := 0

	for v := range reflect.ValueOf(programs).MapKeys() {

		if !alreadySeenIn(v, groups) {
			nrOfGroups++

			queue := []int{}
			connected := []int{}

			current := 0
			queue = append(queue, current)

			for len(queue) > 0 {
				current, queue = queue[0], queue[1:]

				if !alreadySeenIn(current, connected) {
					connected = append(connected, current)
				}

				groups = append(groups, current)

				for _, program := range programs[current] {
					if !alreadySeenIn(program, connected) {
						queue = append(queue, program)
					}
				}
			}
		}
	}

	return nrOfGroups
}

func main() {

	file, err := os.Open("input12.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	programs := make(map[int][]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		program, connectedPrograms := ParseInputRow(scanner.Text())
		Connect(programs, program, connectedPrograms)
	}

	fmt.Printf("CountA: %d\n", CountA(programs))
	fmt.Printf("CountB: %d\n", CountB(programs))

}
