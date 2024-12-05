package day01

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
)

type Day01 struct{}

var _ solver.Solver = (*Day01)(nil)

func New() *Day01 {
	return &Day01{}
}

func (d *Day01) Part1(lines []string) string {
	leftList, rightList := d.parseLines(lines)

	sort.Ints(leftList)
	sort.Ints(rightList)

	total := 0
	for i := 0; i < len(leftList); i++ {
		diff := math.Abs(float64(leftList[i] - rightList[i]))
		total += int(diff)
	}
	return fmt.Sprintf("%d", total)
}

func (d *Day01) Part2(lines []string) string {
	leftList, rightList := d.parseLines(lines)

	count := make(map[int]int)
	for _, num := range rightList {
		if _, ok := count[num]; ok {
			count[num]++
		} else {
			count[num] = 1
		}
	}

	total := 0
	for _, num := range leftList {
		total += num * count[num]
	}
	return fmt.Sprintf("%d", total)
}

func (d *Day01) parseLines(lines []string) ([]int, []int) {
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		leftNum, _ := strconv.Atoi(numbers[0])
		rightNum, _ := strconv.Atoi(numbers[1])
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	return leftList, rightList
}
