func maxArea(height []int) int {

    left := 0
    right := len(height)-1
    max := 0
    temp :=0

    for left < right {
        if(height[left]<height[right]){
            temp = height[left]*(right-left)
            left++
        } else {
            temp = height[right]*(right-left)
            right--
        }
        if(temp> max){
            max = temp
        }
    }
    return max
    
}