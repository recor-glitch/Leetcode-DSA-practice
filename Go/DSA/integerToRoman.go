package DSA

func IntegerToRoman(num int) string {
	valueSlice := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	keySlice := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	for num > 0 {
		for idx, val := range valueSlice {
			if num >= val {
				result += keySlice[idx]
				num -= val
				break
			}
		}
	}
	return result
}
