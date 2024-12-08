package dayXX

import (
	"fmt"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/solver"
)

type DayXX struct{}

var _ solver.Solver = (*DayXX)(nil)

func New() *DayXX {
	return &DayXX{}
}

func (d *DayXX) Part1(lines []string) string {
	return fmt.Sprintf("%d", 123)
}

func (d *DayXX) Part2(lines []string) string {
	return fmt.Sprintf("%d", 123)
}
