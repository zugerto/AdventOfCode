package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type position struct {
	x int
	y int
}

func CalculateNrOfStepsToChild(steps []string) (int, int) {

	position := position{x: 0, y: 0}
	max := 0

	for _, step := range steps {
		switch step {
		case "s":
			position.x, position.y = position.x, position.y-1
		case "n":
			position.x, position.y = position.x, position.y+1
		case "sw":
			position.x, position.y = position.x-1, position.y-1
		case "nw":
			position.x, position.y = position.x-1, position.y+1
		case "se":
			position.x, position.y = position.x+1, position.y-1
		case "ne":
			position.x, position.y = position.x+1, position.y+1
		}
		dist := Max(Abs(position.x), Abs(position.y))
		if dist > max {
			max = dist
		}
	}

	position.x, position.y = Abs(position.x), Abs(position.y)

	return Max(position.x, position.y), max
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fileContent, err := ioutil.ReadFile("input11.txt")
	if err != nil {
		fmt.Print(err)
	}

	steps := strings.Split(string(fileContent), ",")
	nrOfSteps, max := CalculateNrOfStepsToChild(steps)
	fmt.Printf("NrOfSteps: %d, Max: %d\n", nrOfSteps, max)
}
