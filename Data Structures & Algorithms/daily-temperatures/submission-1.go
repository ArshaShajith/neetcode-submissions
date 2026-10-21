func dailyTemperatures(temp []int) []int {
	n := len(temp)
	result := make([]int, n)
	stack := []int{}

	for i:=n-1; i>=0; i--{
		for len(stack)>0 && temp[stack[len(stack)-1]]<=temp[i]{
			stack = stack[:len(stack)-1]
		}

		if len(stack)>0{
			result[i]=stack[len(stack)-1]-i
		}

		stack = append(stack,i)
	}
	return result
}
