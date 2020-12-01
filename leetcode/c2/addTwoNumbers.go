package main

import (
	// "debug/plan9obj"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = new(ListNode)
	var node = head
	var c = 0 //进位
	var v1 = 0
	var v2 = 0
	for l1 != nil || l2 != nil || c != 0 {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		c = c + v1 + v2
		var s = ListNode{Val: c % 10}
		(*node).Next = &s
		node = &s
		c = c / 10
		v1 = 0
		v2 = 0
	}

	// for head != nil{
	//     fmt.Println(*head)
	//     head=head.Next
	// }
	return head.Next
	// return head.Next
	// for l1 != nil{
	//     fmt.Println(*l1)
	//     l1 = l1.Next
	// }
	// for l2 != nil{
	//     println(l2)
	//     l2 = l2.Next
	// }
	// return l1.Next
}

func main() {
	l1 := [4]int{1, 1, 1, 1}
	l2 := [3]int{5, 6, 4}
	// var p = new(int)
	// var p *int

	// fmt.Println(p)
	var head1 = new(ListNode) // equity var ln ListNode; head := &ln
	var node1 *ListNode
	node1 = head1
	// fmt.Println(head1, node1)
	for _, val := range l1 {
		var n = ListNode{Val: val}
		(*node1).Next = &n
		node1 = &n
	}

	var head2 = new(ListNode)
	var node2 *ListNode
	node2 = head2
	for _, val := range l2 {
		var n = ListNode{Val: val}
		(*node2).Next = &n
		node2 = &n
	}
	// 遍历链表
	// for head1 != nil{
	//     fmt.Println(*head1)
	//     head1=head1.Next
	// }

	var head = new(ListNode)
	head = addTwoNumbers(head1.Next, head2.Next)
	for head != nil {
		fmt.Println(*head)
		head = head.Next
	}

	// fmt.Println(addTwoNumbers())
}
