package lib

import (
	"strconv"
	"strings"
)


func FetchSliceOfIntsInString(line string) ([]int, error) {
	var nums []int
	fields := strings.FieldsFunc(line, func(r rune) bool {
		return r == ' '
	})

	for _, field := range fields {
		if field == "" {
			continue
		}
		n, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}
