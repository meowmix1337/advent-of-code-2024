package day05

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
)

type Day05 struct {
}

var _ solver.Solver = (*Day05)(nil)

func New() *Day05 {
	return &Day05{}
}

func (d *Day05) Part1(lines []string) string {
	total := 0
	ruleSet, updates := d.parse(lines)

	for _, update := range updates {
		if d.isValid(update, ruleSet) {
			total += d.getMid(update)
		}
	}
	return fmt.Sprintf("%d", total)
}

func (d *Day05) Part2(lines []string) string {
	total := 0

	ruleSet, updates := d.parse(lines)

	incorrectUpdates := make([][]string, 0)
	for _, update := range updates {
		if !d.isValid(update, ruleSet) {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	for _, update := range incorrectUpdates {
		for !d.isValid(update, ruleSet) {
			d.sortByRuleSet(update, ruleSet)
		}
		total += d.getMid(update)
	}

	return fmt.Sprintf("%d", total)
}

func (d *Day05) parse(lines []string) (map[string][]string, [][]string) {
	rulesSet := make(map[string][]string)
	updates := make([][]string, 0)
	for _, line := range lines {
		if line == "" {
			continue
		} else if strings.Contains(line, ",") {
			updates = append(updates, strings.Split(line, ","))
		} else if strings.Contains(line, "|") {
			// rules
			rules := strings.Split(line, "|")
			if _, ok := rulesSet[rules[0]]; !ok {
				rulesSet[rules[0]] = make([]string, 0)
				rulesSet[rules[0]] = append(rulesSet[rules[0]], rules[1])
			} else {
				rulesSet[rules[0]] = append(rulesSet[rules[0]], rules[1])
			}
		}
	}

	return rulesSet, updates
}

func (d *Day05) isValid(updateList []string, ruleSet map[string][]string) bool {
	for idx, num := range updateList {
		nextIdx := idx + 1
		// Just check the next number, if it is within the current rules set, it is valid
		if nextIdx != len(updateList) && !slices.Contains(ruleSet[num], updateList[nextIdx]) {
			return false
		}
	}

	// if all the numbers pass, we're good
	return true
}

func (d *Day05) sortByRuleSet(updateList []string, ruleSet map[string][]string) {
	for idx, num := range updateList {
		nextIdx := idx + 1
		if nextIdx != len(updateList) && !slices.Contains(ruleSet[num], updateList[nextIdx]) {
			d.swap(updateList, idx, nextIdx, num)
		}
	}
}

func (d *Day05) swap(updateList []string, baseIdx int, idxToSwap int, baseNum string) {
	numToSwap := updateList[idxToSwap]
	updateList[baseIdx] = numToSwap
	updateList[idxToSwap] = baseNum
}

func (d *Day05) getMid(updateList []string) int {
	midNum, _ := strconv.Atoi(updateList[int(math.Floor(float64(len(updateList)/2)))])
	return midNum
}
