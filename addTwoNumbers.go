package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var newNode *ListNode = &ListNode{
		Val: 2,
	}
	var secondNode *ListNode = &ListNode{
		Val: 4,
	}
	newNode.Next = secondNode

	var thirdNode *ListNode = &ListNode{
		Val: 9,
	}
	secondNode.Next = thirdNode

	var newNode1 *ListNode = &ListNode{
		Val: 5,
	}
	var secondNode2 *ListNode = &ListNode{
		Val: 6,
	}
	newNode1.Next = secondNode2

	var thirdNode2 *ListNode = &ListNode{
		Val: 4,
	}
	secondNode2.Next = thirdNode2

	var fourthNode2 *ListNode = &ListNode{
		Val: 9,
	}
	thirdNode2.Next = fourthNode2

	result := addTwoNumbers(newNode, newNode1)
	for result != nil {
		fmt.Println("result %s", result.Val)
		result = result.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newNode = new(ListNode)
	var nowNodePtr *ListNode = newNode
	nums := 0
	for l1 != nil || l2 != nil || nums > 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += nums
		nowNodePtr.Val = sum % 10
		nums = sum / 10

		if l1 != nil || l2 != nil || nums > 0 {
			var nextNode = new(ListNode)
			nowNodePtr.Next = nextNode
			nowNodePtr = nextNode
		}

	}

	return newNode
}
