package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passportValidityList, err := readStringLines("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	validPassports := getTotalValidPassports(passportValidityList)
	fmt.Println(validPassports)
}

func getTotalValidPassports(passports []bool) int {
	var passportCount int
	for _, passportValidity := range passports {
		if passportValidity == true {
			passportCount++
		}
	}
	return passportCount
}

func readStringLines(path string) ([]bool, error) {
	memoryObjects := make([]bool, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(passwordSplitFunc)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "\n", " ")
		memoryObject := readInputLine(line)
		memoryObjects = append(memoryObjects, memoryObject)
	}
	return memoryObjects, nil
}

func passwordSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return len(data), nil, nil
	}

	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func getValueSubstr(data string, key string) (string, bool) {
	if !strings.Contains(data, key) {
		return "", false
	}
	valueStart := strings.Index(data, key) + len(key)
	valueEnd := strings.IndexAny(data[valueStart:], " ")
	if valueEnd == -1 {
		valueEnd = len(data)
	} else {
		valueEnd = valueEnd + valueStart
	}
	value := data[valueStart:valueEnd]
	return value, true
}

func hasValidByr(line string) bool {
	byrStr, found := getValueSubstr(line, "byr:")
	if !found {
		return false
	}
	byr, err := strconv.Atoi(byrStr)
	if err != nil {
		return false
	}
	return byr >= 1920 && byr <= 2002
}

func hasValidIyr(line string) bool {
	iyrStr, found := getValueSubstr(line, "iyr:")
	if !found {
		return false
	}
	iyr, err := strconv.Atoi(iyrStr)
	if err != nil {
		return false
	}
	return iyr >= 2010 && iyr <= 2020
}

func hasValidEyr(line string) bool {
	eyrStr, found := getValueSubstr(line, "eyr:")
	if !found {
		return false
	}
	eyr, err := strconv.Atoi(eyrStr)
	if err != nil {
		return false
	}
	return eyr >= 2020 && eyr <= 2030
}

func hasValidHgt(line string) bool {
	hgt, found := getValueSubstr(line, "hgt:")
	if !found {
		return false
	}

	hgtMagnitude, err := strconv.Atoi(hgt[:len(hgt)-2])
	if err != nil {
		return false
	}
	if strings.HasSuffix(hgt, "in") {
		return hgtMagnitude >= 59 && hgtMagnitude <= 76
	}
	if strings.HasSuffix(hgt, "cm") {
		return hgtMagnitude >= 150 && hgtMagnitude <= 193
	}
	return false
}

func hasValidHcl(line string) bool {
	hcl, found := getValueSubstr(line, "hcl:")
	if !found {
		return false
	}
	matched, _ := regexp.MatchString(`^(#[0-9a-f]{6})$`, hcl)
	return matched
}

func hasValidEcl(line string) bool {
	ecl, found := getValueSubstr(line, "ecl:")
	if !found {
		return false
	}
	validEclList := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	_, hasValidEcl := validEclList[ecl]
	return hasValidEcl
}

func hasValidPid(line string) bool {
	pid, found := getValueSubstr(line, "pid:")
	if !found {
		return false
	}
	matched, _ := regexp.Match(`^\d{9}$`, []byte(pid))
	return matched
}

func readInputLine(line string) bool {

	/*
		byr (Birth Year)
		iyr (Issue Year)
		eyr (Expiration Year)
		hgt (Height)
		hcl (Hair Color)
		ecl (Eye Color)
		pid (Passport ID)
		cid (Country ID)
	*/

	byrFound := hasValidByr(line)
	iyrFound := hasValidIyr(line)
	eyrFound := hasValidEyr(line)
	hgtFound := hasValidHgt(line)
	hclFound := hasValidHcl(line)
	eclFound := hasValidEcl(line)
	pidFound := hasValidPid(line)

	return byrFound && iyrFound && eyrFound && hgtFound && hclFound && eclFound && pidFound
}
