// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:

// Input: x = 123
// Output: 321
// Example 2:

// Input: x = -123
// Output: -321
// Example 3:

// Input: x = 120
// Output: 21

// Constraints:

// -231 <= x <= 231 - 1

package DSA

import (
	"fmt"
	"strconv"
)

func ReverseInteger(x int32) int {
	quotient, reminder, negativeFlag := x, 0, false
	reversedString := ""

	if quotient < 0 {
		quotient = quotient * -1
		negativeFlag = true
	}

	for quotient != 0 {
		reminder = int(quotient) % 10
		quotient = quotient / 10

		reversedString += fmt.Sprint(reminder)
	}

	result, err := strconv.ParseInt(reversedString, 10, 32)
	if err != nil {
		fmt.Printf("Something went wrong while parsing to int")
	}

	if negativeFlag {
		result = result * -1
	}

	return int(result)
}
