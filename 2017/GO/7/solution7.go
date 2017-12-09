package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type program struct {
	name              string
	weight            int
	subProgramsWeight int
	subPrograms       []string
	parent            string
}

func ParseInputRow(row string) (string, *program) {
	re := regexp.MustCompile(`([a-z]+) \((\d+)\)(?: -> )?(.*)?`)
	matches := re.FindAllStringSubmatch(row, -1)

	name := matches[0][1]
	weight, _ := strconv.Atoi(matches[0][2])
	var programs []string
	if matches[0][3] == "" {
		programs = make([]string, 0)
	} else {
		programs = strings.Split(matches[0][3], ", ")
	}

	return name,
		&program{
			name,
			weight,
			0,
			programs,
			"",
		}
}

func CalculateParents(programs map[string]*program) {

	for _, program := range programs {
		for _, programName := range program.subPrograms {
			c := programs[programName]
			c.parent = program.name
		}
	}
}

func RootNodeName(programs map[string]*program) string {

	for _, n := range programs {
		if n.parent == "" {
			return n.name
		}
	}

	return ""
}

func CalculateSubWeight(programs map[string]*program, currentProgram *program) int {
	for _, subProgram := range currentProgram.subPrograms {
		currentProgram.subProgramsWeight += CalculateSubWeight(programs, programs[subProgram])
	}
	return currentProgram.weight + currentProgram.subProgramsWeight
}

func CalculateValueToBalanceFirstUnbalanced(programs map[string]*program) {

	//Eeee, try to ignore the solution to this part of the problem.

	for _, program := range programs {
		if program.subProgramsWeight == 0 {
			continue
		}

		fmt.Printf("name: %s, weight: %d\n", program.name, program.weight)

		for _, subProgramName := range program.subPrograms {
			subProgram := programs[subProgramName]
			fmt.Printf("   name: %s, weight: %d, totalWeight: %d\n", subProgram.name, subProgram.weight, subProgram.weight+subProgram.subProgramsWeight)
		}

	}
}

func main() {

	file, err := os.Open("input7.txt")
	//file, err := os.Open("input7-small.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	programs := make(map[string]*program)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		name, program := ParseInputRow(scanner.Text())
		programs[name] = program
	}

	CalculateParents(programs)

	rootNodeName := RootNodeName(programs)
	fmt.Println(rootNodeName)

	CalculateSubWeight(programs, programs[rootNodeName])
	CalculateValueToBalanceFirstUnbalanced(programs)
}
