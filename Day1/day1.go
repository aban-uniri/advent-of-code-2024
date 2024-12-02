package main

import (
	"fmt"
	"os"
	"sort"

	"adventofcode2024/lib"
)

func main() {
	input, err := lib.ReadFileLineByLine("Day1/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input file:", err)
		os.Exit(1)
	}

	numArrays, err := getNumArraysFromColumns(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing input:", err)
		os.Exit(1)
	}

	sort.Ints(numArrays[0])
	sort.Ints(numArrays[1])

	fmt.Println("Total Distance", getTotalDistance(numArrays))
	fmt.Println("Similarity Score:", getSimilarityScore(numArrays))
}

func getNumArraysFromColumns(input []string) ([][]int, error) {
	output := make([][]int, 2)

	for _, line := range input {
		nums, err := lib.FetchSliceOfIntsInString(line)
		if err != nil {
			return nil, err
		}
		if len(nums) < 2 {
			return nil, fmt.Errorf("not enough numbers in line: %s", line)
		}
		output[0] = append(output[0], nums[0])
		output[1] = append(output[1], nums[1])
	}

	return output, nil
}


func getTotalDistance(arrays [][]int) int {
	sum := 0
	for i := range arrays[0] {
		sum += abs(arrays[0][i] - arrays[1][i])
	}
	return sum
}


func getSimilarityScore(arrays [][]int) int {
	score := 0
	countMap := make(map[int]int)

	for _, num := range arrays[1] {
		countMap[num]++
	}

	for _, num := range arrays[0] {
		if count, exists := countMap[num]; exists {
			score += num * count
		}
	}
	return score
}


func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
