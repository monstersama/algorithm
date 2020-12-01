package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	// m := map[byte]int
	size := len(s)
	lonestLen := 0 //结果
	right := -1    // 右指针
	for i := 0; i < size; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right+1 < size && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}
		lonestLen = max(lonestLen, right-i+1)
	}
	return lonestLen
}

func lengthOfLongestSubstring2(s string) int {
	m := make(map[byte]int)
	left, right := 0, 0
	size := len(s)
	longestLen := 0
	for right < size {
		if m[s[right]] == 0 {
			m[s[right]] = 1
			right++
		} else {
			delete(m, s[left])
			left++
		}
		if len(m) > longestLen {
			longestLen = len(m)
		}
	}
	return longestLen
}

func lengthOfLongestSubstring3(s string) int {
	m := [128]int{0}
	longest := 0
	left := 0
	for right := 0; right < len(s); right++ {
		left = max(left, m[s[right]]) //Determine whether the left pointer needs to point to a new position
		m[s[right]] = right + 1
		longest = max(longest, right-left+1)
	}
	return longest
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func main() {
	var s = "abcadf"
	fmt.Println("longest1:", lengthOfLongestSubstring(s))
	fmt.Println("longest2:", lengthOfLongestSubstring2(s))
	fmt.Println("longest2:", lengthOfLongestSubstring3(s))
}
