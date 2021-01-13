package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type passwordEntry struct {
	first     int
	second    int
	character rune
	password  string
	valid     bool
}

func main() {
	entries, err := readStringLines("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	var totalValid int
	for _, entry := range entries {
		count := strings.Count(entry.password, string(entry.character))
		if count >= entry.first && count <= entry.second {
			totalValid++
		}
	}
	fmt.Println(totalValid)

	totalValid = 0
	for _, entry := range entries {
		passwordSlice := strings.Split(entry.password, "")
		if passwordSlice[entry.first-1] == string(entry.character) && passwordSlice[entry.second-1] != string(entry.character) ||
			passwordSlice[entry.first-1] != string(entry.character) && passwordSlice[entry.second-1] == string(entry.character) {
			totalValid++
		}
	}
	fmt.Println(totalValid)
}

func readStringLines(path string) ([]passwordEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	passwordEntries := make([]passwordEntry, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := readLineAsPasswordEntry(line)
		if err != nil {
			return nil, err
		}
		passwordEntries = append(passwordEntries, entry)
	}
	return passwordEntries, nil
}

func readLineAsPasswordEntry(line string) (passwordEntry, error) {
	entry := new(passwordEntry)
	_, err := fmt.Sscanf(line, "%d-%d %c: %s", &entry.first, &entry.second, &entry.character, &entry.password)
	return *entry, err
}
