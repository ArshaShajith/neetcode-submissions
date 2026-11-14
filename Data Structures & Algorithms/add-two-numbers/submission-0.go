/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	result := dummy

	carry :=0
	num1,num2 := l1, l2

	for num1!=nil || num2!=nil || carry!=0{
		digit1 :=0
		digit2 := 0

		if num1!=nil{
			digit1 = num1.Val
			num1= num1.Next
		}

		if num2 != nil {
			digit2 = num2.Val
			num2 = num2.Next
		}

		total := digit1+digit2+carry
		digit := total%10
		carry = total/10

		result.Next=&ListNode{
			Val: digit,
		}
		result=result.Next
	}
	return dummy.Next
    
}
