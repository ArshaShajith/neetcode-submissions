type Solution struct{}

func (s *Solution) Encode(strs []string) string{
	res := ""
	for _, val := range strs{
		res += strconv.Itoa(len(val))+"#"+val
	}
	return res
}

func (s *Solution) Decode(strs string) []string{
	res := []string{}
	i := 0

	for i<len(strs) {
		j:=i
		for(strs[j] != '#'){
			j++
		}

		length, _ := strconv.Atoi(strs[i:j])
		i = j+1

		res = append(res, strs[i:i+length])
		i = i+length
	}
	return res
}
