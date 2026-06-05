func productExceptSelf(nums []int) []int {
	// Count how many zeros appear

	zerocount := 0
	prod := 1

	for _, val := range nums {
		if (val !=0){
			prod*=val
		} else {
			zerocount++
		}
	}

	res := make([]int, len(nums))

	if (zerocount > 1){
		return res
	}

	for i, num := range nums{
		if zerocount > 0{
			if num == 0{
				res[i]=prod
			} else{
				res[i]=0
			}
		} else {
			res[i]=prod/num
		}
	}
	return res

}
