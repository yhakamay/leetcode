package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
}

func deleteDuplicates(head *ListNode) (dummy *ListNode) {
	dummy = head

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}

	return
}
