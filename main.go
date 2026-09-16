package main

import "fmt"
import "sort"

type Pair struct {
	value int
	index int
}

func twoSum(nums []int, target int) []int {
	pairs := make([]Pair, len(nums))

	for i, num := range nums {
		pairs[i] = Pair{num, i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	left := 0
	right := len(pairs) - 1

	for left < right {
		sum := pairs[left].value + pairs[right].value

		if sum == target {
			return []int{pairs[left].index, pairs[right].index}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	fmt.Println(result)
}
