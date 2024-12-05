package day04

import (
	"fmt"
	"strings"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
)

type Day04 struct {
}

var _ solver.Solver = (*Day04)(nil)

func New() *Day04 {
	return &Day04{}
}

// Directions: Up, Down, Left, Right, and all Diagonals
var directions = []struct {
	dy, dx int
}{
	{-1, 0},  // Up
	{1, 0},   // Down
	{0, -1},  // Left
	{0, 1},   // Right
	{-1, -1}, // Top-left
	{-1, 1},  // Top-right
	{1, -1},  // Bottom-left
	{1, 1},   // Bottom-right
}

const (
	xmas = "XMAS"
	mas  = "MAS"
	sam  = "SAM"
)

func (d *Day04) Part1(lines []string) string {
	total := 0
	wordSearch := d.buildWordSearch(lines)
	for y, row := range wordSearch {
		for x, letter := range row {
			// only care about X
			if letter != "X" {
				continue
			}
			total += d.countXMAS(wordSearch, y, x)
		}
	}
	return fmt.Sprintf("%d", total)
}

func (d *Day04) Part2(lines []string) string {
	total := 0
	wordSearch := d.buildWordSearch(lines)
	for y, row := range wordSearch {
		for x, letter := range row {
			// only care about A
			if letter != "A" {
				continue
			}
			if d.isValidMASOrSAM(wordSearch, x-1, y-1, x+1, y+1) && d.isValidMASOrSAM(wordSearch, x+1, y-1, x-1, y+1) {
				total += 1
			}
		}
	}
	return fmt.Sprintf("%d", total)
}

func (d *Day04) countXMAS(wordSearch [][]string, y, x int) int {
	count := 0
	word := "X"
	// brute force because why not
	for _, direction := range directions {
		dx, dy := x, y
		for word != xmas && len(word) < 4 {
			dx, dy = dx+direction.dx, dy+direction.dy
			if !isInBounds(wordSearch, dx, dy) {
				break
			}
			nextLetter := wordSearch[dy][dx]
			word += nextLetter
		}
		if word == xmas {
			count++
		}
		// reset
		word = "X"
	}
	return count
}

// isInBounds checks if the point (x, y) is within the bounds of the matrix
func isInBounds(wordSearch [][]string, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(wordSearch) && y < len(wordSearch[0])
}

func (d *Day04) buildWordSearch(input []string) [][]string {
	wordSearch := make([][]string, len(input))
	for i, line := range input {
		wordSearch[i] = make([]string, len(line))
		for j, letter := range strings.Split(line, "") {
			wordSearch[i][j] = letter
		}
	}
	return wordSearch
}

// only check diagonals because who cares about the other directions
func (d *Day04) isValidMASOrSAM(wordSearch [][]string, dx1, dy1, dx2, dy2 int) bool {
	word := "*A*"

	if !isInBounds(wordSearch, dx1, dy1) || !isInBounds(wordSearch, dx2, dy2) {
		return false
	}

	// easy, just replace the placeholders and we'll find out if we get what we need
	word = strings.Replace(word, "*", wordSearch[dy1][dx1], 1)
	word = strings.Replace(word, "*", wordSearch[dy2][dx2], 1)

	if word != mas && word != sam {
		return false
	}

	return true
}

func (d *Day04) SolvePart1() {}