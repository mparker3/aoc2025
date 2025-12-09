package main

import (
	"fmt"
	"math"
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
	maxDistance := -1.0
	for i, self := range points {
		for j, other := range points {
			if i == j {
				continue
			}
			distance := (math.Abs(float64(self.x - other.x)) + 1) * (math.Abs(float64(self.y - other.y) + 1))
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}
	distanceInt := int64(maxDistance)
	fmt.Println(distanceInt)
}