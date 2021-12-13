package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Stack struct {
	data []StackFrame
}

func NewStack() *Stack {
	return &Stack {
		data: make([]StackFrame, 0),
	}
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(frame StackFrame) {
	s.data = append(s.data, frame)
}

func (s *Stack) Pop() (*StackFrame, error) {
	if len(s.data) == 0 {
		return nil, errors.New("empty")
	}
	ret := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]

	return &ret, nil
}

type StackFrame struct {
	current point
	prev point
}


func parseInput(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	grid := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rawColumns := strings.Split(line, "")
		columns := make([]int, 0)
		for _, c := range rawColumns {
			num, _ := strconv.Atoi(c)
			columns = append(columns, num)
		}

		grid = append(grid, columns)
		if err != nil {
			return grid, err
		}
	}
	return grid, nil
}

func all(collection []int, predicate func(x int) bool) bool {
	for _, i := range collection {
		if predicate(i) == false {
			return false
		}
	}

	return true
}

type point struct {
	row int
	column int
}

func getLowPoints(grid [][]int) []point {
	lowPoints := make([]point, 0)
	height := len(grid)
	width := len(grid[0])
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			adjacents := make([]int, 0)
			if r > 0 {
				adjacents = append(adjacents, grid[r - 1][c])
			}
			if r < height - 1 {
				adjacents = append(adjacents, grid[r + 1][c])
			}
			if c > 0 {
				adjacents = append(adjacents, grid[r][c - 1])
			}
			if c < width - 1 {
				adjacents = append(adjacents, grid[r][c + 1])
			}

			if all(adjacents, func(x int) bool { return grid[r][c] < x }) {
				lowPoints = append(lowPoints, point { row: r, column: c })
			}
		}
	}

	return lowPoints
}


func getBasinRec(grid [][]int, current point, prev point, basinPoints map[point]bool) {
	height := len(grid)
	width := len(grid[0])

	r, c := current.row, current.column
	if grid[r][c] == 9 {
		return
	}

	if grid[r][c] < grid[prev.row][prev.column] {
		// Haven't flowed upwards from lowpoint
		return
	}

	if r > 0 {
		getBasinRec(grid, point { row: r - 1, column: c }, current, basinPoints)
	}
	if r < height - 1 {
		getBasinRec(grid, point { row: r + 1, column: c }, current, basinPoints)
	}
	if c > 0 {
		getBasinRec(grid, point { row: r, column: c - 1 }, current, basinPoints)
	}
	if c < width - 1 {
		getBasinRec(grid, point { row: r, column: c + 1 }, current, basinPoints)
	}

	basinPoints[current] = true
}

func getBasinStack(grid [][]int, lowPoint point) []point {
	height := len(grid)
	width := len(grid[0])

	basinPoints := make(map[point]bool)
	stack := NewStack()
	stack.Push(StackFrame { current: lowPoint, prev: lowPoint })
	for !stack.Empty() {
		frame, _ := stack.Pop()
		current, prev := frame.current, frame.prev
		r, c := current.row, current.column
		if _, ok := basinPoints[current]; ok {
			continue
		}
		if grid[r][c] == 9 {
			continue
		}

		if grid[r][c] < grid[prev.row][prev.column] {
			// Haven't flowed upwards from lowpoint
			continue
		}

		if r > 0 {
			stack.Push(StackFrame { current: point { row: r - 1, column: c }, prev: current })
		}
		if r < height - 1 {
			stack.Push(StackFrame { current: point { row: r + 1, column: c }, prev: current })
		}
		if c > 0 {
			stack.Push(StackFrame { current: point { row: r, column: c - 1 }, prev: current })
		}
		if c < width - 1 {
			stack.Push(StackFrame { current: point { row: r, column: c + 1 }, prev: current })
		}

		basinPoints[current] = true
	}

	ret := make([]point, 0)
	for p, _ := range basinPoints {
		ret = append(ret, p)
	}

	return ret
}

func getBasin(grid [][]int, lowPoint point) []point {
	basinPoints := make(map[point]bool)
	getBasinRec(grid, lowPoint, lowPoint, basinPoints)

	ret := make([]point, 0)
	for p, _ := range basinPoints {
		ret = append(ret, p)
	}

	return ret
}

func p1(grid [][]int) int {
	sumRiskLevel := 0
	lowPoints := getLowPoints(grid)
	for _, p := range lowPoints {
		sumRiskLevel += grid[p.row][p.column] + 1
	}
	return sumRiskLevel
}

func p2(grid [][]int) int {
	lowPoints := getLowPoints(grid)
	basinSizes := make([]int, 0)
	for _, p := range lowPoints {
		basin := getBasinStack(grid, p)
		basinSizes = append(basinSizes, len(basin))
	}
	sort.Slice(basinSizes, func(i, j int) bool {  return basinSizes[i] > basinSizes[j] })

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func main() {
	input, err := parseInput("9.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}
	part1 := p1(input)
	part2 := p2(input)
	fmt.Printf("Part 1: %d | Part 2: %d", part1, part2)
}
