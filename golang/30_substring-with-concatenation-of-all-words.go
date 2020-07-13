func findSubstring(s string, words []string) []int {

	var r []int
	if len(s) == 0 || len(words) == 0 || len(words[0]) == 0 {
		return r
	}
	lenw := len(words[0])
	lena := len(words[0]) * len(words)
	emap := make(map[string]int)
	for _, v := range words {
		emap[v] += 1
	}
	lmap := make(map[string]int)
	n := 1
	for i := range emap {
		lmap[i] = n
		n++
	}

	amap := make([]int, len(s))
	for i := 0; i < len(s)-lenw+1; i++ {
		amap[i] = lmap[s[i:i+lenw]]
	}

	//fmt.Println(amap)

	nmap := make([]int, len(words)+1)
	for _, v := range words {
		nmap[lmap[v]] = emap[v]
	}

	for i := 0; i < lenw; i++ {

		l := i
		nmap2 := make([]int, len(words)+1)
		copy(nmap2, nmap)
		if i+lena > len(s) {
			break
		}
		for j := l; j < l+lena; j += lenw {
			nmap2[amap[j]]--
		}
		for {

			if test(nmap2) {
				r = append(r, l)
			}
			if l+lena >= len(s) {
				break
			}
			nmap2[amap[l]]++
			nmap2[amap[l+lena]]--
			l += lenw
		}

	}
	return r

}

func test(nmap []int) bool {
	for _, v := range nmap {
		if v != 0 {
			return false
		}
	}
	return true
}
