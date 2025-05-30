// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

// Example 1:

// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
// Example 2:

// Input: digits = ""
// Output: []
// Example 3:

// Input: digits = "2"
// Output: ["a","b","c"]

package DSA

func LetterCombinationOfPhoneNumber(digits string) []string {
	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := []string{}
	var backtrack func(string, string)

	if len(digits) == 0 {
		return result
	}

	backtrack = func(combination string, next_digit string) {
		if len(next_digit) == 0 {
			result = append(result, combination)
		} else {
			letters := phone[next_digit[0]]
			for i := range letters {
				backtrack(combination+string(letters[i]), next_digit[1:])
			}
		}
	}

	backtrack("", digits)
	return result
}
