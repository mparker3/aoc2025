package main

import (
	"fmt"

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

type Board [][]int

func (b *Board) neighbors(x int, y int) []int {
	neighborIndices := [][]int{
		{x-1, y},
		{x+1, y},
		{x, y-1},
		{x, y+1},
		{x-1, y-1},
		{x+1, y-1},
		{x-1, y+1},
		{x+1, y+1},
	}
	neighbors := []int{}
	for _, neighborIndex := range neighborIndices {
		if neighborIndex[0] < 0 || neighborIndex[0] >= len(*b) || neighborIndex[1] < 0 || neighborIndex[1] >= len((*b)[0]) {
			continue
		}
		neighbors = append(neighbors, (*b)[neighborIndex[0]][neighborIndex[1]])
	}
	return neighbors
}

func main() {
	board := Board(helpers.MustParseTo(helpers.InputFile(), transform))
	reachable := 0
	for i := range len(board) {
		for j := range len(board[i]) {
			if board[i][j] == 1 {
				neighbors := board.neighbors(i, j)
				neighborSum := 0
				for _, neighbor := range neighbors {
					neighborSum += neighbor
				}
				if neighborSum < 4 {
					reachable++
				}
			}
		}
	}
	fmt.Println(reachable)
}
