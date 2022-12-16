package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run cmd/cmd.go [input_files]")
		os.Exit(1)
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Could not read file:", os.Args[1])
		os.Exit(1)
	}

	lines := strings.Split(string(file), "\n")

	storage := make(map[int]string)

	readingCrates := true

	for _, line := range lines {
		if readingCrates {
			if line[:2] == " 1" {
				readingCrates = false
				continue
			}

			if line != "" {
				cratePlacements := []string{}
				index := 0

				for index < len(line) {
					cratePlacements = append(cratePlacements, line[index:index+3])
					index += 4
				}

				for i, crate := range cratePlacements {
					if crate != "   " {
						crateLetter := string(strings.Split(crate, "[")[1][0])
						storage[i+1] = crateLetter + storage[i+1]
					}
				}
			}
		} else if line == "" {
			continue
		} else {
			re := regexp.MustCompile("[0-9]+")
			values := re.FindAllString(line, -1)

			numToMove, _ := strconv.Atoi(values[0])
			moveFrom, _ := strconv.Atoi(values[1])
			moveTo, _ := strconv.Atoi(values[2])

			cratesToPlace := storage[moveFrom][len(storage[moveFrom])-numToMove:]
			storage[moveFrom] = storage[moveFrom][:len(storage[moveFrom])-numToMove]
			storage[moveTo] = storage[moveTo] + cratesToPlace
		}
	}

	result := ""

	for i := 1; i <= len(storage); i++ {
		result = result + string(storage[i][len(storage[i])-1])
	}

	fmt.Println(result)
}
