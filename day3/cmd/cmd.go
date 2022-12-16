package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run cmd/cmd.go [input_file]")
		os.Exit(1)
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Could not read file:", os.Args[1])
		os.Exit(1)
	}

	lines := strings.Split(string(file), "\n")

	currentElves := []string{}
	sum := 0
	for _, line := range lines {
		if line != "" {
			currentElves = append(currentElves, line)

			if len(currentElves) == 3 {
				sharedChar := findSharedChar(currentElves)
				sum += convertToValue(sharedChar)

				currentElves = []string{}
			}
		}
	}

	fmt.Println(sum)
}

func findSharedChar(elfList []string) rune {
	for _, char := range elfList[0] {
		if strings.Contains(elfList[1], string(char)) && strings.Contains(elfList[2], string(char)) {
			return char
		}
	}

	return 'a'
}

func convertToValue(char rune) int {
	if char >= 'a' && char <= 'z' {
		return int(char-'a') + 1
	}

	return int(char-'A') + 26 + 1
}
