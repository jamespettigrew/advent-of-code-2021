package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countIncreases(sequence []int) int {
	var prev, increases int
	for _, val := range sequence[1:] {
		if val > prev {
			increases += 1
		}
		prev = val
	}

	return increases
}

func sumSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}

	return sum
}

func countIncreasesSlidingWindow(sequence []int, windowSize int) int {
	var increases, prev int
	for i := windowSize; i < len(sequence); i++ {
		low := i - windowSize
		window := sequence[low: i]
		sum := sumSlice(window)
		if sum > prev {
			increases += 1
		}
		prev = sum
	}

	return increases
}

func readInputSequence(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sequence []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return sequence, err
		}
		sequence = append(sequence, val)
	}

	return sequence, scanner.Err()
}

func main() {
	sequence, err := readInputSequence("1.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}
	result := countIncreases(sequence)
	fmt.Println(result)
	windowedResult := countIncreasesSlidingWindow(sequence, 3)
	fmt.Println(windowedResult)

}
