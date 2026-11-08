func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1)>len(nums2){
		return findMedianSortedArrays(nums2,nums1)
	}

	m:= len(nums1)
	n:= len(nums2)

	l:=0
	r:=m
	half:=(m+n+1)/2

	for l<=r{
		i:= l+(r-l)/2
		j:= half-i

		maxLeft1:= math.MinInt
		maxLeft2:= math.MinInt
		minRight1:= math.MaxInt
		minRight2:= math.MaxInt

		if (i > 0){
			maxLeft1= nums1[i-1]
		}
		if (i<m){
			minRight1= nums1[i]
		}
		if (j>0){
			maxLeft2= nums2[j-1]
		}
		if (j<n){
			minRight2= nums2[j]
		}


		if maxLeft1<=minRight2 && maxLeft2<=minRight1{
			if (m+n)%2==1{
				return float64(max(maxLeft1,maxLeft2))
			}

			leftmax := max(maxLeft1,maxLeft2)
			rightmin := min(minRight1,minRight2)

			return float64(leftmax)/2.0 +float64(rightmin)/2.0
		} else if maxLeft1>minRight2{
			r = i-1
		} else {
			l= i+1
		}
	}
	return 0.0
}
