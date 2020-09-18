package t412

import "strconv"

func fizzBuzz(n int) []string {
	rs := make([]string, 0)
	for i := 1; i <= n; i++ {
		s := ""
		if i%3 != 0 && i%5 != 0 {
			rs = append(rs, strconv.Itoa(i))
			continue
		}
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		rs = append(rs, s)
	}
	return rs
}
