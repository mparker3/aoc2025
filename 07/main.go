package main

import (
	"fmt"
	"slices"

	"aoc2025/helpers"
)

func transform(line string) []string {
	chars := []string{}
	for _, char := range line {
		chars = append(chars, string(char))
	}
	return chars
}

func main() {
	lines := helpers.MustParseTo(helpers.InputFile(), transform)
	rows := len(lines)
	cols := len(lines[0])
	ways := make([][]int, rows)
	for i := range ways {
		ways[i] = make([]int, cols)
	}
	start := slices.Index(lines[0], "S")
	ways[0][start] = 1
	for i, line := range lines {
		if i == 0 {
			// we precompute this
			continue
		}
		for j := range line {
			// ways to reach the current position
			// get the sum of the ways to reach all descendants
			if j-1 >= 0 && lines[i-1][j-1] == "^" {
				// all the ways to get to the left descendant if it's a splitter
				ways[i][j] += ways[i-1][j-1]
			}
			if j+1 < len(lines[i]) && lines[i-1][j+1] == "^" {
				// all the ways to get to the right descendant if it's a splitter
				ways[i][j] += ways[i-1][j+1]
			}
			if lines[i-1][j] == "." || lines[i-1][j] == "S" {
				// all the ways to get to the upper position if it's a laser
				ways[i][j] += ways[i-1][j]
			}
		}
	}

	total := 0
	for _, way := range ways[len(ways)-1] {
		total += way
	}
	fmt.Println(total)

}

