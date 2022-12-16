package main

import (
	"fmt"
	"os"
	"strconv"
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

	overlaps := 0
	for _, line := range lines {
		if line != "" {
			elfSplit := strings.Split(line, ",")
			firstElf := elfSplit[0]
			secondElf := elfSplit[1]

			if isOverlap(generateElfRange(firstElf), generateElfRange(secondElf)) {
				overlaps++
			}
		}
	}

	fmt.Println(overlaps)
}

func isThereSubset(firstElf string, secondElf string) bool {
	firstElfRange := generateElfRange(firstElf)
	secondElfRange := generateElfRange(secondElf)

	return isSubset(firstElfRange, secondElfRange) || isSubset(secondElfRange, firstElfRange)
}

func generateElfRange(elf string) []int {
	roomSplit := strings.Split(elf, "-")
	beginning, _ := strconv.Atoi(roomSplit[0])
	end, _ := strconv.Atoi(roomSplit[1])

	result := []int{}

	for beginning <= end {
		result = append(result, beginning)
		beginning++
	}

	return result
}

func isSubset(firstElf []int, secondElf []int) bool {
	for _, room := range firstElf {
		if !contains(room, secondElf) {
			return false
		}
	}

	return true
}

func contains(val int, list []int) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}

	return false
}

func isOverlap(firstElf []int, secondElf []int) bool {
	for _, room := range firstElf {
		if contains(room, secondElf) {
			return true
		}
	}

	return false
}
