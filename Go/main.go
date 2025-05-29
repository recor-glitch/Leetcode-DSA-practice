package main

import (
	"DSA-practice/DSA"
	Queue "DSA-practice/Problems/Array/Queue"
	Stack "DSA-practice/Problems/Array/Stack"
	"DSA-practice/Problems/LeetCode"
	"fmt"
)

func main() {

	// STACK
	stack := Stack.NewEmptyStack[string]()

	stack.Push("First")
	stack.Push("Second")
	stack.Push("Third")
	stack.Push("Forth")

	ele, ok := stack.Pop()
	if !ok {
		fmt.Println("Error while pop")
	}

	fmt.Println("Deleted Element: ", ele)

	fmt.Println("Top item: ", stack.Top())

	// SIMPLE QUEUE
	queue := Queue.NewEmptyQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	element, ok := queue.Dequeue()
	if !ok {
		fmt.Print("Dequeue functionality failed!")
	}

	fmt.Println("Dequeued element: ", element)

	fmt.Println("First element in the Queue: ", queue.Front())

	// CIRCULAR QUEUE
	cq := Queue.NewCircularQueue[string](5)

	cq.Enqueue("One")
	cq.Enqueue("Two")
	cq.Enqueue("Three")
	cq.Enqueue("Four")

	cqEle, ok := cq.Dequeue()
	if !ok {
		fmt.Println("Dequeue failed!")
	}

	fmt.Println("Dequeued element: ", cqEle)

	cqFront, ok := cq.Front()
	if ok {
		fmt.Println("Front of the circular queue: ", cqFront)
	}

	// LEETCODE

	lgNum := LeetCode.GetLargestNumber([]int{1, 2, 3, 4, 5, 6, 6, 6, 5, 4, 3, 2, 1})

	fmt.Println("My largest number: ", lgNum)

	// TWO SUMS
	results := DSA.TwoSum([]int{3, 4, 6, 3, 8, 3, 2}, 9)
	fmt.Printf("My two sum result: %+v", results)

	// LONGEST SUBSTRING
	LongestSubStringNo := DSA.LongestSubString("ABCABCAS")
	fmt.Printf("My longest sub string: %d", LongestSubStringNo)

	// Find Median from Sorted Arrays
	medianFromSortedArrays := DSA.FindMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Printf("My Median from sorted arrays: %f", medianFromSortedArrays)

	// Longest Palindrome SubString
	longestPalindromeSubString := DSA.GetLongestPalindromeSubString("ad")
	fmt.Printf("My longest palindrome sub stirng: %s", longestPalindromeSubString)

	// Reverse Integer
	reversedInt := DSA.ReverseInteger(123)
	fmt.Printf("My reversed integer: %d", reversedInt)

	// Is Palindrome Number
	IsPalindromeNumber := DSA.IsPalindromeNumber(121)
	fmt.Println("My Palindrome result:", IsPalindromeNumber)

	// Container With Most Water
	containerArea := DSA.GetMaxWaterArea([]int{1, 2})
	fmt.Println("My container with max area: ", containerArea)

	// Integer to Roman
	intToRomanString := DSA.IntegerToRoman(3999)
	fmt.Println("Integer to Roman: ", intToRomanString)

	// Longest Common Prefix
	longestCommonPrefix := DSA.LongestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println("Longest common prefix: ", longestCommonPrefix)
}
