import "strings"

func isNumber(s string) bool {

	s = strings.TrimSpace(s)

	e := strings.Index(s, "e")
	if e == 0 || e == len(s)-1 {
		return false
	}
	if e > -1 {
		if (s[e+1] == '-' || s[e+1] == '+') && e+2 < len(s) || int(s[e+1]) >= int('0') && int(s[e+1]) <= int('9') {
			for i := e + 2; i < len(s); i++ {
				if int(s[i]) < int('0') || int(s[i]) > int('9') {
					return false
				}
			}
		} else {
			return false
		}
		s = s[:e]
	}
	if s[0] == '-' || s[0] == '+' {
		if len(s) == 0 {
			return false
		}
		s = s[1:]
	}
	e = strings.Index(s, ".")
	if e > -1 {
		s = s[:e] + s[e+1:]
	}
	if len(s) == 0 {
		return false
	}
	for _, v := range s {
		if int(v) < int('0') || int(v) > int('9') {
			return false
		}
	}

	return true
}