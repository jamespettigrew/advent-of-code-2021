package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	// "strconv"
	"strings"
)

type entry struct {
	signals [][]string
	digits [][]string
}

type state struct {
	knowns map[string][]string // logicalWiring -> possible signalComponents
	unmappedSignalComponents []string
}

var segments = [][]string {
	{ "A", "B", "C", "D", "E", "G" },
	{ "B", "C" },
	{ "A", "B", "D", "E", "F" },
	{ "A", "B", "C", "D", "F" },
	{ "B", "C", "F", "G" },
	{ "A", "C", "D", "F", "G" },
	{ "A", "C", "D", "E", "F", "G" },
	{ "A", "B", "C" },
	{ "A", "B", "C", "D", "E", "F", "G" },
	{ "A", "B", "C", "D", "F", "G" },
}

func indexOf(haystack []string, needle string) int {
	for i, val := range haystack {
		if val == needle {
			return i
		}
	}

	return -1
}

func remove(s []string, val string) []string {
	i := indexOf(s, val)
	if i == -1 {
		return s
	}

    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func equals(s1 []string, s2 []string) bool {
	set := make(map[string]int)
	for _, s := range s1 {
		set[s] += 1
	}
	for _, s := range s2 {
		set[s] += 1
	}

	for _, v := range set {
		if v != 2 {
			return false
		}
	}

	return true
}

func contains(haystack []string, needle string) bool {
    for _, a := range haystack {
        if a == needle {
            return true
        }
    }
    return false
}

func all(collection map[string][]string, predicate func([]string) bool) bool {
	for _, v := range collection {
		if !predicate(v) {
			return false
		}
	}

	return true
}

func parseInput(filePath string) ([]entry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []entry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawEntry := scanner.Text()
		if err != nil {
			return entries, err
		}
		split := strings.Split(rawEntry, " | ")
		rawSignals := strings.Split(split[0], " ")

		var signals [][]string
		for _, signal := range rawSignals {
			var signalComponents []string
			for _, component := range strings.Split(signal, "") {
				signalComponents = append(signalComponents, component)
			}

			signals = append(signals, signalComponents)
		}
		sort.Slice(signals, func(i, j int) bool {
			return len(signals[i]) < len(signals[j])
		})

		rawDigits := strings.Split(split[1], " ")
		var digits [][]string
		for _, rawDigit := range rawDigits {
			var digitComponents []string
			for _, component := range strings.Split(rawDigit, "") {
				digitComponents = append(digitComponents, component)
			}
			digits = append(digits, digitComponents)
		}
		entries = append(entries, entry { signals: signals, digits: digits })
	}
	return entries, nil
}

func decode(knowns map[string][]string, digitSignals [][]string) int {
	ret := 0
	decoder := make(map[string]string)
	for k, v := range knowns {
		decoder[v[0]] = k
	}

	for _, digitSignal := range digitSignals {
		digitSegments := make([]string, 0)
		for _, signalComponent := range digitSignal {
			digitSegments = append(digitSegments, decoder[signalComponent])
		}

		ret *= 10
		for digit, s := range segments {
			if equals(s, digitSegments) {
				fmt.Println(digit)
				ret += digit
			}
		}
	}

	return ret
}

func p1(entries []entry) int {
	sum := 0
	for _, entry := range entries {
		for _, digitComponents := range entry.digits {
			switch len(digitComponents) {
			case 2, 3, 4, 7:
				sum += 1
			}
		}
	}

	return sum
}

func p2(entries []entry) int {
	sum := 0

	// Compute initial knowns
	for _, entry := range entries {
		unmappedSignalComponents := []string { "a", "b", "c", "d", "e", "f", "g" }
		knowns := make(map[string][]string)
		for _, signalComponents := range entry.signals {
			displayedDigit := 0
			switch len(signalComponents) {
			case 2: // Display 1
				displayedDigit = 1
				break
			case 3: // Display 7
				displayedDigit = 7
				break
			case 4: // Display 4
				displayedDigit = 4
				break
			case 7: // Display 8
				displayedDigit = 8
				break
			default:
				continue
			}

			for _, segment := range segments[displayedDigit] {
				if len(knowns[segment]) == 0 {
					for _, signalComponent := range signalComponents {
						if contains(unmappedSignalComponents, signalComponent) {
							knowns[segment] = append(knowns[segment], signalComponent)
						}
					}
				}
			}

			for _, signalComponent := range signalComponents {
				unmappedSignalComponents = remove(unmappedSignalComponents, signalComponent)
			}
		}

		for _, segment := range []string { "B", "E", "F", } {
			for _, signalComponent := range knowns[segment] {
				possibleSignal := remove([]string { "a", "b", "c", "d", "e", "f", "g" }, signalComponent)
				for _, signalComponents := range entry.signals {
					if equals(possibleSignal, signalComponents) {
						knowns[segment] = []string { signalComponent }
						for otherSegment, _ := range knowns {
							if otherSegment == segment {
								continue
							}
							knowns[otherSegment] = remove(knowns[otherSegment], signalComponent)
						}
						break
					}
				}
			}
		}
		decoded := decode(knowns, entry.digits)
		fmt.Println(decoded)
		sum += decoded
	}


	return sum
}

func main() {
	input, err := parseInput("8.in")
	if err != nil {
		fmt.Println("Error reading file: %s", err)
	}
	part1 := p1(input)
	part2 := p2(input)
	fmt.Printf("Part 1: %d | Part 2: %d", part1, part2)
}
