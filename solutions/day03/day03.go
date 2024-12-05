package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
)

type Day03 struct {
	disabled bool
}

var (
	part1Reg = *regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	part2Reg = *regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)
	numReg   = *regexp.MustCompile(`\d{1,3},\d{1,3}`)
)

const (
	doCmd   = "do()"
	dontCmd = "don't()"
)

var _ solver.Solver = (*Day03)(nil)

func New() *Day03 {
	return &Day03{

		disabled: false,
	}
}

func (d *Day03) Part1(lines []string) string {
	total := 0
	total += d.processLine(lines[0], part1Reg, false)
	return fmt.Sprintf("%d", total)
}

func (d *Day03) Part2(lines []string) string {
	total := 0
	total += d.processLine(lines[0], part2Reg, true)
	return fmt.Sprintf("%d", total)
}

func (d *Day03) processLine(line string, rule regexp.Regexp, dosAndDonts bool) int {
	total := 0
	instructions := rule.FindAllString(line, -1)

	for _, instruct := range instructions {
		if d.skipInstruction(instruct, dosAndDonts) {
			continue
		}
		total += d.multiply(instruct)
	}

	return total
}

func (d *Day03) skipInstruction(instruction string, dosAndDonts bool) bool {
	if !dosAndDonts {
		return false
	}

	if instruction == dontCmd {
		d.disabled = true
	}
	if instruction == doCmd {
		d.disabled = false
		return true // we don't want to process do()s
	}
	if d.disabled {
		return true
	}
	return false
}

func (d *Day03) multiply(instruction string) int {
	nums := strings.Split(numReg.FindAllString(instruction, -1)[0], ",")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	return num1 * num2
}
