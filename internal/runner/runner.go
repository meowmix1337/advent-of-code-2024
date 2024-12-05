package runner

import (
	"fmt"

	"github.com/dvan-sqsp/advent-of-code-2024/internal/parser"
	"github.com/dvan-sqsp/advent-of-code-2024/solutions/day_factory"
	"github.com/rs/zerolog/log"
)

func Run(day int) {
	// get the solver for the specified day
	solver := day_factory.GetDay(day)
	input, err := loadInput(day)
	if err != nil {
		log.Error().Err(err).Msg("failed to load input")
		return
	}

	// print the solutions
	fmt.Printf("Part 1: %s\n", solver.Part1(input))
	fmt.Printf("Part 2: %s\n", solver.Part2(input))
}

func loadInput(day int) ([]string, error) {
	filename := fmt.Sprintf("input/day%02d.txt", day)
	return parser.ReadLines(filename)
}
