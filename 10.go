package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack {
		data: make([]string, 0),
	}
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Top() (*string, error) {
	if s.Empty() {
		return nil, errors.New("empty")
	}

	return &s.data[len(s.data) - 1], nil
}

func (s *Stack) Push(frame string) {
	s.data = append(s.data, frame)
}

func (s *Stack) Pop() (*string, error) {
	if len(s.data) == 0 {
		return nil, errors.New("empty")
	}
	ret := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]

	return &ret, nil
}

func inverseChar(char string) string {
	switch char {
	case "(": return ")"
	case "[": return "]"
	case "{": return "}"
	case "<": return ">"
	default: return ""
	}
}

func illegalCharScore(char string) int {
	switch char {
	case ")": return 3
	case "]": return 57
	case "}": return 1197
	case ">": return 25137
	default: return 0
	}
}

func completionScore(char string) int {
	switch char {
	case ")": return 1
	case "]": return 2
	case "}": return 3
	case ">": return 4
	default: return 0
	}
}

func p1(lines [][]string) int {
	score := 0
	for _, line := range lines {
		stack := NewStack()
		outer: for _, char := range line {
			switch char {
			case "(", "[", "{", "<":
				stack.Push(char)
			case ")", "]", "}", ">":
				popped, err := stack.Pop()
				if err != nil || char != inverseChar(*popped) {
					score += illegalCharScore(char)
					break outer
				}
			}
		}
	}

	return score
}

func p2(lines [][]string) int {
	scores := make([]int, 0)
	for _, line := range lines {
		stack := NewStack()
		corrupted := false
		outer: for _, char := range line {
			switch char {
			case "(", "[", "{", "<":
				stack.Push(char)
			case ")", "]", "}", ">":
				popped, err := stack.Pop()
				if err != nil || char != inverseChar(*popped) {
					corrupted = true
					break outer
				}
			}
		}

		if corrupted {
			continue
		}

		score := 0
		for !stack.Empty() {
			score *= 5
			char, _ := stack.Pop()
			inverse := inverseChar(*char)
			score += completionScore(inverse)

		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	fmt.Println(scores)
	return scores[len(scores) / 2]
}

func parseInput(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")

		lines = append(lines, chars)
		if err != nil {
			return lines, err
		}
	}
	return lines, nil
}

func main() {
	input, err := parseInput("10.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}
	part1 := p1(input)
	part2 := p2(input)
	fmt.Printf("Part 1: %d | Part 2: %d", part1, part2)
}
