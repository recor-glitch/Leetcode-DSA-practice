// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5

package DSA

import "fmt"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedList := make([]int, len(nums2))
	mergedList = append(nums1, mergedList...)

	i, j := len(nums1)-1, len(nums2)-1
	k := (len(nums1) + len(nums2)) - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			mergedList[k] = nums1[i]
			i--
		} else {
			mergedList[k] = nums2[j]
			j-- 
		}
		k--
	}

	fmt.Printf("My merged list: %+v", mergedList)

	return float64(mergedList[len(mergedList)/2])
}
