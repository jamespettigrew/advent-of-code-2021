package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type game struct {
	draw []int
	boards [][][]int
}

func printBoard(board [][]int) {
	fmt.Println("-- BOARD --")
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			number := board[r][c]
			fmt.Printf("%d ", number)
		}
		fmt.Print("\n")
	}
	fmt.Println("-- END --")
}

func boardComplete(board [][]int, drawn map[int]bool) bool {
	for r := 0; r < 5; r++ {
		prevDrawn := true
		for c := 0; c < 5; c++ {
			number := board[r][c]
			if _, ok := drawn[number]; !ok {
				prevDrawn = false
				break
			}
		}

		if prevDrawn == true {
			return true
		}
	}

	for c := 0; c < 5; c++ {
		prevDrawn := true
		for r := 0; r < 5; r++ {
			number := board[r][c]
			if _, ok := drawn[number]; !ok {
				prevDrawn = false
				break
			}
		}

		if prevDrawn == true {
			return true
		}
	}

	return false
}

func sumUnmarked(board [][]int, drawn map[int]bool) int {
	var sum int
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			number := board[r][c]
			if _, ok := drawn[number]; !ok {
				sum += number
			}
		}
	}

	return sum
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

func parseInput(filePath string) (*game, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	movesInput := strings.Split(scanner.Text(), ",")
	moves := stringSliceToIntSlice(movesInput)

	var boards [][][]int
	var board [][]int

	scanner.Scan() // Skip newline
	for scanner.Scan() {
		if err != nil {
			return nil, err
		}

		boardInput := scanner.Text()
		board = append(board, stringSliceToIntSlice(strings.Split(boardInput, " ")))

		if len(board) == 5 {
			boards = append(boards, board)
			board = make([][]int, 0)
			scanner.Scan() // Skip newline
		}
	}

	g := game {
		draw: moves,
		boards: boards,
	}

	return &g, scanner.Err()
}

func part1(g *game) {
	drawn := make(map[int]bool)
	var drawIdx int
	var draw int
	for {
		draw = g.draw[drawIdx]
		drawn[draw] = true
		if drawIdx < 5 {
			drawIdx += 1
			continue
		}

		for _, board := range g.boards {
			if boardComplete(board, drawn) {
				// printBoard(board)
				fmt.Printf("First %d | %d \n", draw, draw * sumUnmarked(board, drawn))
				return
			}
		}

		drawIdx += 1
		if drawIdx >= len(g.draw) {
			break
		}
	}
}

func part2(g *game) {
	drawn := make(map[int]bool)
	var drawIdx int
	var draw int
	remainingBoards := g.boards
	for {
		draw = g.draw[drawIdx]
		drawn[draw] = true
		if drawIdx < 5 {
			drawIdx += 1
			continue
		}

		var boardsToKeep [][][]int
		for _, board := range remainingBoards {
			if boardComplete(board, drawn) {
				// printBoard(board)
				if len(remainingBoards) == 1 {
					fmt.Printf("Last %d | %d \n", draw, draw * sumUnmarked(board, drawn))
				}
			} else {
				boardsToKeep = append(boardsToKeep, board)
			}
		}
		remainingBoards = boardsToKeep

		drawIdx += 1
		if drawIdx >= len(g.draw) {
			break
		}
	}
}

func main() {
	game, err := parseInput("4.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}

	part1(game)
	part2(game)
}
