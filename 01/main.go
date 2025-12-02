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

func (s *Safe) move(instruction Instruction) int {
	zeros := 0
	switch instruction.direction {
	case "R":
		for i := 0; i < instruction.steps; i++ {
			s.currentPosition++
			if s.currentPosition == s.capacity {
				// we're passing zero. whether or not we term on this instruction, we're recording either that we've passed zero or we're stopping at zero
				s.currentPosition = 0
				zeros++
			}
		}
	case "L":
		for i := 0; i < instruction.steps; i++ {
			s.currentPosition--
			if s.currentPosition == 0 {
				zeros++
			} else if s.currentPosition <= 0 {
				s.currentPosition = s.capacity - 1
			}
		}
	}
	return zeros
}

func main() {
	instructions := helpers.MustParseTo(helpers.InputFile(), transform)

	safe := Safe{
		currentPosition: 50,
		capacity: 100,
	}
	atZero := 0
	for _, instruction := range instructions {
		zeros := safe.move(instruction)
		atZero += zeros
		fmt.Println(fmt.Sprintf("passed zero %d times executing instruction %+v", zeros, instruction))
	}
	fmt.Println(atZero)
}
