package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"aoc2025/helpers"
)

func opMultiply(ints []int) int {
	product := 1
	for _, i := range ints {
		product *= i
	}
	return product
}

func opAdd(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}

func toStrArr (inputs string) []string {
	out := []string{}
	for i := range inputs {
		out = append(out, string(inputs[i]))
	}
	slices.Reverse(out) // mixing data logic and app logic >>>>>>>
	return out
}

func main() {
	inputs := helpers.MustParseTo(helpers.InputFile(), toStrArr)
	problems := [][][]string{}
	currentPSet := [][]string{}
	for j := range inputs[0] {
		thisCol := []string{}
		for i := range len(inputs)-1 { // don't handle the last line
			thisCol = append(thisCol, inputs[i][j])
		}
		if isBlankLine(thisCol) {
			problems = append(problems, currentPSet)
			currentPSet = [][]string{}
			continue
		}
		currentPSet = append(currentPSet, thisCol)
	}

	// final one
	problems = append(problems, currentPSet)

	intProbs := [][]int{}
	for _, problem := range problems {
		intProb := []int{}
		for _, inte := range problem {
			joined := strings.Join(inte, "")
			joined = strings.ReplaceAll(joined, " ", "")
			ints, _ := strconv.Atoi(joined)
			intProb = append(intProb, ints)
		}
		intProbs = append(intProbs, intProb)
	}
	answer := 0
	for i, op := range strings.Fields(strings.Join(inputs[len(inputs)-1], " ")) {
		switch op {
		case "*":
			answer += opMultiply(intProbs[i])
		case "+":
			answer += opAdd(intProbs[i])
		}
	}
	fmt.Println(answer)
}

func isBlankLine(line []string) bool {
	for _, char := range line {
		if char != " " {
			return false
		}
	}
	return true
}