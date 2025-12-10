package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2025/helpers"
)

type Engine struct {
	buttonCombos []uint
	desiredState  uint
}

func transform(line string) Engine {
	parts := strings.Split(line, " ")
	rawButtonState := parts[0]
	rawButtonState = strings.ReplaceAll(rawButtonState, "[", "")
	rawButtonState = strings.ReplaceAll(rawButtonState, "]", "")
	var buttonState uint
	lenButton := len(strings.Split(rawButtonState, ""))
	for _, button := range strings.Split(rawButtonState, "") {
		if button == "." {
		} else {
			buttonState |= 1
		}
		buttonState <<= 1
	}
	buttonState >>= 1

	buttonCombos := []uint{}
	for _, rawCombo := range parts[1:len(parts)-1] {
		rawCombo = strings.ReplaceAll(rawCombo, "(", "")
		rawCombo = strings.ReplaceAll(rawCombo, ")", "")
		buttonCombo := 0

		for _, button := range strings.Split(rawCombo, ",") {
			// for each button in the combo, set the corresponding nth bit in the buttonCombo
			buttonInt := mustParseInt(button)
			// which bit should we set? 
			bit := lenButton - buttonInt - 1
			buttonCombo |= 1 << bit
		}
		buttonCombos = append(buttonCombos, uint(buttonCombo))
	}


	return Engine{desiredState: buttonState, buttonCombos: buttonCombos}

}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	engines := helpers.MustParseTo(helpers.InputFile(), transform)
	shortest := []int{}
	for _, engine := range engines {
		shortest = append(shortest, search(0, []uint{0}, engine.desiredState, engine.buttonCombos))
	}
	total := 0
	for _, length := range shortest {
		total += length
	}
	fmt.Println(total)
}

func search(depth int, states []uint, desiredState uint, buttonCombos []uint) int {
	depth++
	newStates := []uint{}
	// push each button and check if that gets us to the desired state
	for _, state := range states {
		for _, combo := range buttonCombos {
			newState := state ^ combo
			if newState == desiredState {
				return depth
			}
			newStates = append(newStates, newState)
		}
	}

	return search(depth, newStates, desiredState, buttonCombos)
}
