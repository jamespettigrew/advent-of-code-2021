package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func stringSliceToIntSlice(s []string) []int {
	result := make([]int, 0)
	for _, x := range s {
		if x == "" {
			continue
		}
		val, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println("Err stringSliceToIntSlice", err)
			continue
		}
		result = append(result, val)
	}

	return result
}

func parseInput(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return stringSliceToIntSlice(strings.Split(scanner.Text(), ",")), nil
}

func bounds(input []int) (min, max int) {
	for _, x := range input {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	return min, max
}


func sumDistances(input []int, y int) int {
	sum := 0
	for _, x := range input {
		sum += int(math.Abs(float64(y - x)))
	}

	return sum
}

func gaussFormula(input []int, y int) int {
	sum := 0
	for _, x := range input {
		delta := int(math.Abs(float64(y - x)))
		sum += delta * (delta + 1) / 2
	}

	return sum
}

func p7_1(input []int) int {
	min, max := bounds(input)

	var minSum int
	for y := min; y <= max; y++ {
		sum := sumDistances(input, y)
		if minSum == 0 || sum < minSum {
			minSum = sum
		}
	}

	return minSum
}

func p7_2(input []int) int {
	min, max := bounds(input)

	var minSum int
	for y := min; y <= max; y++ {
		sum := gaussFormula(input, y)
		if minSum == 0 || sum < minSum {
			minSum = sum
		}
	}

	return minSum
}

func main() {
	input, err := parseInput("7.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}
	part1 := p7_1(input)
	part2 := p7_2(input)
	fmt.Printf("Part 1: %d | Part 2: %d", part1, part2)
}
