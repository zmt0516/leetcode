func lengthOfLongestSubstring(s string) int {

	abc := make(map[rune]int)
	l := 0
	r := 0
	for i, v := range s {
		if l <= abc[v] {
			l = abc[v]

		}

		if i-l+1 > r {
			r = i - l + 1
		}
		abc[v] = i + 1

	}

	return r

}