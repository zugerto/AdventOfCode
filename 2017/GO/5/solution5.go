package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func EscapeMaze(maze []int) int {
	pointer := 0
	stepsToEscapeMaze := 0

	for {
		if pointer > len(maze)-1 {
			break
		}
		stepsToEscapeMaze++
		instruction := maze[pointer]
		maze[pointer]++
		pointer += instruction
	}
	return stepsToEscapeMaze
}

func EscapeMazeStrangeOffset(maze []int) int {
	pointer := 0
	stepsToEscapeMaze := 0

	for {
		if pointer > len(maze)-1 {
			break
		}
		stepsToEscapeMaze++

		instruction := maze[pointer]
		if instruction >= 3 {
			maze[pointer]--
		} else {
			maze[pointer]++
		}

		pointer += instruction
	}

	return stepsToEscapeMaze
}

func main() {

	file, err := os.Open("input5.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	maze := []int{}
	strangerMaze := []int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		instruction, _ := strconv.Atoi(scanner.Text())
		maze = append(maze, instruction)
		strangerMaze = append(strangerMaze, instruction)
	}

	fmt.Println(EscapeMaze(maze))
	fmt.Println(EscapeMazeStrangeOffset(strangerMaze))
}
