package t299

import "fmt"

func getHint(secret string, guess string) string {
	chToNum := make(map[int32]int)
	chIndexMap := make(map[string]bool)
	for i, ch := range secret {
		chToNum[ch] = chToNum[ch] + 1
		chIndexMap[fmt.Sprintf("%d,%d", ch, i)] = true
	}
	a, b := 0, 0
	noAGuess := make([]int32, 0)
	for i, ch := range guess {
		key := fmt.Sprintf("%d,%d", ch, i)
		if _, ok := chIndexMap[key]; ok {
			delete(chIndexMap, key)
			chToNum[ch] = chToNum[ch] - 1
			a++
		} else {
			noAGuess = append(noAGuess, ch)
		}
	}
	for _, ch := range noAGuess {
		if num := chToNum[ch]; num > 0 {
			b++
			chToNum[ch]--
		}
	}
	return fmt.Sprintf("%dA%dB", a, b)
}
