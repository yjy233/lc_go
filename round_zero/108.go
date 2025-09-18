package roundzero

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l <= 0 {
		return nil
	}

	ind := l / 2
	return &TreeNode{
		Val:   nums[ind],
		Left:  sortedArrayToBST(nums[:ind]),
		Right: sortedArrayToBST(nums[ind+1:]),
	}
}
