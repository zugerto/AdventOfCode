package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	register  string
	increase  bool
	amount    int
	condition condition
}

type condition struct {
	register   string
	comparator string
	value      int
}

type registerValues map[string]int

type memory struct {
	registerValues  registerValues
	higestTempValue int
}

func (c condition) evaluate(m *memory) bool {
	value := m.get(c.register)
	switch c.comparator {
	case "==":
		return value == c.value
	case ">":
		return value > c.value
	case "<":
		return value < c.value
	case ">=":
		return value >= c.value
	case "<=":
		return value <= c.value
	case "!=":
		return value != c.value
	default:
		return false
	}
}

func (m *memory) get(register string) int {
	value, ok := m.registerValues[register]
	if ok {
		return value
	}
	m.set(register, 0)
	return 0
}

func (m *memory) set(register string, value int) {
	if value > m.higestTempValue {
		m.higestTempValue = value
	}
	m.registerValues[register] = value
}

func (m *memory) increment(register string, amount int) {
	current := m.get(register)
	m.set(register, current+amount)
}

func (m *memory) decrement(register string, amount int) {
	m.increment(register, -amount)
}

func EvaluateInstructions(instructions []instruction) *memory {
	memory := &memory{
		registerValues:  make(registerValues),
		higestTempValue: 0,
	}

	for _, instruction := range instructions {
		if instruction.condition.evaluate(memory) {
			if instruction.increase {
				memory.increment(instruction.register, instruction.amount)
			} else {
				memory.decrement(instruction.register, instruction.amount)
			}
		}
	}
	return memory
}

func ParseInputRow(partsInRow []string) instruction {
	increase := true
	if partsInRow[1] == "inc" {
		increase = true
	} else {
		increase = false
	}

	amount, _ := strconv.Atoi(partsInRow[2])
	value, _ := strconv.Atoi(partsInRow[6])

	return instruction{
		register: partsInRow[0],
		increase: increase,
		amount:   amount,
		condition: condition{
			register:   partsInRow[4],
			comparator: partsInRow[5],
			value:      value,
		},
	}
}

func HigestValue(m *memory) int {
	highestValue := 0
	for _, v := range m.registerValues {
		if v > highestValue {
			highestValue = v
		}
	}
	return highestValue
}

func HigestTempValue(m *memory) int {
	return m.higestTempValue
}

func main() {
	file, err := os.Open("input8.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	instructions := []instruction{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		partsInRow := strings.Fields(scanner.Text())
		instructions = append(instructions, ParseInputRow(partsInRow))
	}

	memory := EvaluateInstructions(instructions)
	highestValue := HigestValue(memory)
	highestTempValue := HigestTempValue(memory)
	fmt.Printf("HigestValue: %d\n", highestValue)
	fmt.Printf("HighestTempValue: %d\n", highestTempValue)
}
