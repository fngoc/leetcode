package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	fmt.Println(easy.IsPalindrome(11233211))

	fmt.Println(easy.TwoSum([]int{2, 7, 11, 15}, 9))

	fmt.Println(easy.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(easy.LongestCommonPrefix([]string{"dog", "racecar", "car"}))

	fmt.Println(easy.IsValid("()"))
	fmt.Println(easy.IsValid("()[]{}"))
	fmt.Println(easy.IsValid("(]"))
	fmt.Println(easy.IsValid(""))
	fmt.Println(easy.IsValid("){"))

	fmt.Println(easy.PlusOne([]int{1, 2, 3}))
	fmt.Println(easy.PlusOne([]int{1, 6, 0}))
	fmt.Println(easy.PlusOne([]int{1, 7, 9}))
	fmt.Println(easy.PlusOne([]int{9}))
	fmt.Println(easy.PlusOne([]int{9, 9, 9, 9, 9, 9}))

	fmt.Println(easy.RemoveDuplicates([]int{1, 1, 2}))
	fmt.Println(easy.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	fmt.Println(easy.IsPalindromeStr("asdas as"))
	fmt.Println(easy.IsPalindromeStr("asdas asdas LLL das"))
	fmt.Println(easy.IsPalindromeStr("A man, a plan, a canal: Panama"))
	fmt.Println(easy.IsPalindromeStr(" "))
	fmt.Println(easy.IsPalindromeStr("0P"))
	fmt.Println(easy.IsPalindromeStr("0P0"))

	easy.Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	easy.Merge([]int{1}, 1, []int{}, 0)
	easy.Merge([]int{0}, 0, []int{1}, 1)
	easy.Merge([]int{2, 0}, 1, []int{1}, 1)
}
