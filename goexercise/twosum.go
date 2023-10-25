package main

func main() {
	var res = TwoSum([]int{3, 2, 4}, 6)
	if len(res) == 0 {
		println("Empty res")
	} else {
		println(res[0], res[1])
	}
}

// solution 1
func TwoSumMySolution(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// solution 2
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		if tmp, ok := m[target-nums[i]]; !ok {
			// no complementary found, add it into the map
			m[nums[i]] = i
		} else {
			// if ok -> in the m map there is the complementary of this num (tmp) that return target as sum result.
			return []int{i, tmp}
		}
	}
	return []int{}
}
