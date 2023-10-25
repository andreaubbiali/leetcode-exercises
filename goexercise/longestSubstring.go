package main

import "strings"

// Given a string s, find the length of the longest substring without repeating characters.

func main() {
	// println(lengthOfLongestSubstring("pwwkew"))
	println(betterSolution("dvdf"), "should be 3")
	println(betterSolution("bbbbbb"), "should be 1")
	println(betterSolution("pwwkew"), "should be 3")
	println(betterSolution(" "), "should be 1")
	println(betterSolution("abcabcbb"), "should be 3")
	println(betterSolution("pwwkew"), "should be 3")
	println(betterSolution("ckilbkd"), "should be 5")
	println(betterSolution("abba"), "should be 2")
	println(betterSolution("uqinntq"), "should be 4")
	println(betterSolution("wobgrovw"), "should be 6")
}

func betterSolution(s string) int {
	if len(s) == 0 {
		return 0
	}

	var startIdx int = 0
	var max int = 1

	charMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if v, found := charMap[s[i]]; found {
			m := i - startIdx
			println(string(s[startIdx]), string(s[i]), string(s[v+1]))
			if max < m {
				max = m
			}

			if s[v+1] == s[i] {
				startIdx = i
			} else {
				if charMap[byte(s[v+1])] > startIdx {
					startIdx = charMap[byte(s[v+1])]
				}
			}
		}

		charMap[s[i]] = i
	}

	if max < (len(s) - startIdx) {
		return len(s) - startIdx
	} else {
		return max
	}
}

var stopped uint8 = 0
var substrings []string

func lengthOfLongestSubstring(s string) int {
	stopped = 0
	substrings = make([]string, 0)

	bytes := []byte(s)
	if len(bytes) == 0 {
		return 0
	} else if len(bytes) == 1 {
		return 1
	}

	i := 0

	for {
		c := bytes[i]

		checkInSlice(c)

		i++
		if i >= len(bytes) {
			break
		}
	}

	max := 0
	for j := 0; j < len(substrings); j++ {
		if max < len(substrings[j]) {
			max = len(substrings[j])
		}
	}

	return max
}

func checkInSlice(char uint8) {

	for i := stopped; i < uint8(len(substrings)); i++ {
		if strings.Contains(substrings[i], string(char)) {
			stopped = i
		} else {
			substrings[i] = strings.Join([]string{substrings[i], string(char)}, "")
		}
	}
	substrings = append(substrings, string(char))
}
