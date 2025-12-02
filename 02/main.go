package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2025/helpers"
)

func transform(line string) [][]int {
	ranges := [][]int{}
	for _, rangeString := range strings.Split(line, ",") {
		parts := strings.Split(rangeString, "-")
		if len(parts) != 2 {
			panic(fmt.Sprintf("invalid range string: %s", rangeString))
		}
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, []int{min, max})
	}
	return ranges
}

func main() {
	rangesArr := helpers.MustParseTo(helpers.InputFile(), transform)
	ranges := rangesArr[0]
	invalids := []int{}
	for _, rng := range ranges {
		for i := rng[0]; i <= rng[1]; i++ {
			if isRepeatPart1(i) {
				fmt.Println(fmt.Sprintf("found invalid %d for range %+v", i, rng))
				invalids = append(invalids, i)
			}
		}
		fmt.Println(fmt.Sprintf("found %d invalids for range %+v", len(invalids), rng))
	}
	sum := 0
	for _, invalid := range invalids {
		sum += invalid
	}
	fmt.Println(sum)
}

func isRepeat(num int) bool {
	str := strconv.Itoa(num)
	for i := 1; i <= len(str) / 2; i++ {
		substr := str[:i]
		// rule out any case where the string % i != 0
		if len(str) % i != 0 {
			continue
		}
		// check if the substr is repeated in the string
		if strings.Repeat(substr, len(str)/i) == str {
			return true
		}
	}
	return false

}

func isRepeatPart1(num int) bool {
	// TODO(mparker): can we get this to be a pure math operation?
	str := strconv.Itoa(num)
	if len(str) % 2 != 0 {
		return false
	}
	return !(str[:len(str)/2] != str[len(str)/2:])
}
