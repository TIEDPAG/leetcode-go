package t242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chNum := make(map[int32]int)
	for _, ch := range s {
		chNum[ch]++
	}
	for _, ch := range t {
		if _, ok := chNum[ch]; !ok {
			return false
		}
		chNum[ch]--
		if chNum[ch] == 0 {
			delete(chNum, ch)
		}
	}
	return true
}
