package flags

import (
	"strconv"
	"strings"
)

func parseIntRange(s string) ([]int, error) {
	arr := strings.Split(s, "-")
	start, err := strconv.Atoi(arr[0])
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(arr[1])
	if err != nil {
		return nil, err
	}

	if start > end {
		start, end = end, start
	}

	ls := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		ls = append(ls, i)
	}

	return ls, nil
}
