// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"

// Output: true

// Example 2:

// Input: s = "()[]{}"

// Output: true

// Example 3:

// Input: s = "(]"

// Output: false

// Example 4:

// Input: s = "([])"

// Output: true

// Constraints:

// 1 <= s.length <= 104
// s consists of parentheses only '()[]{}'.

package DSA

func IsValidParentheses(s string) bool {

	parentheses := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	stack := []string{}

	for _, letter := range s {
		if len(stack) > 0 && parentheses[string(letter)] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, string(letter))
	}
	return len(stack) == 0
}
