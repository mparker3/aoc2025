package main

import (
	"fmt"
	"sort"
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
	rawRanges := fl[:splitIdx]
	parsedRanges := []Range{}
	for _, rng := range rawRanges {
		parts := strings.Split(rng, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])
		parsedRanges = append(parsedRanges, Range{min, max})
	}
	// let's find the total unique elements in the parsed ranges
	// get the first range
	sort.Slice(parsedRanges, func(i, j int) bool {
		return parsedRanges[i][0] < parsedRanges[j][0]
	})
	sortedRanges := parsedRanges[1:]
	finalRanges := []Range{parsedRanges[0]}
	for _, rng := range sortedRanges {
		// if no overlap with the final range, add a new one
		if rng[0] > finalRanges[len(finalRanges)-1][1] {
			finalRanges = append(finalRanges, rng)
		} else {
			// we're overlapping with the current high range. We know we can't go lower since we sorted. go higher
			finalRanges[len(finalRanges)-1][1] = max(finalRanges[len(finalRanges)-1][1], rng[1])
		}
	}
	sum := 0
	for _, rng := range finalRanges {
		sum += rng[1] - rng[0] + 1
	}
	fmt.Println(sum)
}