package t49

import (
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	rs := make([][]string, 0)
	chNum := make(map[string][]string)
	for _, s := range strs {
		chNumArr := make([]int, 26)
		for _, ch := range s {
			chNumArr[ch-'0']++
		}
		sb := strings.Builder{}
		for _, num := range chNumArr {
			sb.WriteString(strconv.Itoa(num))
			sb.WriteString(",")
		}
		flag := sb.String()
		if _, ok := chNum[flag]; ok {
			chNum[flag] = append(chNum[flag], s)
		} else {
			chNum[flag] = []string{s}
		}
	}
	for _, arr := range chNum {
		rs = append(rs, arr)
	}
	return rs
}
