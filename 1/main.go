package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	numbers, err := scanLines("input")
	if err != nil {
		panic(err)
	}

	pair1, pair2, found := findSumPair(numbers, 2020)
	if !found {
		fmt.Println("No matching pair found")
		return
	}
	fmt.Printf("\nPair found:\n%d * %d = %d\n", pair1, pair2, pair1*pair2)

	triplet1, triplet2, triplet3, found := findSumTriplet(numbers, 2020)
	if !found {
		fmt.Println("No matching triplet found")
		return
	}
	fmt.Printf("\nTriplet found:\n%d * %d * %d = %d\n", triplet1, triplet2, triplet3, triplet1*triplet2*triplet3)
}

func findSumPair(numbers []int, sum int) (int, int, bool) {
	diffs := make(map[int]bool)
	for _, number := range numbers {
		if diffs[number] {
			return number, sum - number, true
		}
		diffs[sum-number] = true
	}
	return 0, 0, false
}

func findSumTriplet(numbers []int, sum int) (int, int, int, bool) {

	for i, triplet1 := range numbers {
		triplet2, triplet3, found := findSumPair(numbers[i:], sum-triplet1)
		if found {
			return triplet1, triplet2, triplet3, true
		}
	}
	return 0, 0, 0, false
}

func scanLines(path string) ([]int, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var numbers []int

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}
