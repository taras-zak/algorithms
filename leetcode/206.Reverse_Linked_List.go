package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head
		head = head.Next
		temp.Next = prev
		prev = temp
	}
	return prev
}

func main() {
	res := reverseList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}})
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
