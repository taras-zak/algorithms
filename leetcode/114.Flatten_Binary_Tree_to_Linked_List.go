package main

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			maxLeft := curr.Left
			for maxLeft.Right != nil {
				maxLeft = maxLeft.Right
			}
			maxLeft.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}
}

func main() {

}
