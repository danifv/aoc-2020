package main

import (
	"adventofcode-2020/readstringlines"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	ns, err := readstringlines.ReadStringLines("input", bufio.ScanLines, numberParser)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	numbers := interfacesToInts(ns)
	numbers = append(numbers, 0)
	sort.Ints(numbers)
	deviceJolt := numbers[len(numbers)-1]
	numbers = append(numbers, deviceJolt+3)
	//numbers are ready at this point
	diffs := diffs(numbers)
	fmt.Println(diffs[1] * diffs[3])

	opts := options(numbers)
	fmt.Println(opts)
}

func options(numbers []int) int {
	optionsAtNumber := make(map[int]int)
	for i := len(numbers) - 1; i >= 0; i-- {
		currentNumber := numbers[i]
		optionsAtNumber[currentNumber] = optionsAtIndex(numbers, optionsAtNumber, i)
	}
	return optionsAtNumber[0]
}

func optionsAtIndex(numbers []int, optionsAtNumber map[int]int, index int) int {
	if index == len(numbers)-1 {
		return 1
	}
	currentNumber := numbers[index]
	var options int
	neighbourIndices := []int{index + 1, index + 2, index + 3}
	neighbourNumbers := make([]int, 0)
	for _, v := range neighbourIndices {
		if v < len(numbers) && numbers[v]-currentNumber < 4 {
			neighbourNumbers = append(neighbourNumbers, numbers[v])
		}
	}
	for _, v := range neighbourNumbers {
		options += optionsAtNumber[v]
	}
	return options
}

func diffs(numbers []int) map[int]int {
	diffs := make(map[int]int)
	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		diffs[diff]++
	}
	return diffs
}

func interfacesToInts(ns []interface{}) []int {
	numbers := make([]int, 0)
	for _, n := range ns {
		numbers = append(numbers, n.(int))
	}
	return numbers
}

func numberParser(scanUnit string) interface{} {
	number, err := strconv.Atoi(scanUnit)
	if err != nil {
		return nil
	}
	return number
}
