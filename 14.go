package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type state struct {
	template string
	pairs    map[string]string
}

func step(polymer string, pairs map[string]string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		buffer.WriteByte(pair[0])
		buffer.WriteString(pairs[pair])
	}
	buffer.WriteByte(polymer[len(polymer)-1])

	return buffer.String()
}

func countElements(polymer string) map[string]int {
	counts := make(map[string]int)
	for _, b := range polymer {
		counts[string(b)] += 1
	}

	return counts
}

func faster(polymer string, pairs map[string]string, steps int) map[string]int {
	pairCounts := make(map[string]int)
	elementCounts := make(map[string]int)
	// Get initial counts
	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		pairCounts[string(pair)] += 1
		elementCounts[string(pair[0])] += 1
	}
	elementCounts[string(polymer[len(polymer)-1])] += 1

	for i := 0; i < steps; i++ {
		newCounts := make(map[string]int)
		for k, v := range pairCounts {
			newPair := string(k[0]) + pairs[k]
			newCounts[newPair] += v
			newPair = pairs[k] + string(k[1])
			elementCounts[pairs[k]] += v
			newCounts[newPair] += v
		}
		pairCounts = newCounts
	}

	return elementCounts
}

func parseInput(filePath string) (state, error) {
	s := state{pairs: make(map[string]string)}

	file, err := os.Open(filePath)
	if err != nil {
		return s, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s.template = scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		split := strings.Split(line, " -> ")
		s.pairs[split[0]] = split[1]
		if err != nil {
			return s, err
		}
	}
	return s, nil
}

func main() {
	state, err := parseInput("14.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}

	// polymer := state.template
	// for i := 0; i < 10; i++ {
	// 	polymer = step(polymer, state.pairs)
	// }
	// counts := countElements(polymer)
	//

	counts := faster(state.template, state.pairs, 40)
	least := 0
	most := 0
	for _, v := range counts {
		if least == 0 || v < least {
			least = v
		}
		if most == 0 || v > most {
			most = v
		}
	}

	fmt.Printf("Result: %d \n", most-least)
}
