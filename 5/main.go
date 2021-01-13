package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type seatInfo struct {
	row    int
	column int
	seatID int
}

func main() {
	seatList, err := readStringLines("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	largestSeatID := searchLargest(seatList)
	fmt.Println(largestSeatID)

	firstMissingSeatID := findFirstMissing(seatList)
	fmt.Println(firstMissingSeatID)
}

func searchLargest(seats []seatInfo) int {
	var n int
	for _, seat := range seats {
		if seat.seatID > n {
			n = seat.seatID
		}
	}
	return n
}

func findFirstMissing(seats []seatInfo) int {
	largestID := searchLargest(seats)
	for i := largestID; i > 0; i-- {
		if !isSeatTaken(seats, i) {
			return i
		}
	}
	return -1
}

func isSeatTaken(seats []seatInfo, seatID int) bool {
	for _, seat := range seats {
		if seat.seatID == seatID {
			return true
		}
	}

	return false
}

func readStringLines(path string) ([]seatInfo, error) {
	memoryObjects := make([]seatInfo, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		memoryObject := readInputLine(line)
		if memoryObject == nil {
			return nil, err
		}
		memoryObjects = append(memoryObjects, *memoryObject)
	}
	return memoryObjects, nil
}

func readInputLine(line string) *seatInfo {
	rowString := line[:7]
	columnString := line[7:]
	rowString = strings.ReplaceAll(rowString, "F", "0")
	rowString = strings.ReplaceAll(rowString, "B", "1")
	columnString = strings.ReplaceAll(columnString, "L", "0")
	columnString = strings.ReplaceAll(columnString, "R", "1")
	row, rowErr := strconv.ParseInt(rowString, 2, 0)
	column, columnErr := strconv.ParseInt(columnString, 2, 0)

	if rowErr != nil || columnErr != nil {
		return nil
	}
	rowInt := int(row)
	columnInt := int(column)
	return &seatInfo{rowInt, columnInt, rowInt*8 + columnInt}
}
