package util

import (
	"strconv"
	"strings"
)

// ReadInts will convert the line into a slice of ints
// helpful for lines like 1,2,3,4,5,6,7
func ReadInts(line string, sep string) ([]int, error) {
	var ints []int
	for _, s := range strings.Split(line, sep) {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
