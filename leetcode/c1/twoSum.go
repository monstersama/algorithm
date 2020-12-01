package main

import (
	"fmt"
)

// hashtable method
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
		fmt.Println(hashTable)
	}
	return nil
}

func main() {
	nums := []int{0, 2, 7, 4}
	target := 9
	fmt.Println(twoSum(nums, target))
}
