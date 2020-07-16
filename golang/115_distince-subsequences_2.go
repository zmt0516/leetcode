func numDistinct(s string, t string) int {
	//var r int
	lent := len(t)
	lens := len(s)
	if lent > lens || lens == 0 || lent == 0 {
		return 0
	}
	emap := make([]int, lens)

	if t[0] == s[0] {
		emap[0] = 1
	}

	//n:=0
	for j := 1; j <= lens-lent; j++ {
		if t[0] == s[j] {
			//n++
			emap[j] = emap[j-1] + 1
		} else {
			emap[j] = emap[j-1]
		}
	}
	//emap[lens-lent+1]=emap[lens-lent]
	//fmt.Println(emap)
	for i := 1; i < lent; i++ {
		//n=0
		//emap[i-1]=0

		//emap[i-1]=0
		//temp:=emap[i-1]
		//emap2:=make([]int,lens)
		//temp:=emap[i-1]
		//emap[i-1]=0
		temp := emap[i-1]
		emap[i-1] = 0
		for j := i; j <= lens-lent+i; j++ {
			//temp2:=emap[j]

			if t[i] == s[j] {
				//n++
				emap[j], temp = temp+emap[j-1], emap[j]

			} else {
				emap[j], temp = emap[j-1], emap[j]
			}
			//temp=temp2

		}
		//copy(emap,emap2)
		//emap[lens-lent+i+1]=emap[lens-lent+i]
		//fmt.Println(emap,i)

	}

	return emap[lens-1]
}