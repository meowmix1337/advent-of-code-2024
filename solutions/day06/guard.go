package day06

import (
	"github.com/dvan-sqsp/advent-of-code-2024/util"
)

type direction int

const (
	Up direction = iota
	Right
	Down
	Left
)

type pos struct{ x, y int }

type Guard struct {
	pos
	direction      direction
	numDistinctPos int
	visited        map[pos]bool
}

func NewGuard(x int, y int) *Guard {
	visited := make(map[pos]bool)
	visited[pos{x: x, y: y}] = true
	return &Guard{
		pos:            pos{x: x, y: y},
		direction:      Up, // always start up
		numDistinctPos: 1,  // starting pos is inclusive
		visited:        visited,
	}
}

func (g *Guard) HasLeftMap(patrolMap [][]string) bool {
	// determine if we hit a "#"
	nextX := g.x
	nextY := g.y

	// preemptively check the next position
	switch g.direction {
	case Up:
		nextY--
	case Right:
		nextX++
	case Down:
		nextY++
	case Left:
		nextX--
	}

	// we're free!
	if !util.IsInBounds(patrolMap, nextX, nextY) {
		return true
	}

	// we're still in bounds, check if we hit something blocking the guard
	// if so, turn
	nextPos := patrolMap[nextY][nextX]
	if nextPos == "#" {
		g.Turn90()
	}

	// actually take the action to move
	switch g.direction {
	case Up:
		g.MoveUp()
	case Right:
		g.MoveRight()
	case Down:
		g.MoveDown()
	case Left:
		g.MoveLeft()
	}

	if _, ok := g.visited[pos{g.x, g.y}]; !ok {
		g.visited[pos{x: g.x, y: g.y}] = true
		g.numDistinctPos++
	}

	return false
}

func (g *Guard) MoveUp() {
	g.y--
}

func (g *Guard) MoveDown() {
	g.y++
}

func (g *Guard) MoveRight() {
	g.x++
}

func (g *Guard) MoveLeft() {
	g.x--
}

func (g *Guard) Turn90() {
	g.direction++
	if g.direction > Left {
		g.direction = Up
	}

}
