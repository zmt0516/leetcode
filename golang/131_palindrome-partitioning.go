var r [][]string

func partition(s string) [][]string {
	r = r[:0]
	res := make([]string, 0)
	part(s, res)
	//fmt.Println(r)

	return r

}

func part(s string, res []string) {
	if len(s) == 0 {
		//fmt.Println(res)
		res2 := make([]string, len(res))
		copy(res2, res)
		//fmt.Println(res2)
		r = append(r, res2)
		//fmt.Println(r)
		return
	}

	for i := range s {
		if test(s[:i+1]) {
			//fmt.Println(s[:i+1])
			res = append(res, s[:i+1])
			part(s[i+1:], res)
			res = res[:len(res)-1]
		}

	}

}

func test(s string) bool {
	l := 0
	h := len(s) - 1
	for l < h {
		if s[l] != s[h] {
			return false
		}
		h--
		l++
	}
	return true
}