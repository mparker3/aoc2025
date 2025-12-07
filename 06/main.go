package main

import (
	"fmt"
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


func main() {
	inputs := helpers.MustParseTo(helpers.InputFile(), strings.Fields)
	fmt.Println(inputs[0])
	numInputs := inputs[:len(inputs)-1]
	ops := inputs[len(inputs)-1]
	numSeries := [][]int{}
	for i := range numInputs[0] {
		thisSeries := []int{}
		for j := range numInputs {
			toInt, _ := strconv.Atoi(numInputs[j][i])
			thisSeries = append(thisSeries, toInt)
		}
		numSeries = append(numSeries, thisSeries)
	}
	result := 0
	for i, series := range numSeries {
		op := ops[i]
		if op == "*" {
			result += opMultiply(series)
		} else {
			result += opAdd(series)
		}
	}
	fmt.Println(result)
}
