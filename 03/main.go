package main

import (
	"aoc2025/helpers"
	"fmt"
	"math"
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
		joltage := findMaxJoltageRecursive(battery.joltages, 12)
		maxJoltage += joltage
	}
	fmt.Println(maxJoltage)
}


func findMaxJoltageRecursive(joltages []int, digitsLeft int) int {
	// so to expand to 12 digits. you essentially want the first instance of the highest number that is not in the last 11 digits.
	// then repeat the process. now, you want the first instance of the highest number that's in the remainder of the array after the index of the first number, 
	// but still not in the last 10 digits. and so on and so forth.  
	// solve the base case
	if digitsLeft == 0 {
		return 0
	}
	
	// find the first instance of the highest number, work backwards
	highestNumber, highestNumberIndex := -1, -1
	for i := 0; i <= len(joltages) - digitsLeft; i++ {
		if joltages[i] > highestNumber {
			highestNumber = joltages[i]
			highestNumberIndex = i
		}
	}
	return (highestNumber * int(math.Pow10(digitsLeft - 1))) + findMaxJoltageRecursive(joltages[highestNumberIndex + 1:], digitsLeft - 1)
}