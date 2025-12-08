package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"aoc2025/helpers"
)

var boxID = 0

type Box struct {
	id      int
	x, y, z int
}

func (b Box) distance(other Box) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(b.x-other.x)), 2) + math.Pow(math.Abs(float64(b.y-other.y)), 2) + math.Pow(math.Abs(float64(b.z-other.z)), 2))
}

type Distance struct {
	distance float64
	self     Box
	other    Box
}


func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func transform(line string) Box {
	parts := strings.Split(line, ",")

	box := Box{
		id: boxID,
		x:  mustParseInt(parts[0]),
		y:  mustParseInt(parts[1]),
		z:  mustParseInt(parts[2]),
	}
	boxID++
	return box
}

func main() {
	boxes := helpers.MustParseTo(helpers.InputFile(), transform)

	// let's brute force to start
	boxDistances := make([]Distance, 0)
	for i, box := range boxes {
		for j, other := range boxes {
			if i != j && i > j { // don't compute a distance twice
				boxDistances = append(boxDistances, Distance{distance: box.distance(other), self: box, other: other})
			}
		}
	}

	sort.Slice(boxDistances, func(i, j int) bool {
		return boxDistances[i].distance < boxDistances[j].distance
	})

	boxesCopy := make([]Box, len(boxes))
	copy(boxesCopy, boxes)

	circuits := make([]*Circuit, len(boxes))
	for i := range boxes {
		circuits[i] = &Circuit{id: i, boxes: []Box{boxes[i]}}
	}

	// already sorted
	for _, distance := range boxDistances {
		circuit1 := circuits[distance.self.id]
		circuit2 := circuits[distance.other.id]
		newCircuit := circuit1.merge(circuit2)
		fmt.Println("merged circuits", circuit1.id, circuit2.id, "->", newCircuit.id)
		fmt.Println("new circuit size", len(newCircuit.boxes))
		// this w
		for _, box := range newCircuit.boxes {
			circuits[box.id] = newCircuit
		}
		if len(newCircuit.boxes) == len(boxes) {
			fmt.Println(distance.self.x * distance.other.x)
			return 
		}
	}

	// get all the unique circuits
	seen := make(map[int]bool)
	uniqueCircuits := make([]*Circuit, 0)
	for _, circuit := range circuits {
		if !seen[circuit.id] {
			seen[circuit.id] = true
			uniqueCircuits = append(uniqueCircuits, circuit)
		}
	}

	// sort the circuits by size descending
	sort.Slice(uniqueCircuits, func(i, j int) bool {
		return len(uniqueCircuits[i].boxes) > len(uniqueCircuits[j].boxes)
	})

	topCircuits := uniqueCircuits[:3]
	product := 1
	for _, circuit := range topCircuits {
		product *= len(circuit.boxes)
	}
	fmt.Println(product)


}

func (c *Circuit) merge(other *Circuit) (*Circuit) {
	if c.id == other.id {
		// we don't need to do anything or update anything here
		return c
	}

	// these are two different circuits
	// create a new circuit with the boxes from both circuits
	newCircuit := &Circuit{id: c.id, boxes: append(c.boxes, other.boxes...)}
	return newCircuit

}


type Circuit struct{
	id int
	boxes []Box
}