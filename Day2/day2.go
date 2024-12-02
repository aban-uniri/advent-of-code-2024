package main

import (
	"fmt"
	"os"

	"adventofcode2024/lib"
)

func main() {
	input, err := lib.ReadFileLineByLine("Day2/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input file:", err)
		os.Exit(1)
	}

	if err := getTotalSafeReportCount(input); err != nil {
		fmt.Fprintln(os.Stderr, "Error processing reports:", err)
		os.Exit(1)
	}
}


func getTotalSafeReportCount(reports []string) error {
	var countSafe, countSafeWithDeletion int

	for _, reportLine := range reports {
		reportNumbers, err := lib.FetchSliceOfIntsInString(reportLine)
		if err != nil {
			return fmt.Errorf("error parsing report: %v", err)
		}

		if isReportSafe(reportNumbers) {
			countSafe++
		} else if checkReportSafetyWithDeletion(reportNumbers) {
			countSafeWithDeletion++
		}
	}

	fmt.Printf("Safe reports: %d \n", countSafe)
	fmt.Printf("Safe reports with deletion: %d", countSafe+countSafeWithDeletion)
	return nil
}


func isReportSafe(reportNumbers []int) bool {
	if len(reportNumbers) < 2 {
		return true
	}

	var increasing bool
	directionSet := false

	for i := 1; i < len(reportNumbers); i++ {
		diff := reportNumbers[i] - reportNumbers[i-1]

		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}

		if !directionSet {
			increasing = diff > 0
			directionSet = true
		} else if (diff > 0) != increasing {
			return false
		}
	}

	return true
}


func checkReportSafetyWithDeletion(reportNumbers []int) bool {
	for i := range len(reportNumbers){
		if isReportSafeSkippingIndex(reportNumbers, i) {
			return true
		}
	}
	return false
}


func isReportSafeSkippingIndex(reportNumbers []int, skipIndex int) bool {
	var prevValue int
	var increasing bool
	directionSet := false
	initialized := false

	for i := range len(reportNumbers){
		if i == skipIndex {
			continue
		}
		if !initialized {
			prevValue = reportNumbers[i]
			initialized = true
			continue
		}

		diff := reportNumbers[i] - prevValue

		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}

		if !directionSet {
			increasing = diff > 0
			directionSet = true
		} else if (diff > 0) != increasing {
			return false
		}

		prevValue = reportNumbers[i]
	}

	return true
}
