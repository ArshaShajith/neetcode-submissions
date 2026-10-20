func isValid(s string) bool {

	stack := make([]byte,0,len(s))

	for i:=0; i<len(s); i++{
		ch:=s[i]

		if (ch=='(' || ch=='[' || ch=='{'){
			stack = append(stack,ch)
			continue
		}

		if (len(stack)==0){
			return false
		}

		top := stack[len(stack)-1]

		if((ch==')' && top!='(') || (ch==']' && top!='[') || (ch=='}' && top!='{') ){
			return false
		}

		stack=stack[:len(stack)-1]
	}
	return len(stack)==0
    
}
