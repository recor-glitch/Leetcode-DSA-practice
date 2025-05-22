package DSA

// Longest Substring Without Repeating Characters

// Given a string s, find the length of the longest substring without duplicate characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

func LongestSubString(s string) int {
	curBest := 0
	set := map[byte]bool{}
	l, r := 0, 0
	for r < len(s) {
		if _, ok := set[s[r]]; !ok { // can grow window
			set[s[r]] = true
			r++
			curBest = max(curBest, r-l)
		} else {
			delete(set, s[l])
			l++
		}
	}
	return curBest
}
