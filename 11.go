package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid [][]int
type Point [2]int

func (g *Grid) Height() int {
	return len((*g))
}

func (g *Grid) Width() int {
	return len((*g)[0])
}

func (g *Grid) Print() {
	for r := 0; r < g.Height(); r++ {
		fmt.Println((*g)[r])
	}
	fmt.Println("")
}

func (g *Grid) Step() int {
	flashes := 0
	flashed := make(map[Point]bool)
	flashQueue := make(chan Point, 100)

	for r := 0; r < g.Height(); r++ {
		for c := 0; c < g.Width(); c++ {
			(*g)[r][c] += 1
			if (*g)[r][c] > 9 {
				(*g)[r][c] = 0
				p := Point { r, c }
				flashed[p] = true
				flashQueue <- p
				flashes += 1
			}
		}
	}

	for {
		select {
		case point := <- flashQueue:
			for _, neighbour := range g.getNeighbours(point) {
				if _, ok := flashed[neighbour]; ok {
					continue
				}

				(*g)[neighbour[0]][neighbour[1]] += 1
				if (*g)[neighbour[0]][neighbour[1]] > 9 {
					(*g)[neighbour[0]][neighbour[1]] = 0
					flashed[neighbour] = true
					flashQueue <- neighbour
					flashes += 1
				}
			}
		default:
			return flashes
		}
	}
}

func (g *Grid) getNeighbours(p Point) []Point {
	neighbours := make([]Point, 0)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			r := p[0] + i
			c := p[1] + j
			if r >= 0 && c >= 0 && r < g.Height() && c < g.Width() {
				neighbours = append(neighbours, Point { r, c })
			}
		}
	}

	return neighbours
}

func parseInput(filePath string) (Grid, error) {
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

func p1(g Grid) int {
	totalFlashes := 0
	for i:= 0; i < 100; i++ {
		totalFlashes += g.Step()
	}

	return totalFlashes
}

func p2(g Grid) int {
	step := 0
	for {
		g.Print()
		flashes := g.Step()
		step += 1
		if flashes == g.Width() * g.Height() {
			return step
		}
	}
}

func main() {
	grid, err := parseInput("11.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}
	duplicate := make([][]int, len(grid))
	for i := range grid {
		duplicate[i] = make([]int, len(grid[i]))
		copy(duplicate[i], grid[i])
	}

	part1 := p1(grid)
	part2 := p2(duplicate)
	fmt.Printf("Part 1: %d | Part 2: %d", part1, part2)
}
