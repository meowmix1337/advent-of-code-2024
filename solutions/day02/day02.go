package day02

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
)

type Day02 struct{}

var _ solver.Solver = (*Day02)(nil)

func New() *Day02 {
	return &Day02{}
}

func (d *Day02) Part1(lines []string) string {
	totalSafe := 0

	reports := d.parseLines(lines)

	for _, report := range reports {
		if d.isSafe(report) {
			totalSafe++
		}
	}

	return fmt.Sprintf("%d", totalSafe)
}

func (d *Day02) Part2(lines []string) string {
	totalSafe := 0

	reports := d.parseLines(lines)

	for _, report := range reports {
		if d.isSafe(report) {
			totalSafe++
		} else if d.isSafeWithDampener(report) {
			totalSafe++
		}
	}

	return fmt.Sprintf("%d", totalSafe)
}

func (d *Day02) parseLines(lines []string) [][]int {
	var reports [][]int
	for _, line := range lines {
		nums := strings.Fields(line)
		report := make([]int, len(nums))
		for i, num := range nums {
			report[i], _ = strconv.Atoi(num)
		}
		reports = append(reports, report)
	}
	return reports
}

func (d *Day02) isSafe(report []int) bool {
	// Check that the numbers are within range (1 - 3)
	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i] - report[i-1]))
		if diff > 3 || diff == 0 {
			return false
		}
	}

	// Check whether the sequence is either all increasing or all decreasing
	increasing := true
	decreasing := true
	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1] {
			decreasing = false
		}
		if report[i] < report[i-1] {
			increasing = false
		}
	}

	return increasing || decreasing
}

func (d *Day02) isSafeWithDampener(report []int) bool {
	// Try removing each level and check if the remainder is safe
	for i := 0; i < len(report); i++ {
		// Create a copy of the slice without modifying the original
		modifiedReport := make([]int, 0, len(report)-1)
		modifiedReport = append(modifiedReport, report[:i]...)   // Add all elements before index i
		modifiedReport = append(modifiedReport, report[i+1:]...) // Add all elements after index i
		if d.isSafe(modifiedReport) {
			return true
		}
	}
	return false
}
