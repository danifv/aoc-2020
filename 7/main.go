package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rules, err := readStringLines("input", bufio.ScanLines, ruleParser)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	shinyGoldHolders := countBagHolders(rules, "shiny gold")
	fmt.Println(len(shinyGoldHolders))
	bagsInGoldBag := countBagContent(rules, "shiny gold")
	fmt.Println(bagsInGoldBag)
}

func countBagContent(rules []interface{}, bagType string) int {
	totalContent := 0
	for _, rule := range rules {
		bagRule := rule.(map[string]map[string]int)
		matchingRule, isRuleMatching := bagRule[bagType]
		if isRuleMatching {
			for containedBag, numberContained := range matchingRule {
				totalContent += numberContained
				totalContent += numberContained * (countBagContent(rules, containedBag))
			}
		}
	}
	return totalContent
}

func countBagHolders(rules []interface{}, bagType string) map[string]bool {
	holders := make(map[string]bool)
	for _, rule := range rules {
		bagRule := rule.(map[string]map[string]int)
		for container, content := range bagRule {
			_, isBagContained := content[bagType]
			if isBagContained {
				holders[container] = true
				mergeContainers(holders, countBagHolders(rules, container))
			}
		}
	}
	return holders
}

func mergeContainers(firstContainer, secondContainer map[string]bool) map[string]bool {
	for k, v := range secondContainer {
		_, isContained := firstContainer[k]
		if !isContained {
			firstContainer[k] = v
		}
	}
	return firstContainer
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

func ruleParser(scanUnit string) interface{} {
	bagRules := make(map[string]map[string]int)

	ruleStrings := strings.Split(scanUnit, " bags contain ")
	container := ruleStrings[0]
	contentString := ruleStrings[1]

	contentStrings := strings.Split(contentString, ", ")
	contentList := make(map[string]int, 0)
	for _, content := range contentStrings {
		var bagPattern, bagColor string
		var bagNumber int
		fmt.Sscanf(content, "%v %v %v bag", &bagNumber, &bagPattern, &bagColor)
		bagType := bagPattern + " " + bagColor
		contentList[bagType] = bagNumber
	}
	bagRules[container] = contentList
	return bagRules

}
