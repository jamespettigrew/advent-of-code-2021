package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type stackFrame struct {
	vertex string
	path []string
	visited map[string]int
	doubleVisited bool
}

type stack struct {
	data []stackFrame
}

func NewStack() *stack {
	return &stack {
		data: make([]stackFrame, 0),
	}
}

func (s *stack) Empty() bool {
	return len(s.data) == 0
}

func (s *stack) Top() (*stackFrame, error) {
	if s.Empty() {
		return nil, errors.New("empty")
	}

	return &s.data[len(s.data) - 1], nil
}

func (s *stack) Push(frame stackFrame) {
	s.data = append(s.data, frame)
}

func (s *stack) Pop() (*stackFrame, error) {
	if len(s.data) == 0 {
		return nil, errors.New("empty")
	}
	ret := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]

	return &ret, nil
}

type Edges [][]string

func parseInput(filePath string) (Edges, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	edges := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		vertices := strings.Split(line, "-")
		edges = append(edges, vertices)
		if err != nil {
			return edges, err
		}
	}
	return edges, nil
}

func countPaths(e Edges) int {
	graph := make(map[string][]string)
	for _, edge := range e {
		v1, v2 := edge[0], edge[1]
		if connections, ok := graph[v1]; ok {
			graph[v1] = append(connections, v2)
		} else {
			graph[v1] = []string { v2 }
		}

		if connections, ok := graph[v2]; ok {
			graph[v2] = append(connections, v1)
		} else {
			graph[v2] = []string { v1 }
		}
	}

	pathCount := 0
	stack := NewStack()
	stack.Push(stackFrame { vertex: "start", path: make([]string, 0), visited: make(map[string]int)})
	for !stack.Empty() {
		frame, _ := stack.Pop()

		if frame.vertex == "end" {
			// fmt.Println(append(frame.path, frame.vertex))
			pathCount += 1
			continue
		}

		// If it's a small cave
		if frame.vertex == strings.ToLower(frame.vertex) {
			if _, ok := frame.visited[frame.vertex]; ok {
				continue
			}
			frame.visited[frame.vertex]++
		}

		neighbours := graph[frame.vertex]
		for _, neighbour := range neighbours {
			if neighbour == "start" {
				continue
			}
			newVisited := make(map[string]int)
			for k,v := range frame.visited {
				newVisited[k] = v
			}
			stack.Push(stackFrame { vertex: neighbour, path: append(frame.path, frame.vertex), visited: newVisited })
		}
	}

	return pathCount
}

func countPaths2(e Edges) int {
	graph := make(map[string][]string)
	for _, edge := range e {
		v1, v2 := edge[0], edge[1]
		if connections, ok := graph[v1]; ok {
			graph[v1] = append(connections, v2)
		} else {
			graph[v1] = []string { v2 }
		}

		if connections, ok := graph[v2]; ok {
			graph[v2] = append(connections, v1)
		} else {
			graph[v2] = []string { v1 }
		}
	}

	pathCount := 0
	stack := NewStack()
	stack.Push(stackFrame { vertex: "start", path: make([]string, 0), visited: make(map[string]int)})
	for !stack.Empty() {
		frame, _ := stack.Pop()

		if frame.vertex == "end" {
			// fmt.Println(append(frame.path, frame.vertex))
			pathCount += 1
			continue
		}

		// If it's a small cave
		if frame.vertex == strings.ToLower(frame.vertex) {
			if _, ok := frame.visited[frame.vertex]; ok {
				if frame.doubleVisited {
					continue
				}
			}
			frame.visited[frame.vertex]++
		}

		neighbours := graph[frame.vertex]
		for _, neighbour := range neighbours {
			if neighbour == "start" {
				continue
			}
			newVisited := make(map[string]int)
			doubleVisited := false
			for k,v := range frame.visited {
				newVisited[k] = v
				if v > 1 {
					doubleVisited = true
				}
			}
			stack.Push(stackFrame { vertex: neighbour, path: append(frame.path, frame.vertex), visited: newVisited, doubleVisited: doubleVisited })
		}
	}

	return pathCount
}

func main() {
	edges, err := parseInput("12.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}

	part1 := countPaths(edges)
	part2 := countPaths2(edges)
	fmt.Printf("Part 1: %d | Part 2: %d", part1, part2)
}
