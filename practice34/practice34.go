package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		l1 = new(ListNode)
	}
	if l2 == nil {
		l2 = new(ListNode)
	}
	p := new(ListNode)
	p.Val = l1.Val + l2.Val
	if p.Val >= 10 {
		p.Val -= 10
		if l1.Next == nil {
			l1.Next = new(ListNode)
		}
		l1.Next.Val++
	}
	if l1.Next != nil || l2.Next != nil { //其中一个不为nil
		p.Next = addTwoNumbers(l1.Next, l2.Next)
	}

	return p
}

func main() {
	test3 := &ListNode{
		Val:  1,
		Next: nil,
	}
	tt2 := &ListNode{
		Val:  9,
		Next: nil,
	}
	tt3 := &ListNode{
		Val:  9,
		Next: tt2,
	}
	t := addTwoNumbers(test3, tt3)
	fmt.Println(t)
	fmt.Println(t.Next)
	fmt.Println(t.Next.Next)
}
