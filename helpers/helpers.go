package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadLines reads a file and returns a slice of strings (one per line)
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// MustReadLines reads a file and returns lines, panicking on error
func MustReadLines(filename string) []string {
	lines, err := ReadLines(filename)
	if err != nil {
		panic(fmt.Sprintf("failed to read file %s: %v", filename, err))
	}
	return lines
}

// ParseTo reads a file and parses each line using the provided parser function
func ParseTo[T any](filename string, parser func(string) T) ([]T, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	result := make([]T, len(lines))
	for i, line := range lines {
		result[i] = parser(line)
	}

	return result, nil
}

// MustParseTo reads a file and parses each line, panicking on error
func MustParseTo[T any](filename string, parser func(string) T) []T {
	result, err := ParseTo(filename, parser)
	if err != nil {
		panic(fmt.Sprintf("failed to parse file %s: %v", filename, err))
	}
	return result
}

// InputFile returns the input filename based on command-line arguments
// Returns "input.txt" if any of these flags are present: -i, --input, -input, input
// Otherwise returns "sample.txt"
func InputFile() string {
	for _, arg := range os.Args[1:] {
		arg = strings.ToLower(strings.TrimSpace(arg))
		if arg == "-i" || arg == "--input" || arg == "-input" || arg == "input" {
			return "input.txt"
		}
	}
	return "sample.txt"
}
