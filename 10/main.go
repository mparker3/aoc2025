package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc2025/helpers"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize/convex/lp"
)

type Engine struct {
	buttonCombos [][]int
	desiredState []int
}

func transform(line string) Engine {
	parts := strings.Split(line, " ")
	rawButtonState := parts[len(parts)-1]
	rawButtonState = strings.ReplaceAll(rawButtonState, "{", "")
	rawButtonState = strings.ReplaceAll(rawButtonState, "}", "")
	var desiredState []int
	for _, button := range strings.Split(rawButtonState, ",") {
		desiredState = append(desiredState, mustParseInt(button))
	}
	buttonCombos := [][]int{}
	for _, rawCombo := range parts[1 : len(parts)-1] {
		rawCombo = strings.ReplaceAll(rawCombo, "(", "")
		rawCombo = strings.ReplaceAll(rawCombo, ")", "")
		buttonCombo := []int{}
		for _, button := range strings.Split(rawCombo, ",") {
			buttonInt := mustParseInt(button)
			buttonCombo = append(buttonCombo, buttonInt)
		}
		buttonCombos = append(buttonCombos, buttonCombo)
	}

	return Engine{desiredState: desiredState, buttonCombos: buttonCombos}

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
	for _, engine := range engines {
		rhs := mat.NewDense(len(engine.buttonCombos), len(engine.desiredState), nil)
		for i, combo := range engine.buttonCombos {
			for _, button := range combo {
				rhs.Set(i, button, 1)
			}
		}
		lhs := mat.NewDense(len(engine.desiredState), 1, nil)
		for i, state := range engine.desiredState {
			lhs.Set(i, 0, float64(state))
		}

		m := len(engine.buttonCombos)   // number of combos (variables)
		n := len(engine.desiredState)   // state dimension (constraints)

		// Aeq = rhsᵀ (n x m)
		var Aeq mat.Dense
		Aeq.CloneFrom(rhs.T())

		// beq = lhs flattened
		beq := make([]float64, n)
		for i := 0; i < n; i++ {
			beq[i] = lhs.At(i, 0)
		}

		// Objective: minimize sum(x_i) or just find any feasible solution.
		c := make([]float64, m)
		for i := range c {
			c[i] = 1 // or 0 if you don't care about minimizing
		}

		// Convert general-form LP (no inequalities, only equalities) to standard form
		cNew, ANew, bNew := lp.Convert(
			c,
			nil, nil,   // G, h : no inequality constraints
			&Aeq, beq,  // Aeq, beq : our equalities rhsᵀ x = lhs
		)

		// Now solve the standard-form LP
		opt, x, err := lp.Simplex(cNew, ANew, bNew, 0, nil)
		if err != nil {
			log.Fatalf("Simplex failed: %v", err)
		}

		fmt.Printf("opt: %v\n", opt)
		fmt.Printf("x (combo coefficients): %v\n", x[:m]) // first m entries correspond to original variables

		// Optional: verify rhsᵀ * x ≈ lhs
		xVec := mat.NewVecDense(m, x[:m])
		var rhsT mat.Dense
		rhsT.CloneFrom(rhs.T())

		var recon mat.VecDense
		recon.MulVec(&rhsT, xVec)

		fmt.Println("reconstructed lhs from rhsᵀ * x:")
		fmt.Println(mat.Formatted(&recon))
	}
}
