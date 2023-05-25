package main

import (
	"fmt"
)

func combineArrays(array1, array2 []string) []string {
	combinedArray := make([]string, 0)

	// Combine arrays without duplicates
	for _, item := range array1 {
		if !contains(combinedArray, item) {
			combinedArray = append(combinedArray, item)
		}
	}

	for _, item := range array2 {
		if !contains(combinedArray, item) {
			combinedArray = append(combinedArray, item)
		}
	}

	return combinedArray
}

func contains(arr []string, item string) bool {
	for _, val := range arr {
		if val == item {
			return true
		}
	}
	return false
}

func noDuplicateArray(array1, array2 []string) []string {
	combinedArray := make([]string, 0)

	// Combine arrays without duplicates
	for _, item1 := range array1 {
		isDuplicate := false
		for _, item2 := range array2 {
			if item1 == item2 {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			combinedArray = append(combinedArray, item1)
		}
	}

	for _, item2 := range array2 {
		isDuplicate := false
		for _, item1 := range array1 {
			if item2 == item1 {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			combinedArray = append(combinedArray, item2)
		}
	}

	return combinedArray
}

func main() {
	array1 := []string{"a", "b", "c"}
	array2 := []string{"b", "c", "d"}

	combinedArray := combineArrays(array1, array2)
	noDuplicateArray := noDuplicateArray(array1, array2)

	fmt.Println("Combined Array:", combinedArray)        // Output: [a b c d]
	fmt.Println("NoDuplicated Array:", noDuplicateArray) // Output: [a d]
}
