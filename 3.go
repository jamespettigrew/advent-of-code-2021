package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func kthBitIsSet(number int, k int) bool {
	return number & (1 << k) != 0
}

func bitValueCounts(numbers []int, bit int) (low int, high int) {
	var highCount int
	for _, number := range numbers {
		if kthBitIsSet(number, bit) {
			highCount += 1
		}
	}

	return len(numbers) - highCount, highCount
}

func gamma(numbers []int, width int) int {
	var result int
	for k := 0; k < width; k++ {
		low, high := bitValueCounts(numbers, width - k - 1)
		if high > low {
			result |= (1 << (width - k - 1))
		}
	}

	return result
}

func oxygenGeneratorRating(numbers []int, width int) int {
	remaining := append([]int{}, numbers...)
	bit := width - 1
	for len(remaining) > 1 {
		low, high := bitValueCounts(remaining, bit)

		var keep []int
		for _, number := range remaining {
			if high >= low && kthBitIsSet(number, bit) {
				keep = append(keep, number)
			} else if low > high && !kthBitIsSet(number, bit) {
				keep = append(keep, number)
			}
		}

		bit -= 1
		remaining = keep
	}

	return remaining[0]
}

func co2GeneratorRating(numbers []int, width int) int {
	remaining := append([]int{}, numbers...)
	bit := width - 1
	for len(remaining) > 1 {
		low, high := bitValueCounts(remaining, bit)

		var keep []int
		for _, number := range remaining {
			if low <= high && !kthBitIsSet(number, bit) {
				keep = append(keep, number)
			} else if high < low && kthBitIsSet(number, bit) {
				keep = append(keep, number)
			}
		}
		bit -= 1
		remaining = keep
	}

	return remaining[0]
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
		x, err := strconv.ParseInt(scanner.Text(), 2, 0)
		if err != nil {
			return nil, err
		}

		sequence = append(sequence, int(x))
	}

	return sequence, scanner.Err()
}

func main() {
	sequence, err := readInputSequence("3.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}
	width := 12

	gamma := gamma(sequence, width)
	epsilon := gamma ^ ((1 << width) - 1)
	fmt.Println(gamma * epsilon)

	ogr := oxygenGeneratorRating(sequence, width)
	cgr := co2GeneratorRating(sequence, width)
	fmt.Println(ogr * cgr)
}
