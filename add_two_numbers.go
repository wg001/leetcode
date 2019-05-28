package main

import "fmt"

func main() {
	//[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	b:=[]int{5,6,4}
	a := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	//a := []int{0}
	//b := []int{0}

	l1 := initListNode(a...)
	l2 := initListNode(b...)
	listNodeGo := addTwoNumbers(l1, l2)
	for ; listNodeGo != nil; listNodeGo = listNodeGo.Next {
		if listNodeGo.Next == nil {
			fmt.Printf("%d", listNodeGo.Val)
		} else {
			fmt.Printf("%d -> ", listNodeGo.Val)
		}
	}
}

func initListNode(a ...int) *ListNode {
	listNodeGo := &ListNode{Val: -1}

	for _, val := range a {
		if listNodeGo.Val < 0 && listNodeGo.Next == nil {
			listNodeGo.Val = val
		} else if listNodeGo == nil {
			listNodeGo.Next = insert(val)
		} else {
			dd := listNodeGo
			for ; dd.Next != nil; dd = dd.Next {

			}
			dd.Next = insert(val)
		}
	}
	if listNodeGo.Val < 0 {
		listNodeGo.Val = 0
	}
	return listNodeGo

	//for ; listNodeGo != nil; listNodeGo = listNodeGo.Next {
	//	if listNodeGo.Next == nil {
	//		fmt.Printf("%d", listNodeGo.Val)
	//	} else {
	//		fmt.Printf("%d -> ", listNodeGo.Val)
	//	}
	//}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carryVal := 0
	listNodeGo := &ListNode{Val: -1}
	for ; l1 != nil || l2 != nil; {
		tempVal := 0
		if l1 != nil && l1.Val > 0 {
			tempVal += l1.Val
		}
		if l2 != nil && l2.Val > 0 {
			tempVal += l2.Val
		}
		tempVal += carryVal
		carryVal = tempVal / 10
		val := tempVal % 10
		if listNodeGo.Val < 0 && listNodeGo.Next == nil {
			listNodeGo.Val = val
		} else if listNodeGo == nil {
			listNodeGo.Next = &ListNode{Val: val}
		} else {
			dd := listNodeGo
			for ; dd.Next != nil; dd = dd.Next {

			}
			dd.Next = &ListNode{Val: val}
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}
	if carryVal > 0 {
		dd := listNodeGo
		for ; dd.Next != nil; dd = dd.Next {

		}
		dd.Next = &ListNode{Val: carryVal}
	}
	if listNodeGo.Val == -1 {
		listNodeGo.Val = 0
	}

	return listNodeGo
}

func insert(val int) *ListNode {

	return &ListNode{Val: val}

}
