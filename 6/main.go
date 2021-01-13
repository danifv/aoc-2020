package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	groups, err := readStringLines("input", answerGroupSplitFunc, voteParser)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	groupVotes := groupVotes(groups)
	distinctVotes := distinctGroupVoteSum(groupVotes)
	commonVotes := commonGroupVoteSum(groups)
	fmt.Println(distinctVotes)
	fmt.Println(commonVotes)
}

func groupVotes(groups []interface{}) []map[rune]int {
	groupVoteList := make([]map[rune]int, 0)
	for _, group := range groups {
		voteMap := make(map[rune]int)
		persons := group.([]string)
		for _, person := range persons {
			for _, vote := range person {
				voteMap[vote]++
			}
		}
		groupVoteList = append(groupVoteList, voteMap)
	}
	return groupVoteList
}

func distinctGroupVoteSum(groupVoteList []map[rune]int) int {
	var distinctVotes int
	for _, groupVotes := range groupVoteList {
		distinctVotes += len(groupVotes)
	}
	return distinctVotes
}

func commonGroupVoteSum(groups []interface{}) int {
	var commonVotes int
	for _, group := range groups {
		persons := group.([]string)
		for _, vote := range persons[0] {
			isCommonVote := true
			for _, person := range persons {
				if !strings.Contains(person, string(vote)) {
					isCommonVote = false
					break
				}
			}
			if isCommonVote {
				commonVotes++
			}
		}
	}
	return commonVotes
}

type scanUnitParser func(scanUnit string) interface{}

func readStringLines(path string, splitFunc bufio.SplitFunc, scanUnitParser scanUnitParser) ([]interface{}, error) {
	memoryObjects := make([]interface{}, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		scanUnit := scanner.Text()
		memoryObject := scanUnitParser(scanUnit)
		if err != nil {
			return nil, err
		}
		memoryObjects = append(memoryObjects, memoryObject)
	}
	return memoryObjects, nil
}

func answerGroupSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return len(data), nil, nil
	}

	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 2, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}
	return
}

func voteParser(scanUnit string) interface{} {
	personVotes := strings.Split(scanUnit, "\n")
	return personVotes
}
