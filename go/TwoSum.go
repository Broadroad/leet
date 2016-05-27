package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		if v, ok := m[num]; ok {
			return []int{v, index}
		}
		m[target-nums[index]] = index
	}
	return nil
}

func main() {
	s := []int{3, 5, 2, 6, 1, 4}
	fmt.Println(s)
	fmt.Println(twoSum(s, 6))
}
