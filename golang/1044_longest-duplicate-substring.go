//It is not a good solution. I'll find some other one.

func longestDupSubstring(S string) string {

	oldm := make(map[int][]int)

	for i, v := range S {

		oldm[((5381<<5)+5381)+int(v)] = append(oldm[((5381<<5)+5381)+int(v)], i)
	}

	for i, v := range oldm {
		if len(v) < 2 {
			delete(oldm, i)
		}
	}
	if len(oldm) == 0 {
		return ""
	}
	n := 1

	for {

		newm := make(map[int][]int)

		for i, v := range oldm {

			for _, v2 := range v {
				if v2 < len(S)-1 {
					newm[((i<<5)+i)+int(S[v2+1])] = append(newm[((i<<5)+i)+int(S[v2+1])], v2+1)

				}

			}

		}
		for i, v := range newm {
			if len(v) < 2 {
				delete(newm, i)
			}
		}

		if len(newm) == 0 {
			for _, v := range oldm {
				return S[v[0]-n+1 : v[0]+1]
			}

		} else {
			oldm = newm

		}
		n++
		if len(oldm) > 100 {
			newm = make(map[int][]int)

			for i, v := range oldm {

				for _, v2 := range v {
					if v2 < len(S)-20 {
						hash := i
						for k := 1; k <= 20; k++ {
							hash = ((hash << 5) + hash) + int(S[v2+k])
						}
						newm[hash] = append(newm[hash], v2+20)

					}

				}

			}
			for i, v := range newm {
				if len(v) < 2 {
					delete(newm, i)
				}
			}

			if len(newm) > 1 {
				oldm = newm
				n += 20
			}

		}

		if len(newm) == 1 {

			for _, v := range newm {
				s := S[v[0]-n+1 : v[0]+1]
				//fmt.Println(n,s)
				sadd := ""
				n = 1
				i := 0
				iold := 0
				j := 1
				for {

					if v[j]+n+1 <= len(S) {
						if iold == 0 {
							if S[v[i]+n] == S[v[j]+n] {
								sadd += string(S[v[i]+n])
								//fmt.Println(i,j,v[i],v[j],n,sadd)
								n++
								continue

							}
						} else {
							if S[v[i]:v[i]+n] == S[v[j]:v[j]+n] {
								sadd = S[v[i] : v[i]+n]
								//fmt.Println(i,j,n,sadd)
								n++
								iold = 0
								continue
							}

						}

					}
					//fmt.Println("??")
					iold = 1
					j++
					if j == len(v) || v[j]+n+1 > len(S) {

						i++
						j = i + 1
					}
					if i == len(v)-1 || v[i]+n+1 > len(S) {
						break

					}

				}
				return s + sadd

			}

		}

		//n++

	}

}