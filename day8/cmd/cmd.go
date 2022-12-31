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

	grid := [][]int{}

	for _, line := range lines {
		if line != "" {
			newRow := []int{}

			for index := range line {
				curr, _ := strconv.Atoi(line[index : index+1])
				newRow = append(newRow, curr)
			}

			grid = append(grid, newRow)
		}
	}

	// count each of the outer rows
	visibleTrees := 0
	numRows := len(grid)
	visibleTrees += numRows * 2

	// count each on the outer columns
	numCols := len(grid[0])
	visibleTrees += (numCols - 2) * 2

	for i := 1; i < len(grid)-1; i = i + 1 {
		for j := 1; j < len(grid[0])-1; j = j + 1 {
			if canBeSeen(grid, i, j) {
				visibleTrees += 1
			}
		}
	}

	fmt.Println("Part 1:", visibleTrees)

	maxScenic := 0

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			currentScenic := scenicScore(grid, i, j)

			if currentScenic > maxScenic {
				maxScenic = currentScenic
			}
		}
	}

	fmt.Println("Part 2:", maxScenic)
}

func canBeSeenNorth(grid [][]int, row int, col int) bool {
	for i := row - 1; i >= 0; i = i - 1 {
		if grid[i][col] >= grid[row][col] {
			return false
		}
	}

	return true
}

func canBeSeenSouth(grid [][]int, row int, col int) bool {
	for i := row + 1; i < len(grid); i = i + 1 {
		if grid[i][col] >= grid[row][col] {
			return false
		}
	}

	return true
}

func canBeSeenEast(grid [][]int, row int, col int) bool {
	for i := col - 1; i >= 0; i = i - 1 {
		if grid[row][i] >= grid[row][col] {
			return false
		}
	}

	return true
}

func canBeSeenWest(grid [][]int, row int, col int) bool {
	for i := col + 1; i < len(grid[0]); i = i + 1 {
		if grid[row][i] >= grid[row][col] {
			return false
		}
	}

	return true
}

func canBeSeen(grid [][]int, row int, col int) bool {
	north := canBeSeenNorth(grid, row, col)
	south := canBeSeenSouth(grid, row, col)
	east := canBeSeenEast(grid, row, col)
	west := canBeSeenWest(grid, row, col)
	return north || south || east || west
	//return canBeSeenNorth(grid, row, col) || canBeSeenSouth(grid, row, col) || canBeSeenEast(grid, row, col) || canBeSeenWest(grid, row, col)
}

func northTrees(grid [][]int, row int, col int) int {
	trees := 0

	for i := row - 1; i >= 0; i-- {
		if grid[i][col] >= grid[row][col] {
			return trees + 1
		} else {
			trees++
		}
	}

	return trees
}

func southTrees(grid [][]int, row int, col int) int {
	trees := 0

	for i := row + 1; i < len(grid); i++ {
		if grid[i][col] >= grid[row][col] {
			return trees + 1
		} else {
			trees++
		}
	}

	return trees
}

func eastTrees(grid [][]int, row int, col int) int {
	trees := 0

	for i := col - 1; i >= 0; i-- {
		if grid[row][i] >= grid[row][col] {
			return trees + 1
		} else {
			trees++
		}
	}

	return trees
}

func westTrees(grid [][]int, row int, col int) int {
	trees := 0

	for i := col + 1; i < len(grid[0]); i++ {
		if grid[row][i] >= grid[row][col] {
			return trees + 1
		} else {
			trees++
		}
	}

	return trees
}

func scenicScore(grid [][]int, row int, col int) int {
	north := northTrees(grid, row, col)
	south := southTrees(grid, row, col)
	east := eastTrees(grid, row, col)
	west := westTrees(grid, row, col)

	return north * south * east * west
}
