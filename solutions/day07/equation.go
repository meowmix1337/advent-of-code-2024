package day07

import (
	"fmt"
	"strconv"
)

type Equation struct {
	targetValue int
	numbers     []int
}

func NewEquation(targetValue int, numbers []int) Equation {
	return Equation{targetValue: targetValue, numbers: numbers}
}

func (e *Equation) isEquationValid(idx int, currentNum int, concatMode bool) bool {
	// base case
	if idx == len(e.numbers) {
		// we've reached the end of the numbers, we should be able to determine if we match it
		return currentNum == e.targetValue
	}

	nextNum := e.numbers[idx]

	// branch off and try multiplication
	if e.isEquationValid(idx+1, currentNum*nextNum, concatMode) {
		return true
	}

	// branch off and try addition
	if e.isEquationValid(idx+1, currentNum+nextNum, concatMode) {
		return true
	}

	if concatMode {
		concatNumStr := fmt.Sprintf("%d%d", currentNum, nextNum)
		concatNum, _ := strconv.Atoi(concatNumStr)
		if e.isEquationValid(idx+1, concatNum, concatMode) {
			return true
		}
	}

	return false
}
