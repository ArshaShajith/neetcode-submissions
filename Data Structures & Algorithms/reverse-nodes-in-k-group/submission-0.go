/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next:head}
	prevGroup := dummy

	for {
		groupEnd := prevGroup
		for i:=0; i<k && groupEnd!=nil ; i++{
			groupEnd = groupEnd.Next
		}
		if groupEnd == nil{
			break
		}

		groupNext:= groupEnd.Next
		groupStart:= prevGroup.Next

		// Reverse group

		prev := groupNext
		curr := groupStart

		for curr!=groupNext {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		prevGroup.Next = groupEnd
		prevGroup = groupStart
	}
    
	return dummy.Next
}
