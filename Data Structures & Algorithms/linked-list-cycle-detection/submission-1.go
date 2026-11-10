/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {

	visited := make(map[*ListNode]bool)
	current := head

	for current!=nil{
		if visited[current]{
			return true
		}
		visited[current]=true
		current=current.Next

	}
	return false
    
}
