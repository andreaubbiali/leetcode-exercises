// Given a string s, return the longest palindromic substring in s.

package main

func main() {
	// println(longestPalindrome("babad"), "expected: bab or aba")
	// println(longestPalindrome("cbbd"), "expected: bb")
	// println(longestPalindrome("bddewwefjn"), "expected: ewwe")
	println(longestPalindrome("abcba"), "expected: abcba")

}

func longestPalindrome(s string) string {
	var res string = string(s[0])

	if len(s) == 1 {
		return res
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}

	for i := 1; i < len(s); i++ {
		println("index: ", i)
		for j := 1; j < len(s); j++ {
			if i-j < 0 || j+i > len(s) {
				break
			}
			s1 := s[i-j : i]
			s2 := s[i : j+i]
			// println(s1, s2)
			if isPalindroma(s1, s2) {
				tmp := s[i-j : j+i]
				if len(tmp) > len(res) {
					res = tmp
				}
			}

			// check is palindroma removing index letter
			if i-j < 0 || j+i+1 > len(s) {
				break
			}
			s1 = s[i-j : i]
			s2 = s[i+1 : j+i+1]
			// println(s1, s2)
			if isPalindroma(s1, s2) {
				tmp := s[i-j : j+i+1]
				if len(tmp) > len(res) {
					res = tmp
				}
			}
		}
	}

	return res
}

func isPalindroma(s1 string, s2 string) bool {
	// println(s1, s2)
	if len(s1) != len(s2) {
		panic("wrong length")
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[len(s1)-i-1] {
			return false
		}
	}

	return true
}
