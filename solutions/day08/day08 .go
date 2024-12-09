package day08 

import (
	"fmt"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
)

type Day08  struct{}

var _ solver.Solver = (*Day08 )(nil)

func New() *Day08  {
	return &Day08 {}
}

func (d *Day08 ) Part1(lines []string) string {
	return fmt.Sprintf("%d", 123)
}

func (d *Day08 ) Part2(lines []string) string {
	return fmt.Sprintf("%d", 123)
}
