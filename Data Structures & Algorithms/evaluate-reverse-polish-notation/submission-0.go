func evalRPN(tokens []string) int {
    stack := make([]int,0,len(tokens))
    var result int
    for i:=0; i<len(tokens);i++{
        num, err := strconv.Atoi(tokens[i])

        if (err==nil){
            stack=append(stack,num)
        } else {
            b := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            a := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            switch tokens[i]{
                case "+":
                    result=a+b
                case "-":
                    result=a-b
                case "*":
                    result=a*b
                case "/":
                    result=a/b
            }

            stack = append(stack,result)

        }
    }
    return stack[0]
}