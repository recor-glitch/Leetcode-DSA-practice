package LeetCode

// find: find the largest number from where the decreasing order starts
// given: a sorted list from increasing to decreasing order
// example: [1, 3, 4, 6, 7, 5, 2, 1] -> 7
// [1, 3, 6, 6, 6, 4, 3, 2, 1] -> 6
// [1, 2, 4, 5, 6, 6, 6, 5, 4, 2, 1] -> 6

func GetLargestNumber(arr []int) int {
	if len(arr) == 0 {
		return -1 // or handle empty case as needed
	}

	i := 0
	j := len(arr) - 1
	maxValue := -1

	for i <= j {

		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if arr[j] > maxValue {
			maxValue = arr[j]
		}

		i++
		j--
	}

	return maxValue
}
