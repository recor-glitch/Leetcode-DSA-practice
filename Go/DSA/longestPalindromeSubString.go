// Given a string s, return the longest palindromic substring in s.

// Example 1:

// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:

// Input: s = "cbbd"
// Output: "bb"

// Constraints:

// 1 <= s.length <= 1000
// s consist of only digits and English letters.

package DSA

func GetLongestPalindromeSubString(s string) string {
	myStr := ""
	j := len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[i] == s[j] {
			myStr += string(s[i])
		}
		j--
	}

	if myStr == "" {
		return string(s[0])
	} else {
		return myStr
	}
}
