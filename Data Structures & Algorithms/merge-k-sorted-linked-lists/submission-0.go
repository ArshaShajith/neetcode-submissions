/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode{
	dummy := &ListNode{}
	current := dummy
	for l1!=nil && l2!=nil{
		if l1.Val<l2.Val{
			current.Next= l1
			l1 = l1.Next
		} else{
			current.Next = l2
			l2 = l2.Next
		}
		current=current.Next
	}

	if l1 != nil{
		current.Next=l1
	}
	if l2!= nil{
		current.Next=l2
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists)==0{
		return nil
	}

	for len(lists)>1{
		var mergedList []*ListNode
		for i:=0; i<len(lists);i+=2{
		l1 := lists[i]
		var l2 *ListNode

		if i+1<len(lists){
			l2= lists[i+1]
		}
		mergedList= append(mergedList, mergeTwoLists(l1,l2))
		}
		lists = mergedList
	}
	return lists[0]

}