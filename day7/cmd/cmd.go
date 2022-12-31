package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	DISK_SPACE     = 70000000
	REQUIRED_SPACE = 30000000
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
	tree := TreeNode{
		fileName: "/",
		size:     0,
		children: []*TreeNode{},
	}

	currentLeaf := &tree

	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		items := strings.Split(line, " ")

		if items[0] == "$" {
			if items[1] == "cd" {
				if items[2] == "/" {
					currentLeaf = &tree
				} else if items[2] == ".." {
					currentLeaf = currentLeaf.parent
				} else {
					for _, node := range currentLeaf.children {
						if node.fileName == items[2] {
							currentLeaf = node
						}
					}
				}
			}
		} else {
			if items[0] == "dir" {
				newNode := &TreeNode{
					fileName: items[1],
					size:     0,
					children: []*TreeNode{},
					parent:   currentLeaf,
				}

				currentLeaf.children = append(currentLeaf.children, newNode)
			} else {
				size, _ := strconv.Atoi(items[0])
				newNode := &TreeNode{
					fileName: items[1],
					size:     size,
					children: nil,
					parent:   currentLeaf,
				}

				currentLeaf.children = append(currentLeaf.children, newNode)
			}
		}
	}

	calculateDirSize(&tree)

	fmt.Println("Part 1:", getTotalSub100M(&tree))

	remainingSpace := DISK_SPACE - tree.size
	neededSpace := REQUIRED_SPACE - remainingSpace

	largerDirs := getAllDirsLarger(&tree, neededSpace)

	sort.Ints(largerDirs)

	fmt.Println("Part 2:", largerDirs[0])
}

func getAllDirsLarger(node *TreeNode, size int) []int {
	if node.children == nil {
		return []int{}
	}

	largerDirs := []int{}

	for _, child := range node.children {
		largerSubDirs := getAllDirsLarger(child, size)

		for _, largerSubDir := range largerSubDirs {
			largerDirs = append(largerDirs, largerSubDir)
		}
	}

	if node.size >= size {
		largerDirs = append(largerDirs, node.size)
	}

	return largerDirs
}

func calculateDirSize(node *TreeNode) {
	if node.size != 0 {
		return
	}

	sum := 0

	for _, child := range node.children {
		if child.size == 0 {
			calculateDirSize(child)
		}

		sum += child.size
	}

	node.size = sum
}

func getTotalSub100M(node *TreeNode) int {
	if node.children == nil {
		return 0
	}

	sum := 0

	for _, child := range node.children {
		sum += getTotalSub100M(child)
	}

	if node.size <= 100000 {
		sum += node.size
	}

	return sum
}

type TreeNode struct {
	fileName string
	size     int
	children []*TreeNode
	parent   *TreeNode
}
