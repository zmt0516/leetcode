func calculate(s string) int {

	var r int

	var emap []int

	nsign := 1

	emap = append(emap, 1)

	for i := 0; i < len(s); i++ {
		v := s[i]
		ntmp := 0
		switch {
		case v == ' ' || v == '+':
			continue

		case v >= '0' && v <= '9':
			j := i
			for ; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
				ntmp *= 10
				ntmp += int(s[j] - '0')
			}
			i = j - 1

		case v == '(':
			emap = append(emap, nsign*emap[len(emap)-1])
			nsign = 1
			continue
		case v == ')':
			emap = emap[:len(emap)-1]
			nsign = 1
			continue
		case v == '-':
			nsign = -nsign
			//fmt.Println(nsign)
			continue

		}
		//fmt.Println(nsign,ntmp,emap[len(emap)-1])
		r += nsign * ntmp * emap[len(emap)-1]
		nsign = 1

	}

	return r

}