var min int

func minCut(s string) int {
	if len(s) == 0 {
		return 0
	}
	var r int
	min = 10000000

	cut(s, r)

	return min - 1

}

func cut(s string, r int) {
	if r >= min {
		return
	}
	if len(s) == 0 {
		min = r
		return
	}

	for i := len(s); i > 0; i-- {

		if test(s[:i]) {
			cut(s[i:], r+1)

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