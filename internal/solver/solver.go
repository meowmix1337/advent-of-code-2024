package solver

// Solver
// most solutions will be a number, but some can be a string.
// converting an int to a string is much easier
// than the opposite so these functions return a string
// for example, the answer is 12. It is easier to just fmt.Sprintf and convert to a string.
// as opposed to if these functions returned an int and the answer is "ABC"
type Solver interface {
	Part1(lines []string) string
	Part2(lines []string) string
}
