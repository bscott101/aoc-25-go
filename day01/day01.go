package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return ""
	}
	return string(data)
}

type SafeInput struct {
	direction string
	turns     int
}

func formatString(data string) []SafeInput {
	lines := strings.Split(data, "\n")
	inputs := []SafeInput{}

	for _, line := range lines[:len(lines)-1] {
		direction := line[0]
		turns, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Error converting string line to int: %s\n", err)
			return []SafeInput{}
		}
		input := SafeInput{
			direction: string(direction),
			turns:     turns,
		}
		inputs = append(inputs, input)
	}
	return inputs
}

func partOneSolution(inputs []SafeInput) int {
	position := 50
	hitCounter := 0
	for _, input := range inputs {
		if input.direction == "L" {
			position -= input.turns
		} else {
			position += input.turns
		}
		// ok turns out turns can be over 100 - rip
		position = position % 100
		switch {
		case position < 0:
			position = 100 + position
		case position > 99:
			position = position - 100
		}
		if position == 0 {
			hitCounter += 1
		}
	}

	return hitCounter
}

func main() {
	filePath := "./day01/input.txt"
	file := ReadFile(filePath)
	data := formatString(file)

	partOneRes := partOneSolution(data)
	fmt.Printf("Part One: %d\n", partOneRes)
}
