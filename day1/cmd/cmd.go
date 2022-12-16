package main

import (
	"fmt"
	"os"
	"sort"
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

	calorieList := []int{}
	currentMax := 0
	for _, line := range lines {
		if line == "" {
			calorieList = append(calorieList, currentMax)
			currentMax = 0
		} else {
			val, _ := strconv.Atoi(line)
			currentMax += val
		}
	}

	sort.Ints(calorieList)

	fmt.Println(calorieList[len(calorieList)-3] + calorieList[len(calorieList)-2] + calorieList[len(calorieList)-1])
}
