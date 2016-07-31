package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	var head, last *ListNode = nil, nil
	carry := 0

	for p1 != nil && p2 != nil {
		carry = p1.Val + p2.Val + carry
		tmp := &ListNode{
			Val:  carry % 10,
			Next: nil,
		}
		carry = carry / 10
		if head == nil {
			head = tmp
			last = tmp
		} else {
			last.Next = tmp
			last = tmp
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	if p1 == nil {
		p1 = p2
	}
	for p1 != nil {
		carry = p1.Val + carry
		tmp := &ListNode{
			Val:  carry % 10,
			Next: nil,
		}
		carry = carry / 10
		if head == nil {
			head = tmp
			last = tmp
		} else {
			last.Next = tmp
			last = tmp
		}
		p1 = p1.Next
	}
	if carry != 0 {
		tmp := &ListNode{
			Val:  carry % 10,
			Next: nil,
		}
		last.Next = tmp
		last = last.Next
	}
	return head
}

func main() {
	/*var l1, l2, l3 ListNode
	l1.Val = 2
	l2.Val = 4
	l3.Val = 3
	l1.Next = &l2
	l2.Next = &l3
	var n1, n2, n3 ListNode
	n1.Val = 5
	n2.Val = 6
	n3.Val = 4
	n1.Next = &n2
	n2.Next = &n3*/

	var l1, l2 ListNode
	l1.Val = 1
	l2.Val = 8
	l1.Next = &l2
	var n1 ListNode
	n1.Val = 0

	ret := addTwoNumbers(&l1, &n1)
	for ret != nil {
		fmt.Print(ret.Val, " ")
		ret = ret.Next
	}
}
