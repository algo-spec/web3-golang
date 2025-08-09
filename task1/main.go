package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1, 1, 3}
	rs := singleNumber(nums)
	fmt.Println("只出现一次的数字：", rs)

	x := 123212
	rs2 := isPalindrome(x)
	fmt.Println("回文数：", rs2)

	s := "()[]{}}"
	rs3 := isValid(s)
	fmt.Println("有效的括号：", rs3)

	strs := []string{"flower", "flow", "flight"}
	rs4 := longestCommonPrefix(strs)
	fmt.Println("最长公共前缀：", rs4)

	digits := []int{9, 9, 9}
	rs5 := plusOne(digits)
	fmt.Println("加一：", rs5)

	nums1 := []int{1, 1, 2}
	rs6 := removeDuplicates(nums1)
	fmt.Println("删除有序数组中的重复项：", rs6)

	nums2 := []int{2, 7, 11, 15}
	target := 9
	rs7 := twoSum(nums2, target)
	fmt.Println("两数之和：", rs7)
}

// 136. 只出现一次的数字
func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

// 9. 回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reverseNumber := 0
	for x > reverseNumber {
		reverseNumber = reverseNumber*10 + x%10
		x /= 10
	}
	return x == reverseNumber || x == reverseNumber/10
}

// 20. 有效的括号
func isValid(s string) bool {
	stack := Stack[string]{}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack.Push(string(char))
		} else {
			if stack.IsEmpty() || stack.Peek() != leftOf(char) {
				return false
			}
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}

func leftOf(char rune) string {
	switch char {
	case ')':
		return "("
	case ']':
		return "["
	case '}':
		return "{"
	}
	return ""
}

/*
Stack 栈，切片实现
*/
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) Peek() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	return item
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	s0 := strs[0]
	for j, c := range s0 {
		for _, s := range strs {
			if j == len(s) || s[j] != byte(c) {
				return s0[:j]
			}
		}
	}
	return s0
}

// 66. 加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

// 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != nums[slowIndex] {
			slowIndex++
			nums[slowIndex] = nums[fastIndex]
		}
	}
	return slowIndex + 1
}

// 56. 合并区间
// func merge(intervals [][]int) [][]int {

// }

// 1. 两数之和
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		other := target - num
		if p, ok := numMap[other]; ok {
			return []int{p, i}
		}
		numMap[num] = i
	}
	return nil
}
