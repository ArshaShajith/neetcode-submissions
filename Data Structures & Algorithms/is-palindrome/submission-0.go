func isPalindrome(s string) bool {
    l := 0
    r := len(s)-1

    for l < r {
        for l<r && !isAlphanum(rune(s[l])){
            l++
        }
        for l<r && !isAlphanum(rune(s[r])){
            r--
        }
        if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])){
            return false
        }
        l++
        r--
    }
    return true
}

func isAlphanum(c rune) bool {
    return unicode.IsLetter(c) || unicode.IsDigit(c)
}
