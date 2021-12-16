package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type axis string

const (
	x axis = "x"
	y      = "y"
)

type point struct {
	x int
	y int
}

type fold struct {
	axis     axis
	position int
}

type state struct {
	dots  []point
	folds []fold
}

func printDots(dots []point) {
	maxX, maxY := 0, 0
	dotLookup := make(map[point]bool)
	for _, dot := range dots {
		if maxX == 0 || dot.x > maxX {
			maxX = dot.x
		}
		if maxY == 0 || dot.y > maxY {
			maxY = dot.y
		}
		dotLookup[dot] = true
	}

	height, width := maxY, maxX
	for r := 0; r <= height; r++ {
		for c := 0; c <= width; c++ {
			if _, ok := dotLookup[point{x: c, y: r}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
}

func performFold(dots []point, f fold) []point {
	unique := make(map[point]bool)
	for _, dot := range dots {
		if f.axis == y && dot.y > f.position {
			dot.y = f.position - (dot.y - f.position)
		} else if f.axis == x && dot.x > f.position {
			dot.x = f.position - (dot.x - f.position)
		}
		unique[dot] = true
	}

	ret := make([]point, 0)
	for k, _ := range unique {
		ret = append(ret, k)
	}

	return ret
}

func parseInput(filePath string) (state, error) {
	s := state{dots: make([]point, 0), folds: make([]fold, 0)}

	file, err := os.Open(filePath)
	if err != nil {
		return s, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		dot := strings.Split(line, ",")
		x, _ := strconv.Atoi(dot[0])
		y, _ := strconv.Atoi(dot[1])
		s.dots = append(s.dots, point{x: x, y: y})
		if err != nil {
			return s, err
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		split = strings.Split(split[2], "=")
		axis := axis(split[0])
		position, _ := strconv.Atoi(split[1])
		s.folds = append(s.folds, fold{axis: axis, position: position})
		if err != nil {
			return s, err
		}

	}
	return s, nil
}

func main() {
	state, err := parseInput("13.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}

	part1 := performFold(state.dots, state.folds[0])
	for _, fold := range state.folds {
		state.dots = performFold(state.dots, fold)
	}

	fmt.Printf("Part 1: %d \n", len(part1))
	printDots(state.dots)
}
