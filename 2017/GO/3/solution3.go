package main

import (
	"fmt"
	"math"
)

func SpiralSteps(steps int) int {

	x := 0.0
	y := 0.0
	nrOfTurns := 0

	for step := 0; step < steps-1; {

		stepsUntilTurn := (nrOfTurns / 2) + 1
		directionOfTurn := nrOfTurns % 4

		for i := 1; i < stepsUntilTurn; i++ {
			if step == steps-1 {
				break
			}
			step++
			if directionOfTurn == 0 {
				x -= 1
			} else if directionOfTurn == 1 {
				y -= 1
			} else if directionOfTurn == 2 {
				x += 1
			} else {
				y += 1
			}
		}
		nrOfTurns += 1
	}

	return int(math.Abs(x) + math.Abs(y))

}

type (
	position struct {
		x int
		y int
	}
	positionToValue map[position]int
)

func NeighbourSum(steps int) int {

	var positionToNeighbourSum positionToValue
	positionToNeighbourSum = make(positionToValue)

	p := position{x: 0, y: 0}
	positionToNeighbourSum[p] = 1

	nrOfTurns := 0

	for step := 0; step < steps-1; {

		stepsUntilTurn := (nrOfTurns / 2) + 1
		directionOfTurn := nrOfTurns % 4

		for i := 1; i < stepsUntilTurn; i++ {
			if step == steps-1 {
				break
			}
			step++
			if directionOfTurn == 0 {
				p = position{x: (p.x - 1), y: p.y}
			} else if directionOfTurn == 1 {
				p = position{x: p.x, y: (p.y - 1)}
			} else if directionOfTurn == 2 {
				p = position{x: (p.x + 1), y: p.y}
			} else {
				p = position{x: p.x, y: (p.y + 1)}
			}
			neighbourSum := CalculateNeighbourSum(p, positionToNeighbourSum)

			if neighbourSum > steps {
				return neighbourSum
			}

			positionToNeighbourSum[p] = neighbourSum

		}
		nrOfTurns += 1
	}

	return positionToNeighbourSum[p]

}

func CalculateNeighbourSum(p position, positionToNeighbourSum positionToValue) int {

	summa := 0
	summa += positionToNeighbourSum[position{x: (p.x + 1), y: p.y}]
	summa += positionToNeighbourSum[position{x: (p.x + 1), y: (p.y + 1)}]
	summa += positionToNeighbourSum[position{x: p.x, y: (p.y + 1)}]
	summa += positionToNeighbourSum[position{x: (p.x - 1), y: (p.y + 1)}]
	summa += positionToNeighbourSum[position{x: (p.x - 1), y: p.y}]
	summa += positionToNeighbourSum[position{x: (p.x - 1), y: (p.y - 1)}]
	summa += positionToNeighbourSum[position{x: p.x, y: (p.y - 1)}]
	summa += positionToNeighbourSum[position{x: (p.x + 1), y: (p.y - 1)}]

	return summa
}

func main() {
	fmt.Println("Spiralsteps: ", SpiralSteps(325489))

	fmt.Println("NeighbourSum: ", NeighbourSum(325489))
}
