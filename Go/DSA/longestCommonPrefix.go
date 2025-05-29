// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

package DSA

func LongestCommonPrefix(str []string) string {

	result := str[0]
	for _, s := range str {
		i := 0
		for ; i < len(s) && i < len(result) && result[i] == s[i]; i++ {
		}
		result = result[:i]
	}
	return result
}
