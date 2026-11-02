func minEatingSpeed(piles []int, h int) int {
	 l, r := 1, 0

	 for _, p := range piles {
		if p > r {
			r = p
		}
	 }

	 res := r

	 for l <= r{

		mid := l + (r-l)/2
		time :=0

		for _, p := range piles {

			time += int(math.Ceil(float64(p)/float64(mid)))
		}
		if time<=h {
			res = mid
			r = mid - 1
		} else {
			l = mid +1
		}
	 }
	 return res

}
