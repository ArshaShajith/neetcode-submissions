/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

	if head==nil {
		return nil
	}

	current := head

	for current!=nil{
		newNode := &Node{
			Val: current.Val,
		}
		newNode.Next = current.Next
		current.Next=newNode
		current=newNode.Next
	}

	current = head

	for current!= nil{
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current= current.Next.Next
	}
    
	current= head
	newcurr := current.Next

	copyhead := newcurr

	for current!=nil{
		current.Next=newcurr.Next 

		if newcurr.Next != nil{
			newcurr.Next = newcurr.Next.Next
		}

		current= current.Next
		newcurr=newcurr.Next
	}
	return copyhead
}
