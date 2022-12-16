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

	sum := 0
	hash := generateHash()
	for _, line := range lines {
		sum += hash[line]
	}

	fmt.Println(sum)
}

func generateHash() map[string]int {
	hash := make(map[string]int)

	hash["A X"] = 3 // them rock, I lose
	hash["A Y"] = 4 // them rock, we tie
	hash["A Z"] = 8 // them rock, I win
	hash["B X"] = 1 // them paper, I lose
	hash["B Y"] = 5 // them paper, we tie
	hash["B Z"] = 9 // them paper, I win
	hash["C X"] = 2 // them scissors, I lose
	hash["C Y"] = 6 // them scissors, we tie
	hash["C Z"] = 7 // them scissors, I win

	return hash
}
