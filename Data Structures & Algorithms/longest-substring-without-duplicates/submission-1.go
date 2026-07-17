func lengthOfLongestSubstring(s string) int {

    l :=0
    max := 0
    seen := make(map[byte]bool) // default value = false

    for r := 0 ; r<len(s); r++{
        for seen[s[r]]{
            seen[s[l]]=false
            l++
        }

        seen[s[r]]=true

        if r-l+1>max{
            max=r-l+1
        }
    }
    return max

}
