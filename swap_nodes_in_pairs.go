package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}


func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var returnHead *ListNode
	for i := 0; ; i++ {
		if i == 0 {
			returnHead = head
		}
		aheadVal := head.Val
		head.Val = head.Next.Val
		head.Next.Val = aheadVal
		head = head.Next.Next
		if head == nil || head.Next == nil {
			break
		}
	}
	return returnHead
}
