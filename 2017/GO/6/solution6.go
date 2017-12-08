package main

import (
	"fmt"
)

func ReallocationCount(banks []int) (int, int) {

	cycles := 0
	seenBanks := make(map[string]int)

	for {
		duplicateBanksFound, betweenOccurances := HasSeenBanks(banks, seenBanks, cycles)
		if duplicateBanksFound {
			return cycles, betweenOccurances
		}

		highestIndex, highestValue := Highest(banks)

		banks[highestIndex] = 0
		index := highestIndex + 1

		for {
			if highestValue == 0 {
				break
			}

			if index >= len(banks) {
				index = 0
			}

			banks[index]++
			highestValue--
			index++
		}

		cycles++
	}
}

func Highest(banks []int) (int, int) {
	highestIndex := 0
	highestValue := 0
	for i, v := range banks {
		if v > highestValue {
			highestIndex = i
			highestValue = v
		}
	}
	return highestIndex, highestValue
}

func HasSeenBanks(banks []int, seenBanks map[string]int, cycles int) (bool, int) {
	hash := fmt.Sprint(banks)
	if previousCycle, found := seenBanks[hash]; found {
		return true, (cycles - previousCycle)
	}
	seenBanks[hash] = cycles
	return false, 0
}

func main() {
	fmt.Println(ReallocationCount([]int{14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4}))
}
