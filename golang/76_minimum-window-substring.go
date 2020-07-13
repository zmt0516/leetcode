func minWindow(s string, t string) string {
	var r string
	if len(s) == 0 || len(t) == 0 {
		return r
	}
	emap := make(map[byte]int)
	for i := range t {
		emap[t[i]] += 1
	}

	l := 0
	var h int

	for h = l; h < len(s); h++ {
		if _, ok := emap[s[h]]; ok {
			emap[s[h]]--
		}
		if test(emap) {
			r = s[l : h+1]
			break
		}
	}

	for h-l+1 > len(t) {
		if _, ok := emap[s[l]]; ok {
			emap[s[l]]++
		}
		l++
		if test(emap) {

			r = s[l : h+1]
		} else {
			h++
			if h >= len(s) {
				break
			}
			if _, ok := emap[s[h]]; ok {
				emap[s[h]]--
			}

		}

	}

	return r

}

func test(emap map[byte]int) bool {

	for _, v := range emap {
		if v > 0 {
			return false
		}
	}
	return true

}