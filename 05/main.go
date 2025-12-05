package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2025/helpers"
)

func transform(line string) []int {
	numbers := []int{}
	for _, char := range line {
		if char == '.' {
			numbers = append(numbers, 0)
		} else {
			numbers = append(numbers, 1)
		}
	}
	return numbers
}

type Range [2]int

func main() {
	fl := helpers.MustReadLines(helpers.InputFile())
	splitIdx := 0
	for i, line := range fl {
		if line == "" {
			splitIdx = i
			break
		}
	}
	ranges := fl[:splitIdx]
	ingredients := fl[splitIdx+1:]
	parsedRanges := []Range{}
	for _, rng := range ranges {
		parts := strings.Split(rng, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])
		parsedRanges = append(parsedRanges, Range{min, max})
	}
	parsedIngredients := []int{}
	for _, ingredient := range ingredients {
		parsedIngredient, _ := strconv.Atoi(ingredient)
		parsedIngredients = append(parsedIngredients, parsedIngredient)
	}
	fmt.Println(parsedRanges)
	fmt.Println(parsedIngredients)
	has := 0
	for _, ingredient := range parsedIngredients {
		fmt.Println(ingredient, "is being checked")
		for _, rng := range parsedRanges {

			if ingredient >= rng[0] && ingredient <= rng[1] {
				has++
				fmt.Println(ingredient, "has", rng)
				break
			}
		}
	}
	fmt.Println(has)
}