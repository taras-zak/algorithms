package main

func nextPermutation(nums []int) {
	var k = -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			k = i
			break
		}
	}
	if k == -1 {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		return
	}
	var i int = 0
	for j := k + 1; j < len(nums); j++ {
		if nums[k] < nums[j] {
			i = j
		}
	}
	nums[k], nums[i] = nums[i], nums[k]
	for a, b := k+1, len(nums)-1; a < b; a, b = a+1, b-1 {
		nums[a], nums[b] = nums[b], nums[a]
	}
	return
}

func main() {

}
