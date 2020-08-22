import "strconv"

func evalRPN(tokens []string) int {

	var emap []int
	for _, v := range tokens {
		lene := len(emap)
		switch v {
		case "+":
			emap[lene-2] += emap[lene-1]
			emap = emap[:lene-1]
		case "-":
			emap[lene-2] -= emap[lene-1]
			emap = emap[:lene-1]
		case "*":
			emap[lene-2] *= emap[lene-1]
			emap = emap[:lene-1]
		case "/":
			emap[lene-2] /= emap[lene-1]
			emap = emap[:lene-1]
		default:
			v2, _ := strconv.Atoi(v)
			emap = append(emap, v2)

		}

	}
	if len(tokens) == 0 {
		return 0
	}
	return emap[0]

}