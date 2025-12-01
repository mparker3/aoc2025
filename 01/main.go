package main

import (
	"fmt"
	"strconv"

	"aoc2025/helpers"
)

type Instruction struct {
	direction string
	steps int
}

func transform(line string) Instruction {
	direction, stepsInt := string(line[0]), line[1:]
	steps, _ := strconv.Atoi(stepsInt)

	return Instruction{
		direction: direction,
		steps: steps,
	}
}

type Safe struct {
	currentPosition int
	capacity int
}

func (s *Safe) move(instruction Instruction) {
	remainder := instruction.steps % s.capacity
	switch instruction.direction {
	case "R":
		s.currentPosition += remainder
		if s.currentPosition >= s.capacity {
			s.currentPosition = s.currentPosition - s.capacity
		}
	case "L":
		s.currentPosition -= remainder
		if s.currentPosition < 0 {
			s.currentPosition = s.capacity + s.currentPosition
		}
	}
}

func main() {
	_ = helpers.MustParseTo(helpers.InputFile(), transform)

	instructions := helpers.MustParseTo(helpers.InputFile(), transform)

	safe := Safe{
		currentPosition: 50,
		capacity: 100,
	}
	atZero := 0
	for _, instruction := range instructions {
		safe.move(instruction)
		if safe.currentPosition == 0 {
			atZero++
		}
		fmt.Println(fmt.Sprintf("Moving %s %d steps, current position: %d", instruction.direction, instruction.steps, safe.currentPosition))
	}
	fmt.Println(atZero)
}
