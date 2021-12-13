package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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

func p5_1(lanternFish []int) int {
	temp := make([]int, len(lanternFish))
	copy(temp, lanternFish)
	lanternFish = temp

	var count int
	for day := 0; day < 80; day++ {
		count = len(lanternFish)
		for f := 0; f < count; f++ {
			lanternFish[f] -= 1
			if lanternFish[f] == -1 {
				lanternFish[f] = 6
				lanternFish = append(lanternFish, 8)
			}
		}
	}

	return len(lanternFish)
}
// 		   0  1  2  3  4  5  6  7  8
// Day 0: [0, 1, 1, 2, 1, 0, 0, 0, 0]
// Day 1: [1, 1, 2, 1, 0, 0, 0, 0, 0]
// Day 2: [1, 2, 1, 0, 0, 0, 1, 0, 1]
func p5_2(lanternFish []int) int64 {
	counts := [9]int { 0, 0, 0, 0, 0, 0, 0, 0, 0 }

	// Seed
	for _, f := range lanternFish {
		counts[f] += 1
	}

	for day := 0; day < 256; day++ {
		// fmt.Printf("DAY %d | 0: %d| 1: %d| 2: %d| 3: %d| 4: %d| 5: %d| 6: %d| 7: %d| 8: %d \n", day, counts[0], counts[1], counts[2], counts[3], counts[4], counts[5], counts[6], counts[7], counts[8])

		var newCounts [9]int
		for i := 1; i < 9; i++ {
			newCounts[i - 1] = counts[i]
		}
		newCounts[8] = counts[0]
		newCounts[6] += counts[0]
		counts = newCounts
	}

	var total int64
	for _, count := range counts {
		total += int64(count)
	}

	return total
}

func main() {
	input, err := parseInput("6.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}
	part1 := p5_1(input)
	part2 := p5_2(input)
	fmt.Printf("Part 1: %d | Part 2: %d", part1, part2)
}
