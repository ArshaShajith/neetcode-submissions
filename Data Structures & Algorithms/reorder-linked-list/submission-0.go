/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {

	if head==nil && head.Next==nil{
		return
	}

	slow, fast:= head,head

	for fast!=nil && fast.Next!=nil{
		slow=slow.Next
		fast= fast.Next.Next
	}

	second:= slow.Next
	slow.Next=nil

	var prev *ListNode

	for second!= nil{
		next:= second.Next
		second.Next= prev
		prev= second
		second= next
	}
    
	second= prev
	first:=head

	for second!=nil {
		next1 := first.Next
		next2 := second.Next

		first.Next=second
		second.Next= next1

		first= next1
		second=next2
	}
}
