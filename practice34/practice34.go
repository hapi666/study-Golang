package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	added := new(ListNode)
	fmt.Println(added.Val)
	p:=added
	fmt.Println(p)
	for l1 != nil && l2 != nil {
		fmt.Println("here!")
		num := l1.Val + l2.Val
		//fmt.Println(num)
		g := num / 10
		if g == 0 {
			added.Val += num
			fmt.Println(added.Val)
		}else {
			added.Val += num % 10
			fmt.Println(added.Val)
			added.Next=new(ListNode)
			added.Next.Val = g
		}

		l1 = l1.Next
		l2 = l2.Next
		added = added.Next
	}
	return p
}

func main() {
	test1 := &ListNode{
		Val:  3,
		Next: nil,
	}
	test2 := &ListNode{
		Val:  4,
		Next: test1,
	}
	test3:=&ListNode{
		Val:2,
		Next:test2,
	}
	// tt := &ListNode{}
	tt1 := &ListNode{
		Val:  4,
		Next: nil,
	}
	tt2 := &ListNode{
		Val:  6,
		Next: tt1,
	}
	tt3:=&ListNode{
		Val:5,
		Next:tt2,
	}
	t := addTwoNumbers(test3, tt3)
	fmt.Println(t)
}
