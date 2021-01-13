package main

import (
	"adventofcode-2020/readstringlines"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type call struct {
	instruction string
	value       int
}

func main() {
	calls, err := readstringlines.ReadStringLines("input", bufio.ScanLines, callParser)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	accumulator, _ := calculateAccumulator(calls)
	fmt.Println(accumulator)
	correctAccumulator := calculateAccumulatorWithInstructionChange(calls)
	fmt.Println(correctAccumulator)
}
func calculateAccumulator(calls []interface{}) (int, bool) {
	accumulator := 0
	executedCalls := make(map[int]bool)
	for i := 0; i < len(calls); {
		c := calls[i].(call)
		executedCalls[i] = true
		accumulatorChange, callOfset := executeCall(c)
		accumulator += accumulatorChange
		i += callOfset
		if wasCallExecuted(executedCalls, i) {
			return accumulator, true
		}
	}
	return accumulator, false
}

func wasCallExecuted(executedCalls map[int]bool, callNumber int) bool {
	_, wasExecuted := executedCalls[callNumber]
	return wasExecuted
}

func calculateAccumulatorWithInstructionChange(calls []interface{}) int {
	var accumulator int
	var hasLoop bool
	for i := 0; i < len(calls); i++ {
		originalCall := calls[i].(call)
		newcall, callChanged := changeCall(calls[i].(call))
		if callChanged {
			calls[i] = newcall
			accumulator, hasLoop = calculateAccumulator(calls)
			if !hasLoop {
				break
			}
			calls[i] = originalCall
		}
	}
	return accumulator
}

func changeCall(c call) (call, bool) {
	switch c.instruction {
	case "jmp":
		c.instruction = "nop"
		return c, true
	case "nop":
		c.instruction = "jmp"
		return c, true
	}
	return c, false
}

func executeCall(c call) (accumulatorChange int, callOffset int) {
	switch c.instruction {
	case "acc":
		return c.value, 1
	case "jmp":
		return 0, c.value
	case "nop":
		return 0, 1
	default:
		return 0, 0
	}
}

func callParser(scanUnit string) interface{} {

	callStrings := strings.Split(scanUnit, " ")
	instruction := callStrings[0]
	value, err := strconv.Atoi(callStrings[1])
	if err != nil {
		return nil
	}
	return call{instruction, value}
}
