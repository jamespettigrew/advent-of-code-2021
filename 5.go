package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type point struct {
	x int
	y int
}

type lineSegment struct {
	p1 point
	p2 point
}

func cmp(p1 point, p2 point) int {
	if p1.x < p2.x {
		return -1
	}
	if p2.x < p1.x {
		return 1
	}
	if p1.y < p2.y {
		return -1
	}
	if p2.y < p1.y {
		return 1
	}

	return 0
}

func stringSliceToIntSlice(s []string) []int {
	result := make([]int, 0)
	for _, x := range s {
		if x == "" {
			continue
		}
		val, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println("Err stringSLiceToIntSlice", err)
			continue
		}
		result = append(result, val)
	}

	return result
}

func parseSegments(filePath string) ([]lineSegment, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	segments := make([]lineSegment, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			return nil, err
		}

		points := strings.Split(scanner.Text(), " -> ")
		coords := stringSliceToIntSlice(strings.Split(points[0], ","))
		p1 := point { x: coords[0], y: coords[1] }
		coords = stringSliceToIntSlice(strings.Split(points[1], ","))
		p2 := point { x: coords[0], y: coords[1] }

		order := cmp(p1, p2)
		if order == 1 {
			p1, p2 = p2, p1
		}
		// fmt.Printf("%d,%d -> %d,%d \n", p1.x, p1.y, p2.x, p2.y)

		segments = append(segments, lineSegment { p1: p1, p2: p2 })
	}

	return segments, scanner.Err()
}

func p5_1(segments []lineSegment) int {
	pointDensities := make(map[point]int)
	for _, segment := range segments {
		// Ignore diagonal lines for Part 1
		p1, p2 := segment.p1, segment.p2
		if p1.x != p2.x && p1.y != p2.y {
			continue
		}
		for x := p1.x; x <= p2.x; x++ {
			for y := p1.y; y <= p2.y; y++ {
				p := point { x: x, y: y }
				pointDensities[p] += 1
			}
		}
	}

	var densityThresholdCount int
	for _, v := range pointDensities {
		if v >= 2 {
			// fmt.Printf("x: %d | y:%d | count: %d\n", p.x, p.y, v)
			densityThresholdCount += 1
		}
	}

	return densityThresholdCount
}

func p5_2(segments []lineSegment) int {
	pointDensities := make(map[point]int)
	for _, segment := range segments {
		p1, p2 := segment.p1, segment.p2
		if p1.x != p2.x && p1.y != p2.y {
			// In case of diagonals, need to walk only diagonal and not entire box demarcated by points
			magnitude := p2.x - p1.x
			up := p1.y < p2.y
			// fmt.Printf("%d, %d -> %d, %d \n", p1.x, p1.y, p2.x, p2.y)
			for i := 0; i <= magnitude; i++ {
				var p point
				if up {
					p = point { x: p1.x + i, y: p1.y + i }
				} else {
					p = point { x: p1.x + i, y: p1.y - i }
				}
				// fmt.Printf("%d, %d \n", p.x, p.y)
				pointDensities[p] += 1
			}
			// fmt.Println("")
		} else {
			for x := p1.x; x <= p2.x; x++ {
				for y := p1.y; y <= p2.y; y++ {
					p := point { x: x, y: y }
					pointDensities[p] += 1
				}
			}
		}
	}

	var densityThresholdCount int
	for _, v := range pointDensities {
		if v >= 2 {
			// fmt.Printf("x: %d | y:%d | count: %d\n", p.x, p.y, v)
			densityThresholdCount += 1
		}
	}

	return densityThresholdCount
}

func main() {
	segments, err := parseSegments("5.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}
	part1 := p5_1(segments)
	part2 := p5_2(segments)
	fmt.Printf("Part 1: %d | Part 2: %d", part1, part2)
}
