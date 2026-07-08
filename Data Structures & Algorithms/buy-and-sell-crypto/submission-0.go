func maxProfit(prices []int) int {
    max :=0
    min:= prices[0]
    for i:=0 ; i<len(prices);i++ {
        if(min > prices[i]){
            min=prices[i]
        }
        profit := prices[i]-min

        if(max<profit){
            max=profit
        }
    }
    return max
}