package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	s := make([][]int, 0, 0)
	for i, v := range nums {
		s = append(s, []int{i, v})
	}
	nn := sort(s)
	for i := 0; i < len(nn); i++ {
		for j := len(nn) - 1; j > 0; j-- {
			ss := nn[i][1] + nn[j][1]
			if ss == target {
				return []int{nn[i][0], nn[j][0]}
			}
		}
	}
	return nil
}

func sort(s [][]int) [][]int {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i][1] > s[j][1] {
				s[i], s[j] = s[j], s[i]
				break
			}
		}
	}
	return s
}

func main() {
	var s = []int{3, 3}
	fmt.Println(twoSum(s, 6))
}
