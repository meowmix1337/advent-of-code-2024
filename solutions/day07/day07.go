package day07

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
)

type Day07 struct{}

var _ solver.Solver = (*Day07)(nil)

func New() *Day07 {
	return &Day07{}
}

func (d *Day07) Part1(lines []string) string {
	equations := d.parseEquations(lines)

	validEquations := d.ValidateEquations(equations, false)
	d.ValidateEquationsWaitGroup(equations, false)

	return fmt.Sprintf("%d", d.sumValidEquations(validEquations))
}

func (d *Day07) Part2(lines []string) string {
	equations := d.parseEquations(lines)

	validEquations := d.ValidateEquations(equations, true)

	d.ValidateEquationsWaitGroup(equations, true)

	return fmt.Sprintf("%d", d.sumValidEquations(validEquations))
}

func (d *Day07) parseEquations(lines []string) []Equation {
	equations := make([]Equation, 0)
	for _, line := range lines {
		equationStr := strings.Split(line, ":")
		targetVal, _ := strconv.Atoi(equationStr[0])
		numberStrs := strings.TrimSpace(equationStr[1])

		numbers := make([]int, 0)
		for _, numberStr := range strings.Split(numberStrs, " ") {
			num, _ := strconv.Atoi(numberStr)
			numbers = append(numbers, num)
		}

		equations = append(equations, NewEquation(targetVal, numbers))
	}

	return equations
}

func (d *Day07) ValidateEquations(equations []Equation, concatMode bool) []Equation {
	start := time.Now()
	validEquations := make([]Equation, 0)
	for _, equation := range equations {
		// recursively determine if the equations are valid
		if equation.isEquationValid(1, equation.numbers[0], concatMode) {
			validEquations = append(validEquations, equation)
		}
	}

	end := time.Now()
	totalTime := end.Sub(start)
	fmt.Print("Total Time: ", totalTime, "\n")

	return validEquations
}

func (d *Day07) ValidateEquationsWaitGroup(equations []Equation, concatMode bool) []Equation {
	start := time.Now()
	validEquations := make([]Equation, 0)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, equation := range equations {
		wg.Add(1)
		go func(equation Equation) {
			defer wg.Done()
			// recursively determine if the equations are valid
			if equation.isEquationValid(1, equation.numbers[0], concatMode) {
				mu.Lock()
				validEquations = append(validEquations, equation)
				mu.Unlock()
			}
		}(equation)
	}

	wg.Wait()
	end := time.Now()

	totalTime := end.Sub(start)
	fmt.Print("Total Time WG: ", totalTime, "\n")

	return validEquations
}

func (d *Day07) sumValidEquations(equations []Equation) int {
	sum := 0
	for _, equation := range equations {
		sum += equation.targetValue
	}
	return sum
}
