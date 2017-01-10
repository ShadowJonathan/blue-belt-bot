package Belt

func IsSpace(L byte) bool {
	return L == ' ' || L == '\t' || L == '\r'
}

func GetArgs(S string) []string {
	Args := []string{}
	Length := len(S)
	for i := 0; i < Length; i++ {
		Letter := S[i]
		if !IsSpace(Letter) {
			var startCMD int
			start = i
			i++
			for i < Length && IsSpace(s[i]) {
				i++
			}
			Args = append(Args, S[startCMD:i])
		}
	}
	return Args
}
