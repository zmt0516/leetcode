func findWords(words []string) []string {
	var r []string
	emap := make(map[byte]int, 52)
	s := "qwertyuiopQWERTYUIOP"
	for _, v := range s {
		emap[byte(v)] = 0
	}
	s = "asdfghjklASDFGHJKL"
	for _, v := range s {
		emap[byte(v)] = 1
	}
	s = "zxcvbnmZXCVBNM"
	for _, v := range s {
		emap[byte(v)] = 2
	}
LABEL:
	for _, v := range words {
		for i := 1; i < len(v); i++ {
			if emap[v[i]] != emap[v[i-1]] {
				continue LABEL
			}
		}
		r = append(r, v)

	}

	return r

}