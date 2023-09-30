package main

func twoSum(nums []int, target int) []int {

	lengthNums := len(nums)
	for i := 0; i < lengthNums; i++ {
		for j := 1; j < lengthNums; j++ {
			if i != j {
				if nums[i]+nums[j] == target {
					return []int{i, j}
				}
			}

		}
	}

	return make([]int, 2)

}
func main() {}
