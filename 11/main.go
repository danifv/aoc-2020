package main

import (
	"adventofcode-2020/readstringlines"
	"bufio"
	"fmt"
	"os"
)

func main() {
	s, err := readstringlines.ReadStringLines("input", bufio.ScanLines, seatParser)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	seats := interfaceToRune(s)
	rounds := 0
	change := true
	for change != false {
		fmt.Println(seats)
		seats, change = changeSeats(seats)
		rounds++
	}
	fmt.Println(countOccupied(seats))
}

func countOccupied(seats [][]rune) int {
	var occupied int
	for _, row := range seats {
		for _, s := range row {
			if s == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func interfaceToRune(seats []interface{}) [][]rune {
	runeSeats := make([][]rune, 0)
	for _, r := range seats {
		runeSeats = append(runeSeats, r.([]rune))
	}
	return runeSeats
}

func changeSeats(seats [][]rune) ([][]rune, bool) {
	rows := len(seats)
	columns := len(seats[0])

	newSeats := make([][]rune, rows)
	hasSeatingChanged := false
	var seatChanged bool
	for y, row := range seats {
		newSeats[y] = make([]rune, columns)
		for x := range row {
			newSeats[y][x], seatChanged = newSeatValue(seats, x, y)
			if seatChanged {
				hasSeatingChanged = true
			}
		}
	}
	return newSeats, hasSeatingChanged
}

func newSeatValue(seats [][]rune, xSelf int, ySelf int) (rune, bool) {
	var newSeat, oldSeat rune
	oldSeat = seats[ySelf][xSelf]
	occupiedNeighbours := occupiedNeighbours(seats, xSelf, ySelf)
	switch {
	case oldSeat == 'L' && occupiedNeighbours == 0:
		newSeat = '#'
	case oldSeat == '#' && occupiedNeighbours >= 5:
		newSeat = 'L'
	default:
		newSeat = oldSeat
	}
	seatChanged := newSeat != oldSeat
	return newSeat, seatChanged
}

func occupiedNeighbours(seats [][]rune, xSelf int, ySelf int) int {
	var occupiedNeighbours int
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}
			if isOccupiedInDirection(seats, xSelf, ySelf, i, j) {
				occupiedNeighbours++
			}
		}
	}
	return occupiedNeighbours

}

func isOccupiedInDirection(seats [][]rune, xSelf int, ySelf int, xDirection int, yDirection int) bool {
	for (ySelf+yDirection >= 0 && ySelf+yDirection < len(seats)) && (xSelf+xDirection >= 0 && xSelf+xDirection < len(seats[0])) {
		ySelf += yDirection
		xSelf += xDirection
		if seats[ySelf][xSelf] == '#' {
			return true
		}
		if seats[ySelf][xSelf] == 'L' {
			return false
		}
	}
	return false
}

func seatParser(scanUnit string) interface{} {
	return []rune(scanUnit)
}
