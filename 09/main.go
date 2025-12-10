package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aoc2025/helpers"
)

type Point struct {
	x int
	y int
}
func transform(line string) Point {
	parts := strings.Split(line, ",")
	return Point{x: mustParseInt(parts[0]), y: mustParseInt(parts[1])}
}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	points := helpers.MustParseTo(helpers.InputFile(), transform)
	compressedYCoordinates := make(map[int]int)
	compressedXCoordinates := make(map[int]int)
	sort.Slice(points, func(i, j int) bool {
		return points[i].y < points[j].y
	})
	xPoints := make([]Point, len(points))
	copy(xPoints, points)
	// reduce xPoints, remove dupes

	sort.Slice(xPoints, func(i, j int) bool {
		return xPoints[i].x < xPoints[j].x
	})
	currCord := 0
	for _, point := range xPoints {
		if _, ok := compressedXCoordinates[point.x]; ok {
			continue	
		}
		compressedXCoordinates[point.x] = currCord
		currCord++
	}
	yPoints := make([]Point, len(points))
	copy(yPoints, points)
	sort.Slice(yPoints, func(i, j int) bool {
		return yPoints[i].y < yPoints[j].y
	})
	currCord = 0 
	for _, point := range yPoints {
		if _, ok := compressedYCoordinates[point.y]; ok {
			continue
		}
		compressedYCoordinates[point.y] = currCord
		currCord++
	}

	// and now, we can proceed with dumbass brute force
	board := make([][]string, len(compressedYCoordinates))
	for i := range board {
		board[i] = make([]string, len(compressedXCoordinates))
	}
	for _, point := range points {
		board[compressedYCoordinates[point.y]][compressedXCoordinates[point.x]] = "#"
	}
	for _, row := range board {
		fmt.Println(row)
	}







}