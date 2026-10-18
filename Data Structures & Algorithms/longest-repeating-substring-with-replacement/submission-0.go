func characterReplacement(s string, k int) int {

	freq := make(map[rune]int)
	l := 0
	maxwindow :=0
	maxfreq :=0

	for r, val := range s {
		freq[val]++
		if ( freq[val]>maxfreq){
			maxfreq = freq[val]
		}

		windowlength := r-l+1

		if ( windowlength - maxfreq > k){
			freq[rune(s[l])]--
			l++
		}

		windowlength = r-l+1

		if(windowlength>maxwindow){
			maxwindow=windowlength
		}
	}
	return maxwindow

}
