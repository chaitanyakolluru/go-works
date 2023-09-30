package main

func singleNumber(nums []int) int {
	counter := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if counter[nums[i]] == 0 {
			counter[nums[i]] = 1
			continue
		}

		if counter[nums[i]] == 1 {
			delete(counter, nums[i])
		}
	}

	keys := make([]int, 0, len(counter))
	for k := range counter {
		keys = append(keys, k)
	}

	return keys[0]
}
func main() {}
