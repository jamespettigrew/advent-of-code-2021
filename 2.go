package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func computePositionAndDepth(sequence []string) (int, int) {
	var position, depth int
	for _, x := range sequence {
		split := strings.Split(x, " ")
		command := split[0]
		value, _ := strconv.Atoi(split[1])
		switch command {
		case "forward":
			position += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	return position, depth
}

func computeAimedPositionAndDepth(sequence []string) (int, int) {
	var position, depth, aim int
	for _, x := range sequence {
		split := strings.Split(x, " ")
		command := split[0]
		value, _ := strconv.Atoi(split[1])
		switch command {
		case "forward":
			position += value
			depth += value * aim
			if depth < 0 {
				depth = 0
			}
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	return position, depth
}

func readInputSequence(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sequence []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sequence = append(sequence, scanner.Text())
	}

	return sequence, scanner.Err()
}

func main() {
	sequence, err := readInputSequence("2.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}
	position, depth := computePositionAndDepth(sequence)
	fmt.Println(position * depth)

	position, depth = computeAimedPositionAndDepth(sequence)
	fmt.Println(position * depth)
}
