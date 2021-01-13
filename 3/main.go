package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	snowmap, err := readStringLines("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	trees1 := getSlopeTrees(snowmap, 1, 1)
	trees2 := getSlopeTrees(snowmap, 3, 1)
	trees3 := getSlopeTrees(snowmap, 5, 1)
	trees4 := getSlopeTrees(snowmap, 7, 1)
	trees5 := getSlopeTrees(snowmap, 1, 2)
	fmt.Println(trees1 * trees2 * trees3 * trees4 * trees5)
}

func getSlopeTrees(snowmap [][]string, horizontalStep, verticalStep int) int {
	var treeCount int
	height := len(snowmap)
	width := len(snowmap[0])
	j := 0
	for i := 0; i < height; i += verticalStep {
		if snowmap[i][j] == "#" {
			treeCount++
		}
		j = (j + horizontalStep) % width
	}
	fmt.Println(treeCount)
	return treeCount
}

func readStringLines(path string) ([][]string, error) {
	memoryObjects := make([][]string, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		memoryObject := readInputLine(line)
		if err != nil {
			return nil, err
		}
		memoryObjects = append(memoryObjects, [][]string{memoryObject}...)
	}
	return memoryObjects, nil
}

func readInputLine(line string) []string {
	return strings.Split(line, "")
}
