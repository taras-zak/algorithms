package main

func findTarget(root *TreeNode, k int) bool {
	nums := make([]int, 0)
	nums = inorder(root, nums)
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == k {
			return true
		}
		if sum > k {
			j--
		} else {
			i++
		}
	}
	return false
}

func inorder(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}
	nums = inorder(root.Left, nums)
	nums = append(nums, root.Val)
	return inorder(root.Right, nums)
}
