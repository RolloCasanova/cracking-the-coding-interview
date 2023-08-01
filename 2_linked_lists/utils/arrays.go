package utils

import (
	"fmt"
	"strconv"
)

func StringArrayToIntArray(strs []string) ([]int, error) {
	ints := make([]int, len(strs))

	for i, str := range strs {
		v, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("error converting string to int: %w", err)
		}

		ints[i] = v
	}

	return ints, nil
}
