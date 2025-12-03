package main

import (
	"aoc2025/helpers"
	"fmt"
	"strconv"
)

// TODO(mparker): set up gofmt on save
type battery struct {
	joltages []int
}

func transform(line string) battery {
	joltages := []int{}
	for _, char := range line {
		intJoltage, _ := strconv.Atoi(string(char))
		joltages = append(joltages, intJoltage)
	}
	return battery{
		joltages: joltages,
	}
}

func main() {
	batteries := helpers.MustParseTo(helpers.InputFile(), transform)
	maxJoltage := 0
	for _, battery := range batteries {
		maxJoltage += findMaxJoltage(battery)
	}
	fmt.Println(maxJoltage)
}

func findMaxJoltage(battery battery) int {
	// find max first digit. work backwards from len(battery.joltages) - 2
	maxFirstDigit, maxFirstDigitIndex := -1, -1 // zero is a valid digit, but it might not appear here 
	for i := len(battery.joltages) - 2; i >= 0; i-- {
		if battery.joltages[i] >= maxFirstDigit { // >= because we want to give ourselves as many options as possible
			maxFirstDigit = battery.joltages[i]
			maxFirstDigitIndex = i
		}
	}

	maxLastDigit := -1
	// work forwards from maxFirstDigitIndex + 1
	for i := maxFirstDigitIndex + 1; i < len(battery.joltages); i++ {
		if battery.joltages[i] > maxLastDigit {
			maxLastDigit = battery.joltages[i]
		}
	}
	return (maxFirstDigit * 10) + maxLastDigit
}