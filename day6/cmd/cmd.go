package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run cmd/cmd.go [input_file]")
		os.Exit(1)
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Could not read file:", os.Args[1])
	}

	line := string(file)
	fmt.Println(line)

	for i := 0; i < len(line) - 13; i++ {
		hash := make(map[byte]int)
		for j := 0; j <= 13; j++ {
			hash[line[i + j]]++
		}

		found_dups := false
		for _, v := range hash {
			if v != 1 {
				found_dups = true
			}
		}

		if !found_dups {
			fmt.Println(i + 14)
			os.Exit(0)
		}
	}
}
