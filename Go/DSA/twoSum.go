package DSA

func TwoSum(nums []int, target int) []int {
	myMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := myMap[complement]; found {
			return []int{nums[index], nums[i]}
		}
		myMap[num] = i
	}

	return nil
}
